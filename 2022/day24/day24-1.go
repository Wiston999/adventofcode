package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/oleiade/lane/v2"

	"github.com/urfave/cli/v2"
)

const (
	WALL = iota
	GROUND
	ICEU
	ICED
	ICEL
	ICER
)

type Element int

type State struct {
	P     Coord
	Steps int
}

type Coord struct {
	I, J int
}

type Problem struct {
	Grid          map[Coord]Element
	Start, Target Coord
	MaxI, MaxJ    int
	Best          int
}

func (p *Problem) CheckState(s State) (result bool) {
	if p.Grid[s.P] == WALL {
		return false
	}
	for c, e := range p.Grid {
		if (s.P.I == c.I || s.P.J == c.J) && e != GROUND{
			position := c
			switch e {
			case ICEU:
				position.I = mod(c.I-1-s.Steps, p.MaxI-1) + 1
			case ICED:
				position.I = mod(c.I-1+s.Steps, p.MaxI-1) + 1
			case ICER:
				position.J = mod(c.J-1+s.Steps, p.MaxJ-1) + 1
			case ICEL:
				position.J = mod(c.J-1-s.Steps, p.MaxJ-1) + 1
			}
			if s.P == position {
				return false
			}
		}
	}
	return true
}

func (p *Problem) Neighbours(c State) (cs []State) {
	candidates := []State{
		State{Coord{c.P.I - 1, c.P.J}, c.Steps + 1},
		State{Coord{c.P.I + 1, c.P.J}, c.Steps + 1},
		State{Coord{c.P.I, c.P.J - 1}, c.Steps + 1},
		State{Coord{c.P.I, c.P.J + 1}, c.Steps + 1},
		State{Coord{c.P.I, c.P.J}, c.Steps + 1},
	}

	for _, ns := range candidates {
		if p.CheckState(ns) {
			cs = append(cs, ns)
		}
	}

	sort.Slice(cs, func(a, b int) bool {
		return dist(cs[a].P, p.Target) < dist(cs[b].P, p.Target)
	})

	return
}

func (p *Problem) PrintState(s State) (result string) {
	shifted := make(map[Coord]Element)
	for c, e := range p.Grid {
		position := c
		switch e {
		case ICEU:
			position.I = mod(c.I-1-s.Steps, p.MaxI-1) + 1
		case ICED:
			position.I = mod(c.I-1+s.Steps, p.MaxI-1) + 1
		case ICER:
			position.J = mod(c.J-1+s.Steps, p.MaxJ-1) + 1
		case ICEL:
			position.J = mod(c.J-1-s.Steps, p.MaxJ-1) + 1
		}
		if e != WALL && e != GROUND {
			shifted[position] = e
		}
	}
	for i := 0; i <= p.MaxI; i++ {
		for j := 0; j <= p.MaxJ; j++ {
			c := Coord{i, j}
			if s.P == c {
				result += "*"
				continue
			}
			if i == 0 || i == p.MaxI || j == 0 || j == p.MaxJ {
				result += "#"
				continue
			}
			switch shifted[c] {
			case ICEU:
				result += "^"
			case ICED:
				result += "v"
			case ICEL:
				result += "<"
			case ICER:
				result += ">"
			default:
				result += "."
			}
		}
		result += "\n"
	}
	return
}

func (p *Problem) SolutionAStar() (path []State, score int) {
	start := State{p.Start, 0}
	pending := lane.NewMinPriorityQueue[State, int]()
	pending.Push(start, 0)

	gScore := make(map[State]int)
	gScore[start] = 0

	cameFrom := make(map[State]State)
	cameFrom[start] = start

	fScore := make(map[State]int)
	fScore[State{p.Start, 0}] = dist(p.Start, p.Target)
	for pending.Size() > 0 {
		current, _, _ := pending.Pop()
		if current.P == p.Target {
			log.Info(fmt.Sprintf("Found solution %v", gScore[current]))
			p.Best = gScore[current]
			curr := current
			for curr != start {
				path = append(path, curr)
				curr = cameFrom[curr]
			}
			return path, gScore[current]
		}

		for _, n := range p.Neighbours(current) {
			tentative := gScore[current] + 1
			if v, ok := gScore[n]; !ok || tentative < v {
				gScore[n] = tentative
				fScore[n] = tentative + dist(n.P, p.Target)
				pending.Push(n, fScore[n])
				cameFrom[n] = current
			}
		}
	}
	return
}

func dist(a, b Coord) int {
	return int(math.Abs(float64(a.I-b.I)) + math.Abs(float64(a.J-b.J)))
}

func mod(a, b int) int {
	return (a%b + b) % b
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
	data.Grid = make(map[Coord]Element)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			switch c {
			case '.':
				data.Grid[Coord{i, j}] = GROUND
			case '#':
				data.Grid[Coord{i, j}] = WALL
			case '^':
				data.Grid[Coord{i, j}] = ICEU
			case 'v':
				data.Grid[Coord{i, j}] = ICED
			case '<':
				data.Grid[Coord{i, j}] = ICEL
			case '>':
				data.Grid[Coord{i, j}] = ICER
			}
			if i == 0 && data.Grid[Coord{i, j}] == GROUND {
				data.Start = Coord{i, j}
			}
			if j > data.MaxJ {
				data.MaxJ = j
			}
		}
		if i > data.MaxI {
			data.MaxI = i
		}
	}
	for j := 0; j <= data.MaxJ; j++ {
		if data.Grid[Coord{data.MaxI, j}] == GROUND {
			data.Target = Coord{data.MaxI, j}
		}
	}
	data.Best = data.MaxI * data.MaxJ * 3
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
	path, score := problem.SolutionAStar()
	for i, s := range path {
		log.Info(fmt.Sprintf("[%03d][%03d]: %v\n%s", score, i, s, problem.PrintState(s)))
	}
	result = len(path)

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
