package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/oleiade/lane/v2"

	"github.com/urfave/cli/v2"
)

type void struct{}

var null void

type Problem struct {
	Connections map[string]map[string]void
}

func (p *Problem) Part1() (result int) {
	groups := make(map[string]void)
	for c1, g1 := range p.Connections {
		for c2 := range g1 {
			for c3 := range g1 {
				if _, ok := p.Connections[c2][c3]; ok && c2 != c3 {
					if c1[0] == 't' || c2[0] == 't' || c3[0] == 't' {
						sorted := []string{c1, c2, c3}
						sort.Strings(sorted)
						groups[strings.Join(sorted, ",")] = null
					}
				}
			}
		}
	}
	result = len(groups)
	return
}

func (p *Problem) BuildCluster(current string, cluster map[string]void) {
	for member := range cluster {
		if _, ok := p.Connections[member][current]; !ok {
			return
		}
	}
	cluster[current] = null
	for conn := range p.Connections[current] {
		p.BuildCluster(conn, cluster)
	}
}

func (p *Problem) Part2() (result string) {
	maximum := 0
	for c1 := range p.Connections {
		c := make(map[string]void)
		p.BuildCluster(c1, c)
		if len(c) > maximum {
			maximum = len(c)
			keys := make([]string, 0)
			for k := range c {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			result = strings.Join(keys, ",")
		}
		log.Debug(fmt.Sprintf("Cluster of %v: %v", c1, c))
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
	p.Connections = make(map[string]map[string]void)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		parts := strings.Split(l, "-")
		if _, ok := p.Connections[parts[0]]; !ok {
			p.Connections[parts[0]] = make(map[string]void)
		}
		if _, ok := p.Connections[parts[1]]; !ok {
			p.Connections[parts[1]] = make(map[string]void)
		}
		p.Connections[parts[1]][parts[0]] = null
		p.Connections[parts[0]][parts[1]] = null
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
