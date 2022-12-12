package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"

	astar "github.com/fzipp/astar"
)

type Coord struct {
	X, Y int
}

type Problem struct {
	Tiles           map[Coord]int
	Current, Target Coord
	MaxX, MaxY      int
}

func (p *Problem) Print() (result string) {
	for i := 0; i <= p.MaxX; i++ {
		for j := 0; j <= p.MaxY; j++ {
			if p.Current.X == i && p.Current.Y == j {
				result += "C  "
			} else if p.Target.X == i && p.Target.Y == j {
				result += "T  "
			} else {
				result += fmt.Sprintf("%02d ", p.Tiles[Coord{i, j}])
			}
		}
		result += "\n"
	}
	return
}

func (p *Problem) Neighbours(c Coord) (cs []Coord) {
	coordList := []Coord{
		Coord{c.X, c.Y + 1},
		Coord{c.X, c.Y - 1},
		Coord{c.X + 1, c.Y},
		Coord{c.X - 1, c.Y},
	}

	for _, nc := range coordList {
		if _, ok := p.Tiles[nc]; ok {
			if (p.Tiles[nc] - p.Tiles[c]) <= 1 {
				cs = append(cs, nc)
			}
		}
	}

	log.Debug(fmt.Sprintf("Neighbours from {%02d, %02d}: %v", c.X, c.Y, cs))
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

func parseInput(input string) (data Problem, err error) {
	byteData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	strData := string(byteData)
	data.Tiles = make(map[Coord]int)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range strings.TrimSpace(l) {
			if c == 'S' {
				data.Current = Coord{i, j}
				c = 'a'
			} else if c == 'E' {
				data.Target = Coord{i, j}
				c = 'z'
			}
			data.Tiles[Coord{i, j}] = int(byte(c) - 'a')
			if j > data.MaxY {
				data.MaxY = j
			}
		}
		if i > data.MaxX {
			data.MaxX = i
		}
	}
	return
}

func cost(e, s Coord) float64 {
	return 1
}

// Manhattan
func heuristic(s, e Coord) float64 {
	return 0 // math.Abs(float64(s.X-e.X)) + math.Abs(float64(s.Y-e.Y))
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	echo(fmt.Sprintf("Read Map: \n%s", problem.Print()), context.String("output"))
	path := astar.FindPath[Coord](&problem, problem.Current, problem.Target, cost, heuristic)
	for i, p := range path {
		log.Info(fmt.Sprintf("[%03d] Step: %v", i, p))
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
