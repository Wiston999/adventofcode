package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"

	astar "github.com/fzipp/astar"
)

type Node struct {
	Total, Used, Available, Percentage int
}

type Coord struct {
	X, Y int
}

type Grid map[Coord]Node

func (g Grid) Size() (x, y int) {
	for c := range g {
		if c.X > x {
			x = c.X
		}
		if c.Y > y {
			y = c.Y
		}
	}
	return
}

type State struct {
	Empty, Current Coord
}

func (g Grid) Neighbours(c State) (cs []State) {
	coords := []Coord{
		Coord{c.Empty.X - 1, c.Empty.Y},
		Coord{c.Empty.X + 1, c.Empty.Y},
		Coord{c.Empty.X, c.Empty.Y - 1},
		Coord{c.Empty.X, c.Empty.Y + 1},
	}

	for _, coord := range coords {
		if cell, ok := g[coord]; ok && coord != c.Current && !g.TooBig(coord) {
			if cell.Used <= g[c.Empty].Total {
				cs = append(cs, State{coord, c.Current})
			}
		}
	}
	if (c.Current.X == c.Empty.X) && math.Abs(float64(c.Current.Y-c.Empty.Y)) == 1 {
		cs = append(cs, State{
			c.Current,
			c.Empty,
		})
	}
	if (c.Current.Y == c.Empty.Y) && math.Abs(float64(c.Current.X-c.Empty.X)) == 1 {
		cs = append(cs, State{
			c.Current,
			c.Empty,
		})
	}
	log.Debug(fmt.Sprintf("Generated from %#v neighours %#v", c, cs))
	return
}

func (g Grid) Empties() (cs []Coord) {
	for c := range g {
		if g[c].Used == 0 {
			cs = append(cs, c)
		}
	}
	return
}

func (g Grid) TooBig(c Coord) (result bool) {
	empties := g.Empties()
	var walls []Coord
	for _, b := range empties {
		if g[b].Total < g[c].Used {
			walls = append(walls, b)
		}
	}
	return len(empties) == len(walls)
}

func (g Grid) Print() (result string) {
	X, Y := g.Size()
	for i := 0; i <= Y; i++ {
		for j := 0; j <= X; j++ {
			if i == 0 && j == 0 {
				result += "(.)"
			} else if i == 0 && j == X {
				result += " G "
			} else if g[Coord{j, i}].Used == 0 {
				result += " _ "
			} else if g.TooBig(Coord{j, i}) {
				result += " # "
			} else {
				result += " . "
			}
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

func atoi(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
}

func parseInput(input string) (data Grid, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	regex := *regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)%`)
	res := regex.FindAllStringSubmatch(string(tmpData), -1)
	data = make(Grid)
	for _, r := range res {
		c := Coord{X: atoi(r[1]), Y: atoi(r[2])}
		data[c] = Node{
			Total:      atoi(r[3]),
			Used:       atoi(r[4]),
			Available:  atoi(r[5]),
			Percentage: atoi(r[6]),
		}
		log.Debug(fmt.Sprintf("Parsed node at %#v: %#v", c, data[c]))
	}
	return
}

func heuristic(s, e State) float64 {
	return float64(s.Current.Y + s.Current.X + s.Empty.Y + s.Empty.X)
}

func cost(e, s State) float64 {
	return 1
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	data, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	fmt.Println(data.Print())
	x, _ := data.Size()
	path := astar.FindPath[State](&data, State{data.Empties()[0], Coord{x, 0}}, State{Coord{1, 0}, Coord{0, 0}}, cost, heuristic)
	for i, s := range path {
		log.Info(fmt.Sprintf("[%03d] Path: %#v", i, s))
	}
	result = len(path) - 1
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
