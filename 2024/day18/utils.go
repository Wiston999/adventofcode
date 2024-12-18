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
	Bytes      []Coord
	Map        map[Coord]string
	Start, End Coord
	MaxX, MaxY int
}

func (p *Problem) Print(path []State) (result string) {
	pathCoords := make(map[Coord]void)
	for _, c := range path {
		pathCoords[c.Current] = null
	}
	for i := 0; i <= p.MaxX; i += 1 {
		for j := 0; j <= p.MaxY; j += 1 {
			if p.Map[Coord{i, j}] == "#" {
				result += "#"
			} else if _, ok := pathCoords[Coord{i, j}]; ok {
				result += "O"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return
}

func (p *Problem) Part1() (result int) {
	p.Map = make(map[Coord]string)
	for i := 0; i < 1024; i += 1 {
		p.Map[p.Bytes[i]] = "#"
	}
	pathFinder := PathFinder{
		P:     *p,
		Start: State{Coord{0, 0}},
		Goal: func(s State) bool {
			return s.Current.X == p.MaxX && s.Current.Y == p.MaxY
		},
		Cost: func(n, c State) float64 {
			return 1
		},
		Heuristic: func(s State) float64 {
			return s.Current.Manhattan(Coord{p.MaxX, p.MaxY})
		},
	}
	path, score := pathFinder.Search()
	path = append(path, State{Coord{}})
	log.Info(fmt.Sprintf("Found path with score %f: %v", score, path))
	fmt.Println(p.Print(path))
	result = int(score)
	return
}

func (p *Problem) Part2() (result string) {
	for i := 1024; i < len(p.Bytes); i += 1 {
		p.Map[p.Bytes[i]] = "#"
		pathFinder := PathFinder{
			P:     *p,
			Start: State{Coord{0, 0}},
			Goal: func(s State) bool {
				return s.Current.X == p.MaxX && s.Current.Y == p.MaxY
			},
			Cost: func(n, c State) float64 {
				return 1
			},
			Heuristic: func(s State) float64 {
				return s.Current.Manhattan(Coord{p.MaxX, p.MaxY})
			},
		}
		path, _ := pathFinder.Search()
		if len(path) == 0 {
			log.Info(fmt.Sprintf("Found no path after adding %v", p.Bytes[i]))
			result = fmt.Sprintf("%d,%d", p.Bytes[i].Y, p.Bytes[i].X)
			break
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
	p.MaxX, p.MaxY = 70, 70
	if strings.Contains(input, "test") {
		p.MaxX, p.MaxY = 6, 6
	}
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		c := Coord{}
		fmt.Sscanf(l, "%d,%d", &c.Y, &c.X)
		p.Bytes = append(p.Bytes, c)
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
	Current Coord
}

func (s *State) Neighbours(p Problem) (ns []State) {
	candidates := []Coord{
		{s.Current.X + 1, s.Current.Y},
		{s.Current.X - 1, s.Current.Y},
		{s.Current.X, s.Current.Y + 1},
		{s.Current.X, s.Current.Y - 1},
	}
	for _, c := range candidates {
		if c.X > p.MaxX || c.Y > p.MaxY || c.X < 0 || c.Y < 0 {
			continue
		}
		if v := p.Map[c]; v == "#" {
			continue
		}
		ns = append(ns, State{c})
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
