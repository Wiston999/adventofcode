package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	U string = "U"
	D        = "D"
	L        = "L"
	R        = "R"
)

type Coord struct {
	X, Y int
}

type Step struct {
	Direction string
	Count     int
}

type Problem struct {
	Visited map[Coord]int
	Rope    []Coord
	Tails   int
	Steps   []Step
}

func (p *Problem) Distance(t int) int {
	if p.Rope[t-1].X == p.Rope[t].X {
		return int(math.Abs(float64(p.Rope[t-1].Y) - float64(p.Rope[t].Y)))
	}

	if p.Rope[t-1].Y == p.Rope[t].Y {
		return int(math.Abs(float64(p.Rope[t-1].X) - float64(p.Rope[t].X)))
	}
	return int(math.Abs(float64(p.Rope[t-1].X)-float64(p.Rope[t].X))+math.Abs(float64(p.Rope[t-1].Y)-float64(p.Rope[t].Y))) - 1
}

func (p *Problem) MoveTail(t int) {
	if p.Rope[t].X == p.Rope[t-1].X {
		if p.Rope[t].Y < p.Rope[t-1].Y {
			p.Rope[t].Y++
		} else {
			p.Rope[t].Y--
		}
	} else if p.Rope[t].Y == p.Rope[t-1].Y {
		if p.Rope[t].X < p.Rope[t-1].X {
			p.Rope[t].X++
		} else {
			p.Rope[t].X--
		}
	} else { // Move diagonally
		if p.Rope[t].X < p.Rope[t-1].X {
			p.Rope[t].X++
		} else {
			p.Rope[t].X--
		}
		if p.Rope[t].Y < p.Rope[t-1].Y {
			p.Rope[t].Y++
		} else {
			p.Rope[t].Y--
		}
	}

	log.Debug(fmt.Sprintf("Moved Rope[t] to {%02d, %02d} - {%02d, %02d}", p.Rope[t].X, p.Rope[t].Y, p.Rope[t-1].X, p.Rope[t-1].Y))
	if t == p.Tails {
		p.Visited[p.Rope[t]]++
	}
}

func (p *Problem) Run() {
	p.Visited[p.Rope[p.Tails]]++
	for i, s := range p.Steps {
		log.Debug(fmt.Sprintf("[%03d] Processing step: %v", i, s))
		for j := 0; j < s.Count; j++ {
			switch s.Direction {
			case U:
				p.Rope[0].Y--
			case D:
				p.Rope[0].Y++
			case R:
				p.Rope[0].X++
			case L:
				p.Rope[0].X--
			}
			for t := 1; t <= p.Tails; t++ {
				if p.Distance(t) > 1 {
					p.MoveTail(t)
				}
			}
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
	data.Visited = make(map[Coord]int)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		parts := strings.Split(l, " ")
		c, _ := strconv.Atoi(parts[1])
		m := Step{
			Direction: parts[0],
			Count:     c,
		}
		data.Steps = append(data.Steps, m)
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var tails = context.Int("tails")
	problem, err := parseInput(input)
	problem.Tails = tails
	problem.Rope = make([]Coord, tails+1)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	problem.Run()
	log.Info(fmt.Sprintf("Head Finished at {%02d, %02d}", problem.Rope[0].X, problem.Rope[0].Y))
	log.Info(fmt.Sprintf("Tail Finished at {%02d, %02d}", problem.Rope[problem.Tails].X, problem.Rope[problem.Tails].Y))

	for c, v := range problem.Visited {
		if v > 0 {
			result++
			log.Debug(fmt.Sprintf("{%02d, %02d} visited %03d times", c.X, c.Y, v))
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
				Name:    "tails",
				Aliases: []string{"t"},
				Value:   1,
				Usage:   "Tail length, use 1 for part 1, use 9 for part 2",
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
