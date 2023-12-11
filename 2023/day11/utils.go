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

type Problem struct {
	Galaxy         Map
	EmptyX, EmptyY map[int]bool
}

type Map struct {
	Data       map[Coord]Planet
	MaxX, MaxY int
}

type Pair struct {
	C1, C2 Coord
}

func (m *Map) Manhattan(x, y Coord) (result int) {
	return int(math.Abs(float64(x.X-y.X))) + int(math.Abs(float64(x.Y-y.Y)))
}

type Planet string

const (
	Void = "."
	Full = "#"
)

func (p *Problem) StoreEmpty() {
	p.EmptyX, p.EmptyY = make(map[int]bool), make(map[int]bool)
	for i := 0; i <= p.Galaxy.MaxX; i++ {
		empty := true
		for j := 0; j <= p.Galaxy.MaxY; j++ {
			if _, found := p.Galaxy.Data[Coord{i, j}]; found {
				empty = false
				break
			}
		}
		if empty {
			log.Debug(fmt.Sprintf("Found empty column: %03d", i))
			p.EmptyX[i] = true
		}
	}
	for i := 0; i <= p.Galaxy.MaxY; i++ {
		empty := true
		for j := 0; j <= p.Galaxy.MaxX; j++ {
			if _, found := p.Galaxy.Data[Coord{j, i}]; found {
				empty = false
				break
			}
		}
		if empty {
			log.Debug(fmt.Sprintf("Found column row: %03d", i))
			p.EmptyY[i] = true
		}
	}
}

func (p *Problem) ComputeDistances(inc int) (result int) {
	visitedMap := make(map[Pair]bool)
	for c1 := range p.Galaxy.Data {
		for c2 := range p.Galaxy.Data {
			visited := false
			if _, found := visitedMap[Pair{c1, c2}]; found {
				visited = true
			}
			if _, found := visitedMap[Pair{c2, c1}]; found {
				visited = true
			}
			if c1 != c2 && !visited {
				visitedMap[Pair{c1, c2}] = true
				distance := p.Galaxy.Manhattan(c1, c2)
				log.Debug(fmt.Sprintf("Distance from %v to %v: Base %d", c1, c2, distance))
				first, last := c1.X, c2.X
				if first > last {
					first, last = last, first
				}
				for i := first; i <= last; i++ {
					if _, ok := p.EmptyX[i]; ok {
						log.Debug(fmt.Sprintf("Distance from %v to %v: Crossed empty row %d", c1, c2, i))
						distance += inc
						if inc > 1 {
							distance -= 1
						}
					}
				}
				first, last = c1.Y, c2.Y
				if first > last {
					first, last = last, first
				}
				for i := first; i <= last; i++ {
					if _, ok := p.EmptyY[i]; ok {
						log.Debug(fmt.Sprintf("Distance from %v to %v: Crossed empty column %d", c1, c2, i))
						distance += inc
						if inc > 1 {
							distance -= 1
						}
					}
				}
				log.Debug(fmt.Sprintf("Distance from %v to %v: %03d", c1, c2, distance))
				result += distance
			}
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	result = p.ComputeDistances(1)
	return
}

func (p *Problem) Part2() (result int) {
	result = p.ComputeDistances(1000000)
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
	p.Galaxy.Data = make(map[Coord]Planet)
	p.EmptyX = make(map[int]bool)
	p.EmptyY = make(map[int]bool)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			if c == '#' {
				p.Galaxy.Data[Coord{j, i}] = Planet(c)
			}
			if j > p.Galaxy.MaxX {
				p.Galaxy.MaxX = j
			}
		}
		if i > p.Galaxy.MaxY {
			p.Galaxy.MaxY = i
		}
	}
	p.StoreEmpty()
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
