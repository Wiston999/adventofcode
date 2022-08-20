package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	VERTICAL   = '|'
	HORIZONTAL = '-'
	CORNER     = '+'
)

type NodeType rune

type Node NodeType

type Coord struct {
	I, J int
}

type Problem struct {
	Nodes   map[Coord]Node
	Visited map[Coord]int
	Start   Coord
	Path    string
	Steps   int
}

func (p *Problem) Trace() string {
	current := p.Start
	inertia := 1 // 1 is horizontal, -1 vertical
	if p.Nodes[p.Start] == VERTICAL {
		inertia = -1
	}
	for {
		p.Steps++
		log.Info(fmt.Sprintf("Visiting %v", current))
		p.Visited[current] = inertia
		if p.Nodes[current] >= 'A' && p.Nodes[current] <= 'Z' {
			p.Path += string(p.Nodes[current])
			log.Info(fmt.Sprintf("Found letter in the path %c: %s", p.Nodes[current], p.Path))
		}
		neighbours := []Coord{}
		if inertia == -1 {
			neighbours = []Coord{
				Coord{current.I + 1, current.J},
				Coord{current.I - 1, current.J},
			}
		} else {
			neighbours = []Coord{
				Coord{current.I, current.J + 1},
				Coord{current.I, current.J - 1},
			}
		}
		next := current
		for _, n := range neighbours {
			_, exist := p.Nodes[n]
			vinertia, visited := p.Visited[n]
			log.Debug(fmt.Sprintf("Testing neighbour %v of %v: E: %v V: %v", n, current, exist, visited))
			if exist && !visited {
				next = n
			} else if exist && visited && vinertia != inertia {
				next = n
			}
		}
		if next == current {
			// No new path
			break
		}
		if p.Nodes[next] == CORNER {
			// Corners must change its direction
			inertia *= -1
		}
		current = next
	}
	return p.Path
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
	data.Nodes = make(map[Coord]Node)
	data.Visited = make(map[Coord]int)
	for i, l := range strings.Split(strData, "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			var n Node
			if c == '|' {
				n = VERTICAL
			} else if c == '-' {
				n = HORIZONTAL
			} else if c == '+' {
				n = CORNER
			} else if c != ' ' {
				n = Node(c)
			}
			if n != 0 {
				if i == 0 || j == 0 {
					data.Start = Coord{i, j}
				}
				data.Nodes[Coord{i, j}] = n
			}
		}
	}
	return
}

func solution(context *cli.Context) (result string) {
	var input = context.String("input")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))

	result = problem.Trace()
	log.Info(fmt.Sprintf("Visited %d places", problem.Steps))
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
