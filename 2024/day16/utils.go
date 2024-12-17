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
	MaxX, MaxY          int
	Current, Start, End Coord
	Dist                map[State]int
	Prev                map[State]State
}

func (p *Problem) Dijkstra() {
	p.Dist = make(map[State]int)
	p.Prev = make(map[State]State)
	pending := lane.NewMinPriorityQueue[State, float64]()
	pending.Push(State{p.Start, East}, 0)
	for k, v := range p.Map {
		if v != "#" {
			for i := North; i <= West; i += 1 {
				p.Dist[State{k, Dir(i)}] = 10000000000
				p.Prev[State{k, Dir(i)}] = State{}
				pending.Push(State{k, Dir(i)}, 10000000000)
			}
		}
	}
	p.Dist[State{p.Start, East}] = 0

	for pending.Size() > 0 {
		vertex, _, _ := pending.Pop()
		for _, n := range vertex.Neighbours(*p) {
			cost := 1
			if n.Direction != vertex.Direction {
				cost = 1000
			}
			alt := p.Dist[vertex] + cost
			if alt < p.Dist[n] {
				p.Prev[n] = vertex
				p.Dist[n] = alt
				pending.Push(n, float64(alt))
			}
		}
	}
}

func (p *Problem) Part1() (result int) {
	result = p.Dist[State{p.End, North}]
	for i := North; i <= West; i += 1 {
		if p.Dist[State{p.End, Dir(i)}] < result {
			result = p.Dist[State{p.End, Dir(i)}]
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	best := p.Dist[State{p.End, North}]
	var bestDirection Dir
	for i := North; i <= West; i += 1 {
		if p.Dist[State{p.End, Dir(i)}] < best {
			best = p.Dist[State{p.End, Dir(i)}]
			bestDirection = Dir(i)
		}
	}
	pending := lane.NewMinPriorityQueue[State, float64]()
	pending.Push(State{p.End, bestDirection}, 0)
	visited := make(map[State]void)
	unique := make(map[Coord]void)
	unique[p.Start] = null
	unique[p.End] = null
	for pending.Size() > 0 {
		current, _, _ := pending.Pop()
		if _, ok := visited[current]; ok {
			continue
		}
		nullState := State{}
		if p.Prev[current] == nullState {
			continue
		}
		visited[current] = null
		unique[current.Current] = null
		pending.Push(p.Prev[current], 0)
		log.Debug(fmt.Sprintf("%v: %d", current, p.Dist[current]))
		for _, n := range current.neighbours(*p, true) {
			log.Debug(fmt.Sprintf("%v %v: %d %d", current, n, p.Dist[current], p.Dist[n]))
			cost := 1
			if n.Direction != current.Direction {
				cost = 1000
			}
			if p.Dist[n]+cost == p.Dist[current] {
				log.Debug(fmt.Sprintf("Going to visit extra neighbour: %v", n))
				pending.Push(n, 0)
			}
		}
	}
	result = len(unique)
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
	p.Dijkstra()
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

func (s *State) neighbours(p Problem, reverse bool) (ns []State) {
	candidates := []State{}
	var straightCandidate State
	direction := s.Direction
	if reverse {
		direction = (s.Direction + 2) % 4
	}
	switch direction {
	case North:
		straightCandidate = State{
			Coord{s.Current.X - 1, s.Current.Y},
			s.Direction,
		}
	case South:
		straightCandidate = State{
			Coord{s.Current.X + 1, s.Current.Y},
			s.Direction,
		}
	case East:
		straightCandidate = State{
			Coord{s.Current.X, s.Current.Y + 1},
			s.Direction,
		}
	case West:
		straightCandidate = State{
			Coord{s.Current.X, s.Current.Y - 1},
			s.Direction,
		}
	}
	candidates = append(candidates, straightCandidate)
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

func (s *State) Neighbours(p Problem) (ns []State) {
	return s.neighbours(p, false)
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
