package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/oleiade/lane/v2"

	"github.com/urfave/cli/v2"
)

type void struct{}

var null void

type Problem struct {
	Map                         map[Coord]string
	MaxX, MaxY                  int
	InitialGuard, Guard         Coord
	InitialDirection, Direction string
}

func (p *Problem) Print(guard Coord, direction string) (result string) {
	for i := 0; i <= p.MaxX; i++ {
		for j := 0; j <= p.MaxY; j++ {
			if guard.X == i && guard.Y == j {
				result += direction
			} else {
				result += p.Map[Coord{i, j}]
			}
		}
		result += "\n"
	}
	return
}

func (p *Problem) Walk() (visited []Coord, repeat bool) {
	p.Guard = p.InitialGuard
	p.Direction = p.InitialDirection
	hits := make(map[Coord][]string)
	for {
		// Out of map
		if _, ok := p.Map[p.Guard]; !ok {
			break
		}
		log.Debug(fmt.Sprintf("Visiting %v: %d", p.Guard, len(visited)))
		visited = append(visited, p.Guard)
		next := p.Guard
		switch p.Direction {
		case "^":
			next.X -= 1
		case "v":
			next.X += 1
		case "<":
			next.Y -= 1
		case ">":
			next.Y += 1
		}
		if p.Map[next] != "#" {
			p.Guard = next
		} else {
			// Already hit that obstacle in the same direction, we are going to loop
			if slices.Contains(hits[p.Guard], p.Direction) {
				repeat = true
				break
			}
			hits[p.Guard] = append(hits[p.Guard], p.Direction)
			switch p.Direction {
			case "^":
				p.Direction = ">"
			case "v":
				p.Direction = "<"
			case "<":
				p.Direction = "^"
			case ">":
				p.Direction = "v"
			}
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	path, _ := p.Walk()
	visited := make(map[Coord]void)
	for _, step := range path {
		visited[step] = null
	}
	result = len(visited)
	return
}

func (p *Problem) Part2() (result int) {
	path, _ := p.Walk()
	tested := make(map[Coord]void)
	for _, c := range path {
		candidates := []Coord{
			{c.X - 1, c.Y},
			{c.X, c.Y - 1},
			{c.X, c.Y + 1},
			{c.X + 1, c.Y},
		}
		nc := make([]Coord, 0)
		for _, c := range candidates {
			if _, ok := p.Map[c]; ok {
				if _, ok := tested[c]; !ok {
					tested[c] = null
					nc = append(nc, c)
				}
			}
		}
		for _, c := range nc {
			log.Debug(fmt.Sprintf("Testing placing obstacle at %v", c))
			previous := p.Map[c]
			p.Map[c] = "#"
			if _, loop := p.Walk(); loop {
				log.Info(fmt.Sprintf("Detected loop placing obstacle at %v", c))
				result += 1
			}
			p.Map[c] = previous
		}
	}
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
			if c != "." && c != "#" {
				p.InitialDirection = c
				p.InitialGuard = Coord{i, j}
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
	ID int
}

func (s *State) Neighbours(p Problem) (ns []State) {

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
