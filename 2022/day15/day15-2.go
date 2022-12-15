package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Range struct {
	Start, End int
}

type Coord struct {
	X, Y int
}

const (
	BLANK = iota
	SENSOR
	BEACON
	COVERED
)

type Element int

type Pair struct {
	Sensor, Beacon Coord
	Radius         int
}

func (p *Pair) GetRadius() int {
	if p.Radius == 0 {
		p.Radius = p.GetDistance(p.Beacon)
	}
	return p.Radius
}

func (p *Pair) GetDistance(c Coord) int {
	return int(math.Abs(float64(p.Sensor.X)-float64(c.X)) + math.Abs(float64(p.Sensor.Y)-float64(c.Y)))
}

type Problem struct {
	Ranges map[int][]Range
	Pairs  []Pair
}

func (p *Problem) GetCoord(c Coord) (result Element) {
	for _, r := range p.Ranges[c.X] {
		if r.Start <= c.Y && r.End >= c.Y {
			return COVERED
		}
	}
	return BLANK
}

func (p *Problem) CoverPlaces(limit int) {
	for _, pair := range p.Pairs {
		r := pair.GetRadius()
		for i := -r; i <= r; i++ {
			realI := pair.Sensor.X + i
			if realI >= 0 && realI <= limit {
				absI := int(math.Abs(float64(i)))
				r := Range{pair.Sensor.Y - r + absI, pair.Sensor.Y + r - absI}
				p.Ranges[realI] = append(p.Ranges[realI], r)
				currLen := len(p.Ranges[realI]) + 1
				for currLen != len(p.Ranges[realI]) {
					currLen = len(p.Ranges[realI])
					p.MergeRanges(realI)
				}
			}
		}
	}
}

func (p *Problem) MergeRanges(i int) {
	sort.Slice(p.Ranges[i], func(a, b int) bool {
		return p.Ranges[i][a].Start < p.Ranges[i][b].Start
	})
	log.Debug(fmt.Sprintf("Merging %v", p.Ranges[i]))
	ranges := make([]Range, 0)
	for j := 0; j < len(p.Ranges[i]); j++ {
		nr := p.Ranges[i][j]
		for k := j + 1; k < len(p.Ranges[i]); k++ {
			if (p.Ranges[i][k].Start-1) <= nr.End && nr.End <= p.Ranges[i][k].End {
				nr.End = p.Ranges[i][k].End
				j++
			} else if p.Ranges[i][k].Start >= nr.Start && p.Ranges[i][k].End <= nr.End {
				j++
			}
		}
		ranges = append(ranges, nr)
	}
	p.Ranges[i] = ranges
	log.Debug(fmt.Sprintf("Merged %v", p.Ranges[i]))
}

func setLogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.Fatal(fmt.Sprintf("Unkown log level %s", level))
	}
}

func echo(msg, f string) {
	var file *os.File
	if f == "-" {
		file = os.Stdout
	} else {
		var err error
		file, err = os.OpenFile(f, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Error(fmt.Sprintf("Error opening file %s for writing output: %v", f, err))
			return
		}
		defer file.Close()
	}
	fmt.Fprintf(file, msg)
}

func parseInput(input string) (data Problem, err error) {
	byteData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
		pair := Pair{}
		read, err := fmt.Sscanf(
			l,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n",
			&pair.Sensor.X, &pair.Sensor.Y, &pair.Beacon.X, &pair.Beacon.Y,
		)
		if err != nil {
			log.Warn(fmt.Sprintf("Error scanning input [%02d]: %s", read, err))
		}
		data.Pairs = append(data.Pairs, pair)
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var limit = context.Int("limit")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	problem.Ranges = make(map[int][]Range)
	problem.CoverPlaces(limit)
	log.Debug(fmt.Sprintf("Parsed problem [limit=%d] %#v", limit, problem))
	log.Info("Places computed")
	for i, r := range problem.Ranges {
		if len(r) > 1 {
			result = 4000000*i + r[0].End + 1
		}
	}
	// for i := 0; i < limit; i++ {
	//
	// 	for j := 0; j < limit; j++ {
	// 		if problem.GetCoord(Coord{i, j}) == BLANK {
	// 			log.Info(fmt.Sprintf("%v is not covered", Coord{i, j}))
	// 			result = 4000000*i + j
	// 			break
	// 		}
	// 	}
	// 	if result != 0 {
	// 		break
	// 	}
	// }

	return
}

func main() {
	app := &cli.App{
		Name:  "AOC",
		Usage: "Solve AdventOfCode problem!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "loglevel",
				Aliases: []string{"l"},
				Value:   "info",
				Usage:   "Log level output",
			},
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   "input.txt",
				Usage:   "Input file path",
			},
			&cli.IntFlag{
				Name:    "limit",
				Aliases: []string{"x"},
				Value:   4000000,
				Usage:   "Row to check for elements",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Value:   "-",
				Usage:   "Output file path, use - for stdout",
			},
		},
		Action: func(c *cli.Context) error {
			setLogLevel(c.String("loglevel"))
			log.Debug("Received parameters:")
			log.Debug(fmt.Sprintf("Input file name:  %s", c.String("input")))
			log.Debug(fmt.Sprintf("Output file name: %s", c.String("output")))
			log.Debug(fmt.Sprintf("Log level:        %s", c.String("loglevel")))
			echo(fmt.Sprintf("Solution is %v", solution(c)), c.String("output"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
