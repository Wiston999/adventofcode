package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"

	astar "github.com/fzipp/astar"
)

const (
	Wall = iota
	Way
	Target
)

type State struct {
	C       Coord
	Targets string
}

type Coord struct {
	X, Y int
}

type TileType int

type Map struct {
	Tiles   map[Coord]TileType
	Targets map[Coord]int
}

func (m Map) Neighbours(s State) (states []State) {
	candidates := []State{
		State{C: Coord{s.C.X + 1, s.C.Y}},
		State{C: Coord{s.C.X - 1, s.C.Y}},
		State{C: Coord{s.C.X, s.C.Y + 1}},
		State{C: Coord{s.C.X, s.C.Y - 1}},
	}
	for _, c := range candidates {
		if v, ok := m.Tiles[c.C]; ok && v != Wall {
			c.Targets = strings.Clone(s.Targets)
			if v == Target {
				index := m.Targets[c.C]
				c.Targets = c.Targets[:index] + "1" + c.Targets[index+1:]
			}
			states = append(states, c)
		}
	}
	log.Debug(fmt.Sprintf("Generated states from %#v: %#v", s, states))

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

func parseInput(input string) (data Map, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	data.Tiles = make(map[Coord]TileType)
	data.Targets = make(map[Coord]int)
	for j, l := range strings.Split(strings.TrimSpace(string(tmpData)), "\n") {
		for i, c := range strings.TrimSuffix(l, "\n") {
			var v TileType
			v = Wall
			if c == '.' {
				v = Way
			} else if c != '#' {
				v = Target
				n, _ := strconv.Atoi(string(c))
				data.Targets[Coord{i, j}] = n
			}
			log.Debug(fmt.Sprintf("Parsed tile %#v: %v [%3v]", Coord{i, j}, v, data.Targets))
			data.Tiles[Coord{i, j}] = v
		}
	}
	return
}

func cost(e, s State) float64 {
	return 1
}

func heuristic(s, e State) float64 {
	h := 0.0
	for i := range s.Targets {
		if s.Targets[i] == '0' {
			h += 10.0
			h += math.Abs(float64(s.C.X-e.C.X)) + math.Abs(float64(s.C.Y-e.C.Y))
		}
	}
	return 0
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	data, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	startCoord := Coord{}
	for c, v := range data.Targets {
		if v == 0 {
			startCoord = c
			break
		}
	}
	result = 100000000
	for c, v := range data.Targets {
		if v != 0 {
			start := State{C: startCoord, Targets: "1" + strings.Repeat("0", len(data.Targets)-1)}
			target := State{C: c, Targets: strings.Repeat("1", len(data.Targets))}

			log.Info(fmt.Sprintf("Searching from %#v to %#v", startCoord, c))
			path := astar.FindPath[State](data, start, target, cost, heuristic)
			log.Debug(fmt.Sprintf("Found path: (%03d) %#v", len(path), path))
			if len(path) > 0 && len(path) < result {
				result = len(path) - 1
				for i, step := range path {
					log.Info(fmt.Sprintf("Found shortest path: (%03d) %#v", i, step))
				}
			}
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
