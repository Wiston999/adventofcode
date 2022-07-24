package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"

	astar "github.com/fzipp/astar"
)

type Coord struct {
	X, Y int
}

type Map struct {
	Tiles    map[Coord]bool
	Designer int
	Current  Coord
}

// True means position cannot be used
func (m *Map) Compute(c Coord) bool {
	if c.X < 0 || c.Y < 0 {
		return true
	}
	v := c.X*c.X + 3*c.X + 2*c.X*c.Y + c.Y + c.Y*c.Y + m.Designer
	binary := fmt.Sprintf("%b", v)
	if (strings.Count(binary, "1") % 2) == 0 {
		return false
	}
	return true
}

func (m *Map) Get(c Coord) bool {
	if v, ok := m.Tiles[c]; !ok {
		v = m.Compute(c)
		return v
	} else {
		return v
	}
}

func (m *Map) Print(c Coord, path []Coord) (result string) {
	for i := 0; i <= c.Y; i++ {
		for j := 0; j <= c.X; j++ {
			current := Coord{j, i}
			if m.Get(current) {
				result += "#"
			} else if len(path) > 0 {
				if path[0] == current {
					result += "S"
				} else if path[len(path)-1] == current {
					result += "E"
				} else if contains(current, path) {
					result += "O"
				} else {
					result += "."
				}
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return
}

func (m *Map) Neighbours(c Coord) (cs []Coord) {
	coordList := []Coord{
		Coord{c.X, c.Y + 1},
		Coord{c.X, c.Y - 1},
		Coord{c.X + 1, c.Y},
		Coord{c.X - 1, c.Y},
	}

	for _, nc := range coordList {
		if !m.Get(nc) {
			cs = append(cs, nc)
		}
	}
	return
}

func contains(a Coord, list []Coord) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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

func step(m Map, c Coord) (cs []Coord) {
	coordList := []Coord{
		Coord{c.X, c.Y + 1},
		Coord{c.X, c.Y - 1},
		Coord{c.X + 1, c.Y},
		Coord{c.X - 1, c.Y},
	}

	for _, nc := range coordList {
		if !m.Get(nc) {
			cs = append(cs, nc)
		}
	}
	return
}

func cost(e, s Coord) float64 {
	return 1
}

// Manhattan
func heuristic(s, e Coord) float64 {
	return math.Abs(float64(s.X-e.X)) + math.Abs(float64(s.Y-e.Y))
}

func solution(context *cli.Context) (result int) {
	m := Map{
		Designer: context.Int("input"),
		Current:  Coord{1, 1},
	}
	target := Coord{
		context.Int("x"),
		context.Int("y"),
	}
	log.Info(fmt.Sprintf("Looking for path to (%d, %d) with designer number %d", target.X, target.Y, m.Designer))
	log.Info(fmt.Sprintf("Initial map setup is:\n%s", m.Print(target, []Coord{})))
	path := astar.FindPath[Coord](&m, m.Current, target, cost, heuristic)
	if path == nil {
		log.Error("No path found with input provided")
	} else {
		result = len(path) - 1
		maxCoord := target
		for _, c := range path {
			if c.X > maxCoord.X {
				maxCoord.X = c.X
			}
			if c.Y > maxCoord.Y {
				maxCoord.Y = c.Y
			}
		}
		log.Info(fmt.Sprintf("Found shortest path:\n%s", m.Print(maxCoord, path)))
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
			&cli.IntFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   1352,
				Usage:   "Input value",
			},
			&cli.IntFlag{
				Name:  "x",
				Value: 31,
				Usage: "X coordinate target",
			},
			&cli.IntFlag{
				Name:  "y",
				Value: 39,
				Usage: "Y coordinate target",
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
			log.Debug(fmt.Sprintf("Input value:  %d", c.Int("input")))
			log.Debug(fmt.Sprintf("Target:  (%d, %d)", c.Int("x"), c.Int("y")))
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
