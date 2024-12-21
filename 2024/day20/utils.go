package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/oleiade/lane/v2"

	"github.com/urfave/cli/v2"
)

type void struct{}

var null void

type Problem struct {
	Map        map[Coord]string
	MaxX, MaxY int
	Start, End Coord
}

func (p *Problem) FindPath() (result int) {
	pathFinder := PathFinder{
		P:     *p,
		Start: State{p.Start},
		Goal: func(s State) bool {
			return s.C == p.End
		},
		Cost: func(n, c State) float64 {
			return 1
		},
		Heuristic: func(s State) float64 {
			return s.C.Manhattan(p.End)
		},
	}
	_, score := pathFinder.Search()
	result = int(score)
	return
}

type Cheat struct {
	A, B Coord
}

func (p *Problem) Part1() (result int) {
	cheats := make(map[Cheat]int)

	base := p.FindPath()
	for c, t := range p.Map {
		if t == "#" {
			if p.Map[Coord{c.X - 1, c.Y}] == "." && p.Map[Coord{c.X + 1, c.Y}] == "." {
				p.Map[c] = "."
				cheats[Cheat{c, Coord{c.X + 1, c.Y}}] = base - p.FindPath()
				p.Map[c] = "#"
			}
			if p.Map[Coord{c.X, c.Y + 1}] == "." && p.Map[Coord{c.X, c.Y - 1}] == "." {
				p.Map[c] = "."
				cheats[Cheat{c, Coord{c.X, c.Y + 1}}] = base - p.FindPath()
				p.Map[c] = "#"
			}
		}
	}

	for c, v := range cheats {
		log.Debug(fmt.Sprintf("Cheat %v saved %d", c, v))
		if v >= 100 {
			result += 1
		}
	}
	log.Info(fmt.Sprintf("Found %d cheats and %d improves 100 picoseconds", len(cheats), result))
	return
}

func (p *Problem) Part2() (result int) {
	return
}

func NewProblem(ctx *cli.Context) (p *Problem, err error) {
	input := ctx.String("input")
	byteData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	p = new(Problem)
	p.Map = make(map[Coord]string)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range strings.Split(strings.TrimSpace(l), "") {
			if c == "S" {
				p.Start = Coord{i, j}
				p.Map[Coord{i, j}] = "."
			} else if c == "E" {
				p.End = Coord{i, j}
				p.Map[Coord{i, j}] = "."
			} else {
				p.Map[Coord{i, j}] = c
			}
			if j > p.MaxY {
				p.MaxY = j
			}
		}
		if i > p.MaxX {
			p.MaxX = i
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

func atoi(a string) (i int) {
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

type State struct {
	C Coord
}

func (s *State) Neighbours(p Problem) (ns []State) {
	candidates := []State{
		{Coord{s.C.X + 1, s.C.Y}},
		{Coord{s.C.X - 1, s.C.Y}},
		{Coord{s.C.X, s.C.Y + 1}},
		{Coord{s.C.X, s.C.Y - 1}},
	}
	for _, c := range candidates {
		if p.Map[c.C] != "#" {
			ns = append(ns, c)
		}
	}
	return
}

type PathFinder struct {
	P         Problem
	Start     State
	Goal      func(State) bool
	Cost      func(State, State) float64
	Heuristic func(State) float64
}

func (p *PathFinder) Search() (path []State, score float64) {
	start := p.Start
	pending := lane.NewMinPriorityQueue[State, float64]()
	pending.Push(start, 0)

	gScore := make(map[State]float64)
	gScore[start] = 0

	cameFrom := make(map[State]State)
	cameFrom[start] = start

	fScore := make(map[State]float64)
	fScore[p.Start] = p.Heuristic(p.Start)
	for pending.Size() > 0 {
		current, _, _ := pending.Pop()
		if p.Goal(current) {
			log.Info(fmt.Sprintf("PathFinder Found solution %v", gScore[current]))
			curr := current
			for curr != start {
				path = append(path, curr)
				curr = cameFrom[curr]
			}
			return path, gScore[current]
		}

		for _, n := range current.Neighbours(p.P) {
			tentative := gScore[current] + p.Cost(n, current)
			if v, ok := gScore[n]; !ok || tentative < v {
				gScore[n] = tentative
				fScore[n] = tentative + p.Heuristic(n)
				pending.Push(n, fScore[n])
				cameFrom[n] = current
			}
		}
	}
	return
}

type Coord struct {
	X, Y int
}

func (c *Coord) Manhattan(oc Coord) float64 {
	return math.Abs(float64(oc.X-c.X)) + math.Abs(float64(oc.Y-c.Y))
}
