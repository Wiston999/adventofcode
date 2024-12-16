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
	Map                 map[Coord]string
	Current, Start, End Coord
	Shortest            []State
	Best                int
}

func (p *Problem) ShortestPath(start, end State) (path []State, score float64) {
	pathFinder := PathFinder{
		P:     *p,
		Start: start,
		Goal: func(s State) bool {
			return s.Current.X == end.Current.X && s.Current.Y == end.Current.Y
		},
		Cost: func(n, c State) float64 {
			if n.Direction == c.Direction {
				return 1.0
			} else {
				return 1000.0
			}
		},
		Heuristic: func(s State) float64 {
			rotateCost := 0.0
			switch s.Direction {
			case South:
				rotateCost = 2
			case East:
				rotateCost = 1
			case West:
				rotateCost = 1
			}
			return s.Current.Manhattan(p.End) + 1000*rotateCost
		},
	}
	path, score = pathFinder.Search()
	path = append(path, start)
	return
}

func (p *Problem) PathScore(path []State) (result int) {
	current := path[0]
	for i := 1; i < len(path); i += 1 {
		if current.Direction != path[i].Direction {
			result += 1000
		} else {
			result += 1
		}
		current = path[i]
	}
	return
}

func (p *Problem) Part1() (result int) {
	path, score := p.ShortestPath(State{p.Start, East}, State{p.End, North})
	log.Debug(fmt.Sprintf("Found path to %v [%02f]: %v", p.End, score, path))
	p.Shortest = path
	p.Best = int(score)
	result = p.Best
	return
}

func (p *Problem) Part2() (result int) {
	if len(p.Shortest) == 0 {
		log.Fatal("First part must be run to get best path")
	}
	pending := lane.NewMinPriorityQueue[State, float64]()
	pending.Push(State{p.Start, East}, 0)
	visited := make(map[State]void)
	solution := make(map[Coord]void)

	for pending.Size() > 0 {
		current, _, _ := pending.Pop()
		_, toCurrent := p.ShortestPath(State{p.Start, East}, current)
		_, toEnd := p.ShortestPath(current, State{p.End, North})
		// log.Debug(fmt.Sprintf("From %v: %f + %f <= %d", current, toCurrent, toEnd, p.Best))
		if int(toCurrent+toEnd) == p.Best {
			solution[current.Current] = null
		}
		for _, n := range current.Neighbours(*p) {
			if _, ok := visited[n]; !ok {
				pending.Push(n, 0)
				visited[n] = null
			}
		}
	}
	result = len(solution)
	log.Debug(fmt.Sprintf("%v", solution))
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
	Current   Coord
	Direction Dir
}

type Dir int

const (
	North = iota
	East
	South
	West
)

func (s *State) Neighbours(p Problem) (ns []State) {
	candidates := []State{}
	switch s.Direction {
	case North:
		candidates = append(candidates, State{
			Coord{s.Current.X - 1, s.Current.Y},
			s.Direction,
		})
	case South:
		candidates = append(candidates, State{
			Coord{s.Current.X + 1, s.Current.Y},
			s.Direction,
		})
	case East:
		candidates = append(candidates, State{
			Coord{s.Current.X, s.Current.Y - 1},
			s.Direction,
		})
	case West:
		candidates = append(candidates, State{
			Coord{s.Current.X, s.Current.Y + 1},
			s.Direction,
		})
	}
	candidates = append(candidates, State{
		Coord{s.Current.X, s.Current.Y},
		(s.Direction + 1) % 4,
	})
	candidates = append(candidates, State{
		Coord{s.Current.X, s.Current.Y},
		(s.Direction + 3) % 4,
	})
	for _, c := range candidates {
		if v, ok := p.Map[c.Current]; ok && v != "#" {
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
