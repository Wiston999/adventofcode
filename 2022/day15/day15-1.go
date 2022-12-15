package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Coord struct {
	X, Y int
}

const (
	BLANK = iota
	SENSOR
	BEACON
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
	Grid       map[Coord]Element
	Pairs      []Pair
	MinX, MinY int
	MaxX, MaxY int
}

func (p *Problem) UpdateBoundaries() {
	for _, pair := range p.Pairs {
		r := pair.GetRadius()
		if (pair.Sensor.X - r) < p.MinX {
			p.MinX = pair.Sensor.X - r
		}
		if (pair.Sensor.X + r) > p.MaxX {
			p.MaxX = pair.Sensor.X + r
		}
		if (pair.Sensor.Y - r) < p.MinY {
			p.MinY = pair.Sensor.Y - r
		}
		if (pair.Sensor.Y + r) > p.MaxY {
			p.MaxY = pair.Sensor.Y + r
		}
	}
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
	data.Grid = make(map[Coord]Element)
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
		data.Grid[pair.Sensor] = SENSOR
		data.Grid[pair.Beacon] = BEACON
		data.Pairs = append(data.Pairs, pair)
	}
	data.UpdateBoundaries()
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var Y = context.Int("row")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))

	for i := problem.MinX; i <= problem.MaxX; i++ {
		closer := false
		for _, p := range problem.Pairs {
			if p.GetDistance(Coord{i, Y}) <= p.GetRadius() {
				closer = true
				break
			}
		}
		if closer && problem.Grid[Coord{i, Y}] == BLANK {
			log.Debug(fmt.Sprintf("%v cannot contain any sensor", Coord{i, Y}))
			result += 1
		}
	}

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
				Name:    "row",
				Aliases: []string{"y"},
				Value:   2000000,
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
