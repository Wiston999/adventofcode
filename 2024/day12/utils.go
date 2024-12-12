package main

import (
	"fmt"
	"maps"
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

type Vertex struct {
	A, B Coord
}

type Problem struct {
	Map        map[Coord]string
	Computed   map[Coord]void
	MaxX, MaxY int
}

func (p *Problem) CountVertices(current Coord, visited map[Coord]void) (result int) {
	if _, ok := visited[current]; ok {
		return
	}
	visited[current] = null
	p.Computed[current] = null
	neighbors := []Coord{
		{current.X - 1, current.Y},
		{current.X + 1, current.Y},
		{current.X, current.Y - 1},
		{current.X, current.Y + 1},
	}
	diff := 0
	for _, n := range neighbors {
		if p.Map[n] == p.Map[current] {
			result += p.CountVertices(n, visited)
		} else {
			diff += 1
		}
	}
	switch diff {
	case 1:
		c := Coord{current.X - 1, current.Y}
		if _, ok := visited[c]; !ok && p.Map[current] != p.Map[c] {
			visited[c] = null
			if p.Map[current] == p.Map[Coord{current.X - 1, current.Y + 1}] {
				result += 1
			}
			if p.Map[current] == p.Map[Coord{current.X - 1, current.Y - 1}] {
				result += 1
			}
		}
		c = Coord{current.X + 1, current.Y}
		if _, ok := visited[c]; !ok && p.Map[current] != p.Map[c] {
			visited[c] = null
			if p.Map[current] == p.Map[Coord{current.X + 1, current.Y + 1}] {
				result += 1
			}
			if p.Map[current] == p.Map[Coord{current.X + 1, current.Y - 1}] {
				result += 1
			}
		}
		c = Coord{current.X, current.Y - 1}
		if _, ok := visited[c]; !ok && p.Map[current] != p.Map[c] {
			visited[c] = null
			if p.Map[current] == p.Map[Coord{current.X - 1, current.Y - 1}] {
				result += 1
			}
			if p.Map[current] == p.Map[Coord{current.X + 1, current.Y - 1}] {
				result += 1
			}
		}
		c = Coord{current.X, current.Y + 1}
		if _, ok := visited[c]; !ok && p.Map[current] != p.Map[c] {
			visited[c] = null
			if p.Map[current] == p.Map[Coord{current.X - 1, current.Y + 1}] {
				result += 1
			}
			if p.Map[current] == p.Map[Coord{current.X + 1, current.Y + 1}] {
				result += 1
			}
		}
	case 2:
		c1, c2 := Coord{current.X - 1, current.Y}, Coord{current.X, current.Y - 1}
		c3, c4 := Coord{current.X + 1, current.Y}, Coord{current.X, current.Y + 1}
		_, ok1 := visited[c1]
		_, ok2 := visited[c2]
		_, ok3 := visited[c3]
		_, ok4 := visited[c4]
		if !ok1 && !ok2 && p.Map[current] != p.Map[c1] && p.Map[current] != p.Map[c2] {
			visited[c1], visited[c2] = null, null
			if p.Map[current] != p.Map[Coord{current.X + 1, current.Y + 1}] {
				result += 2
			} else {
				result += 1
			}
		} else if !ok1 && !ok4 && p.Map[current] != p.Map[c1] && p.Map[current] != p.Map[c4] {
			visited[c1], visited[c4] = null, null
			if p.Map[current] != p.Map[Coord{current.X + 1, current.Y - 1}] {
				result += 2
			} else {
				result += 1
			}
		} else if !ok3 && !ok2 && p.Map[current] != p.Map[c3] && p.Map[current] != p.Map[c2] {
			visited[c3], visited[c2] = null, null
			if p.Map[current] != p.Map[Coord{current.X - 1, current.Y + 1}] {
				result += 2
			} else {
				result += 1
			}
		} else if !ok3 && !ok4 && p.Map[current] != p.Map[c3] && p.Map[current] != p.Map[c4] {
			visited[c3], visited[c4] = null, null
			if p.Map[current] != p.Map[Coord{current.X - 1, current.Y - 1}] {
				result += 2
			} else {
				result += 1
			}
		} else {
			if _, ok := visited[Coord{current.X - 1, current.Y - 1}]; !ok && p.Map[current] == p.Map[Coord{current.X - 1, current.Y - 1}] {
				visited[Coord{current.X - 1, current.Y - 1}] = null
				result += 1
			}
			if _, ok := visited[Coord{current.X - 1, current.Y + 1}]; !ok && p.Map[current] == p.Map[Coord{current.X - 1, current.Y + 1}] {
				visited[Coord{current.X - 1, current.Y - 1}] = null
				result += 1
			}
			if _, ok := visited[Coord{current.X + 1, current.Y - 1}]; !ok && p.Map[current] == p.Map[Coord{current.X + 1, current.Y - 1}] {
				visited[Coord{current.X - 1, current.Y - 1}] = null
				result += 1
			}
			if _, ok := visited[Coord{current.X + 1, current.Y + 1}]; !ok && p.Map[current] == p.Map[Coord{current.X + 1, current.Y + 1}] {
				visited[Coord{current.X - 1, current.Y - 1}] = null
				result += 1
			}
		}
	case 3:
		result += 2
	case 4:
		result += 4
	}

	switch diff {
	case 1:
		c := Coord{current.X - 1, current.Y}
		if _, ok := visited[c]; !ok && p.Map[current] != p.Map[c] {
			visited[c] = null
			if p.Map[current] == p.Map[Coord{current.X - 1, current.Y + 1}] {
				result += 1
			}
			if p.Map[current] == p.Map[Coord{current.X - 1, current.Y - 1}] {
				result += 1
			}
		}
		c = Coord{current.X + 1, current.Y}
		if _, ok := visited[c]; !ok && p.Map[current] != p.Map[c] {
			visited[c] = null
			if p.Map[current] == p.Map[Coord{current.X + 1, current.Y + 1}] {
				result += 1
			}
			if p.Map[current] == p.Map[Coord{current.X + 1, current.Y - 1}] {
				result += 1
			}
		}
		c = Coord{current.X, current.Y - 1}
		if _, ok := visited[c]; !ok && p.Map[current] != p.Map[c] {
			visited[c] = null
			if p.Map[current] == p.Map[Coord{current.X - 1, current.Y - 1}] {
				result += 1
			}
			if p.Map[current] == p.Map[Coord{current.X + 1, current.Y - 1}] {
				result += 1
			}
		}
		c = Coord{current.X, current.Y + 1}
		if _, ok := visited[c]; !ok && p.Map[current] != p.Map[c] {
			visited[c] = null
			if p.Map[current] == p.Map[Coord{current.X - 1, current.Y + 1}] {
				result += 1
			}
			if p.Map[current] == p.Map[Coord{current.X + 1, current.Y + 1}] {
				result += 1
			}
		}
	case 2:
		c1, c2 := Coord{current.X - 1, current.Y}, Coord{current.X, current.Y - 1}
		c3, c4 := Coord{current.X + 1, current.Y}, Coord{current.X, current.Y + 1}
		_, ok1 := visited[c1]
		_, ok2 := visited[c2]
		_, ok3 := visited[c3]
		_, ok4 := visited[c4]
		if !ok1 && !ok2 && p.Map[current] != p.Map[c1] && p.Map[current] != p.Map[c2] {
			visited[c1], visited[c2] = null, null
			if p.Map[current] != p.Map[Coord{current.X + 1, current.Y + 1}] {
				result += 2
			} else {
				result += 1
			}
		} else if !ok1 && !ok4 && p.Map[current] != p.Map[c1] && p.Map[current] != p.Map[c4] {
			visited[c1], visited[c4] = null, null
			if p.Map[current] != p.Map[Coord{current.X + 1, current.Y - 1}] {
				result += 2
			} else {
				result += 1
			}
		} else if !ok3 && !ok2 && p.Map[current] != p.Map[c3] && p.Map[current] != p.Map[c2] {
			visited[c3], visited[c2] = null, null
			if p.Map[current] != p.Map[Coord{current.X - 1, current.Y + 1}] {
				result += 2
			} else {
				result += 1
			}
		} else if !ok3 && !ok4 && p.Map[current] != p.Map[c3] && p.Map[current] != p.Map[c4] {
			visited[c3], visited[c4] = null, null
			if p.Map[current] != p.Map[Coord{current.X - 1, current.Y - 1}] {
				result += 2
			} else {
				result += 1
			}
		} else {
			if _, ok := visited[Coord{current.X - 1, current.Y - 1}]; !ok && p.Map[current] == p.Map[Coord{current.X - 1, current.Y - 1}] {
				visited[Coord{current.X - 1, current.Y - 1}] = null
				result += 1
			}
			if _, ok := visited[Coord{current.X - 1, current.Y + 1}]; !ok && p.Map[current] == p.Map[Coord{current.X - 1, current.Y + 1}] {
				visited[Coord{current.X - 1, current.Y - 1}] = null
				result += 1
			}
			if _, ok := visited[Coord{current.X + 1, current.Y - 1}]; !ok && p.Map[current] == p.Map[Coord{current.X + 1, current.Y - 1}] {
				visited[Coord{current.X - 1, current.Y - 1}] = null
				result += 1
			}
			if _, ok := visited[Coord{current.X + 1, current.Y + 1}]; !ok && p.Map[current] == p.Map[Coord{current.X + 1, current.Y + 1}] {
				visited[Coord{current.X - 1, current.Y - 1}] = null
				result += 1
			}
		}
	case 3:
		result += 2
	case 4:
		result += 4
	}
	return
}

func (p *Problem) FenceValue(start Coord) (area, perimeter int, vertices map[Vertex]int) {
	vertices = make(map[Vertex]int)
	if _, ok := p.Computed[start]; ok {
		return
	}
	current := p.Map[start]
	p.Computed[start] = null
	neighbors := []Coord{
		{start.X - 1, start.Y},
		{start.X + 1, start.Y},
		{start.X, start.Y - 1},
		{start.X, start.Y + 1},
	}
	area += 1
	diff := 0
	for _, n := range neighbors {
		if p.Map[n] == current {
			a, p, v := p.FenceValue(n)
			area += a
			perimeter += p
			maps.Copy(vertices, v)
		} else {
			diff += 1
			perimeter += 1
		}
	}
	tl := p.Map[Coord{start.X - 1, start.Y - 1}]
	tt := p.Map[Coord{start.X - 1, start.Y}]
	tr := p.Map[Coord{start.X - 1, start.Y + 1}]
	l := p.Map[Coord{start.X, start.Y - 1}]
	r := p.Map[Coord{start.X, start.Y + 1}]
	bl := p.Map[Coord{start.X + 1, start.Y - 1}]
	bb := p.Map[Coord{start.X + 1, start.Y}]
	br := p.Map[Coord{start.X + 1, start.Y + 1}]
	log.Debug(fmt.Sprintf("%v\n%v %v %v\n%v %v %v\n%v %v %v", start, tl, tt, tr, l, current, r, bl, bb, br))
	if (tt != current && l != current) || (tt != l && (tl == current)) {
		log.Debug(fmt.Sprintf("Detected vertex: %v %v", start, Coord{start.X - 1, start.Y - 1}))
		vertices[Vertex{start, Coord{start.X - 1, start.Y - 1}}] = 1
		if tt != current && l != current && tl == current {
			vertices[Vertex{start, Coord{start.X - 1, start.Y - 1}}] += 1
		}
	}
	if (tt != current && r != current) || (tt != r && (tr == current)) {
		log.Debug(fmt.Sprintf("Detected vertex: %v %v", start, Coord{start.X - 1, start.Y + 1}))
		vertices[Vertex{start, Coord{start.X - 1, start.Y + 1}}] = 1
		if tt != current && r != current && tr == current {
			vertices[Vertex{start, Coord{start.X - 1, start.Y + 1}}] += 1
		}
	}
	if (bb != current && l != current) || (bb != l && (bl == current)) {
		log.Debug(fmt.Sprintf("Detected vertex: %v %v", start, Coord{start.X + 1, start.Y - 1}))
		vertices[Vertex{start, Coord{start.X + 1, start.Y - 1}}] = 1
		if bb != current && l != current && bl == current {
			vertices[Vertex{start, Coord{start.X + 1, start.Y - 1}}] += 1
		}
	}
	if (bb != current && r != current) || (bb != r && (br == current)) {
		log.Debug(fmt.Sprintf("Detected vertex: %v %v", start, Coord{start.X + 1, start.Y + 1}))
		vertices[Vertex{start, Coord{start.X + 1, start.Y + 1}}] = 1
		if bb != current && r != current && br == current {
			vertices[Vertex{start, Coord{start.X + 1, start.Y + 1}}] += 1
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	p.Computed = make(map[Coord]void)
	for i := 0; i <= p.MaxX; i += 1 {
		for j := 0; j <= p.MaxY; j += 1 {
			area, perimeter, _ := p.FenceValue(Coord{i, j})
			log.Debug(fmt.Sprintf("Area and perimeter of %v: %d %d", Coord{i, j}, area, perimeter))
			result += area * perimeter
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	p.Computed = make(map[Coord]void)
	for i := 0; i <= p.MaxX; i += 1 {
		for j := 0; j <= p.MaxY; j += 1 {
			area, _, vertices := p.FenceValue(Coord{i, j})
			unique := make(map[Vertex]int)
			for v := range vertices {
				_, ok1 := unique[v]
				_, ok2 := unique[Vertex{v.B, v.A}]
				if !ok1 && !ok2 {
					if value, opposite := vertices[Vertex{v.B, v.A}]; opposite {
						unique[v] = value
					} else {
						unique[v] = 1
					}
				}
			}
			log.Debug(fmt.Sprintf("Vertices of %v: %v [%v]", Coord{i, j}, vertices, unique))
			log.Debug(fmt.Sprintf("Area and vertices of %v: %d %d", Coord{i, j}, area, len(unique)))
			for _, v := range unique {
				result += v * area
			}
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
			p.Map[Coord{i, j}] = c
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
