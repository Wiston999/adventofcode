package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/pforemski/gouda/interpolate"
	"github.com/pforemski/gouda/point"

	"github.com/oleiade/lane/v2"

	"github.com/urfave/cli/v2"
)

type Problem struct {
	Maze       map[Coord]Dot
	MaxX, MaxY int
	Start      Coord
}

type Position struct {
	C     Coord
	Depth int
}

type Dot string

const (
	Start  = "S"
	Garden = "."
	Rock   = "#"
)

func (p *Problem) Print(visited map[Coord]int) (result string) {
	for i := 0; i <= p.MaxX; i++ {
		for j := 0; j <= p.MaxY; j++ {
			if v, ok := visited[Coord{j, i}]; ok {
				result += fmt.Sprintf("%03d ", v)
			} else {
				result += string(p.Maze[Coord{j, i}]) + "   "
			}
		}
		result += "\n"
	}
	return
}

func mod(n, m int) int {
	return (n%m + m) % m
}

func (p *Problem) Visit(limit int) (visited map[Coord]int) {
	visited = make(map[Coord]int)
	pending := lane.NewMinPriorityQueue[Position, int]()
	pending.Push(Position{p.Start, 0}, 0)
	for pending.Size() > 0 {
		current, _, _ := pending.Pop()
		if _, ok := visited[current.C]; ok || current.Depth > limit {
			continue
		}
		visited[current.C] = current.Depth

		candidates := []Coord{
			Coord{current.C.X + 1, current.C.Y},
			Coord{current.C.X - 1, current.C.Y},
			Coord{current.C.X, current.C.Y + 1},
			Coord{current.C.X, current.C.Y - 1},
		}

		for _, c := range candidates {
			if v, ok := p.Maze[Coord{mod(c.X, p.MaxX+1), mod(c.Y, p.MaxY+1)}]; ok && v != Rock {
				pending.Push(Position{c, current.Depth + 1}, current.Depth+1)
			}
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	visited := p.Visit(64)
	for _, v := range visited {
		if v%2 == 0 {
			result += 1
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	x0, x1, x2 := 0, 0, 0
	visited := p.Visit(65)
	for _, v := range visited {
		if v%2 == 1 {
			x0 += 1
		}
	}
	visited = p.Visit(65 + 131)
	for _, v := range visited {
		if v%2 == 0 {
			x1 += 1
		}
	}
	visited = p.Visit(65 + 131*2)
	for _, v := range visited {
		if v%2 == 1 {
			x2 += 1
		}
	}
	log.Debug(fmt.Sprintf("Interpolating for a, b, c = %d, %d, %d", x0, x1, x2))
	points := point.NewZeros(3, 2) // 3 points in 2 axes
	points[0].V[0] = 65
	points[0].V[1] = float64(x0)
	points[1].V[0] = 65 + 131
	points[1].V[1] = float64(x1)
	points[2].V[0] = 65 + 2*131
	points[2].V[1] = float64(x2)
	lagrange, _ := interpolate.NewLagrange(points)
	result = int(lagrange.Interpolate(26501365))
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
	p.Maze = make(map[Coord]Dot)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			p.Maze[Coord{j, i}] = Dot(c)
			if p.Maze[Coord{j, i}] == Start {
				p.Start = Coord{j, i}
			}
			if i > p.MaxY {
				p.MaxY = i
			}
			if j > p.MaxX {
				p.MaxX = j
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
