package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Coord struct {
	X, Y int
}

const (
	BLANK = iota
	ROCK
	SAND
)

type Edge int

func (e Edge) Print() string {
	if e == BLANK {
		return "."
	} else if e == ROCK {
		return "#"
	}
	return "o"
}

type Problem struct {
	Map        map[Coord]Edge
	MinX, MinY int
	MaxX, MaxY int
	SandStart  Coord
}

func (p *Problem) DropSand() (stable bool) {
	current := p.SandStart
	for current.Y <= p.MaxY {
		log.Debug(fmt.Sprintf("Sand at %v", current))
		if p.Map[Coord{current.X, current.Y + 1}] == BLANK {
			current = Coord{current.X, current.Y + 1}
		} else if p.Map[Coord{current.X - 1, current.Y + 1}] == BLANK {
			current = Coord{current.X - 1, current.Y + 1}
		} else if p.Map[Coord{current.X + 1, current.Y + 1}] == BLANK {
			current = Coord{current.X + 1, current.Y + 1}
		} else {
			p.Map[current] = SAND
			return true
		}
	}
	return
}

func (p *Problem) Print() (result string) {
	result += "\n"
	for i := p.MinY; i <= p.MaxY; i++ {
		for j := p.MinX; j <= p.MaxX; j++ {
			result += p.Map[Coord{j, i}].Print()
		}
		result += "\n"
	}
	return
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

func Atoi(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
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
	data.Map = make(map[Coord]Edge)
	data.MinX = 99999999999
	data.MinY = 99999999999
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		steps := strings.Split(l, " -> ")
		c0 := strings.Split(steps[0], ",")
		curr := Coord{Atoi(c0[0]), Atoi(c0[1])}
		for j := 1; j < len(steps); j++ {
			cNext := strings.Split(steps[j], ",")
			next := Coord{Atoi(cNext[0]), Atoi(cNext[1])}
			if curr.X == next.X {
				higher, lower := curr.Y, next.Y
				if higher < lower {
					higher, lower = lower, higher
				}
				if lower < data.MinY {
					data.MinY = lower
				}
				if higher > data.MaxY {
					data.MaxY = higher
				}
				log.Debug(fmt.Sprintf("Rock from: {%d, %d} to {%d, %d}", curr.X, lower, curr.X, higher))
				for k := lower; k <= higher; k++ {
					data.Map[Coord{curr.X, k}] = ROCK
				}
			} else {
				higher, lower := curr.X, next.X
				if higher < lower {
					higher, lower = lower, higher
				}
				if lower < data.MinX {
					data.MinX = lower
				}
				if higher > data.MaxX {
					data.MaxX = higher
				}
				log.Debug(fmt.Sprintf("Rock from: {%d, %d} to {%d, %d}", lower, curr.Y, higher, curr.Y))
				for k := lower; k <= higher; k++ {
					data.Map[Coord{k, curr.Y}] = ROCK
				}

			}
			curr = next
		}
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	problem.SandStart = Coord{500, 0}
	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	log.Debug(fmt.Sprintf("Map:%s", problem.Print()))
	for result = 1; problem.DropSand(); result++ {
		log.Info(fmt.Sprintf("[%03d] Sand rested", result))
		if log.GetLevel() == log.DebugLevel {
			log.Debug(fmt.Sprintf("Map:%s", problem.Print()))
		}
	}
	result--

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
