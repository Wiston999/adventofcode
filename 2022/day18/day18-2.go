package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Coord struct {
	X, Y, Z int
}

func (c *Coord) Neighbours(min, max int) (cs []Coord) {
	candidates := []Coord{
		Coord{c.X + 1, c.Y, c.Z},
		Coord{c.X, c.Y + 1, c.Z},
		Coord{c.X, c.Y, c.Z + 1},
		Coord{c.X - 1, c.Y, c.Z},
		Coord{c.X, c.Y - 1, c.Z},
		Coord{c.X, c.Y, c.Z - 1},
	}
	for _, cand := range candidates {
		if cand.X < min || cand.X > max {
			continue
		}
		if cand.Y < min || cand.Y > max {
			continue
		}
		if cand.Z < min || cand.Z > max {
			continue
		}
		cs = append(cs, cand)
	}
	return
}

type void struct{}

type Problem struct {
	Space    map[Coord]void
	Min, Max int
}

func (p *Problem) CountFaces() (result int) {
	// Ensure we start out of everything
	pending := []Coord{Coord{p.Min - 1, p.Min - 1, p.Min - 1}}
	visited := make(map[Coord]void)
	for len(pending) > 0 {
		current := pending[0]
		pending = pending[1:]
		log.Debug(fmt.Sprintf("Exploring %v", current))

		for _, n := range current.Neighbours(p.Min-1, p.Max+1) {
			if _, ok := visited[n]; ok {
				log.Debug(fmt.Sprintf("[%v] %v already visited", current, n))
				continue
			}
			if _, ok := p.Space[n]; ok {
				log.Debug(fmt.Sprintf("[%v] %v is LAVA", current, n))
				result += 1
			} else {
				log.Debug(fmt.Sprintf("[%v] %v is not LAVA, expanding search", current, n))
				visited[n] = void{}
				pending = append(pending, n)
			}
		}
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

func atoiSafe(a string) (i int) {
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
	strData := string(byteData)
	data.Space = make(map[Coord]void)
	data.Min = 99999
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		c := Coord{}
		fmt.Sscanf(l, "%d,%d,%d", &c.X, &c.Y, &c.Z)
		data.Space[c] = void{}
		if data.Min > c.X {
			data.Min = c.X
		}
		if data.Min > c.Y {
			data.Min = c.Y
		}
		if data.Min > c.Z {
			data.Min = c.Z
		}
		if data.Max < c.X {
			data.Max = c.X
		}
		if data.Max < c.Y {
			data.Max = c.Y
		}
		if data.Max < c.Z {
			data.Max = c.Z
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

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	result = problem.CountFaces()

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
			start := time.Now()
			echo(fmt.Sprintf("Solution is %v in %s", solution(c), time.Since(start)), c.String("output"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
