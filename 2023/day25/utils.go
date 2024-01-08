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
	Vertices map[string]*Vertex
	Edges    map[EdgeName]Edge
}

type Vertex struct {
	Name  string
	Neigh []*Vertex
}

type EdgeName struct {
	A, B string
}

type Edge struct {
	A, B *Vertex
}

func (p *Problem) EdgeExists(a, b string, ignore map[EdgeName]void) bool {
	_, ok1 := p.Edges[EdgeName{a, b}]
	_, ok2 := p.Edges[EdgeName{b, a}]
	_, ig1 := ignore[EdgeName{a, b}]
	_, ig2 := ignore[EdgeName{b, a}]
	return (ok1 || ok2) && !ig1 && !ig2
}

func (p *Problem) BFS(start, end *Vertex, ignore map[EdgeName]void) (edges map[EdgeName]void, visited map[string]void) {
	type queueItem struct {
		Current *Vertex
		Prev    *queueItem
		Depth   int
	}
	edges = make(map[EdgeName]void)
	visited = make(map[string]void)
	pending := lane.NewMinPriorityQueue[*queueItem, int]()
	pending.Push(&queueItem{Current: start}, 0)

	for pending.Size() > 0 {
		current, _, _ := pending.Pop()
		if _, ok := visited[current.Current.Name]; ok {
			continue
		}
		visited[current.Current.Name] = null
		if end != nil && current.Current.Name == end.Name {
			path := make(map[EdgeName]void)
			vertices := make(map[string]void)
			for itr := current; itr.Prev != nil; itr = itr.Prev {
				path[EdgeName{itr.Current.Name, itr.Prev.Current.Name}] = null
				vertices[itr.Current.Name] = null
			}
			return path, vertices
		}
		for _, n := range current.Current.Neigh {
			if p.EdgeExists(current.Current.Name, n.Name, ignore) {
				edges[EdgeName{current.Current.Name, n.Name}] = null
				pending.Push(&queueItem{n, current, current.Depth + 1}, current.Depth+1)
			}
		}
	}
	// Can't be here if looking for path, reset result
	if end != nil {
		edges = make(map[EdgeName]void)
		visited = make(map[string]void)
	}
	return
}

func (p *Problem) AddVertex(name string) (result *Vertex) {
	if v, ok := p.Vertices[name]; ok {
		result = v
	} else {
		result = new(Vertex)
		result.Name = name
		result.Neigh = make([]*Vertex, 0)
		p.Vertices[name] = result
	}
	log.Debug(fmt.Sprintf("Added vertex %s: %p", result.Name, result))
	return
}

func (p *Problem) AddEdge(a, b string) {
	p.Edges[EdgeName{a, b}] = Edge{
		p.Vertices[a],
		p.Vertices[b],
	}
	log.Debug(fmt.Sprintf("Added edge %s-%s: %v", a, b, p.Edges[EdgeName{a, b}]))
}

func (p *Problem) CountPaths(a, b *Vertex) (count int, used map[EdgeName]void) {
	used = make(map[EdgeName]void)
	for ; ; count += 1 {
		path, _ := p.BFS(a, b, used)
		for edge := range path {
			used[edge] = null
		}
		if len(path) == 0 {
			break
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	var src *Vertex
	for _, v := range p.Vertices {
		src = v
		break
	}
	log.Info(fmt.Sprintf("Read %d vertices and %d edges", len(p.Vertices), len(p.Edges)))
	for k, v := range p.Vertices {
		if src.Name != k {
			paths, used := p.CountPaths(src, v)
			if paths <= 3 {
				_, vertices := p.BFS(v, nil, used)
				result = len(vertices) * (len(p.Vertices) - len(vertices))
				break
			}
			log.Debug(fmt.Sprintf("%s -> %s: %v\n", src.Name, k, paths))
		}
	}
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
	p.Edges = make(map[EdgeName]Edge)
	p.Vertices = make(map[string]*Vertex)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		colon := strings.Split(l, ": ")
		a := p.AddVertex(colon[0])
		for _, n := range strings.Split(colon[1], " ") {
			tmp := p.AddVertex(n)
			a.Neigh = append(a.Neigh, tmp)
			tmp.Neigh = append(tmp.Neigh, a)
			p.AddEdge(colon[0], n)
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
