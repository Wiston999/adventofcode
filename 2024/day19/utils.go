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

type Pair struct {
	A, B string
}

type Problem struct {
	Towels  []string
	Designs []string
	Cache   map[string]int
}

func (p *Problem) IsPossible(design string) (result bool) {
	if len(design) == 0 {
		return true
	}
	for _, t := range p.Towels {
		if strip, could := strings.CutPrefix(design, t); could {
			log.Debug(fmt.Sprintf("Checking if %s is possible after removing %s [%s]", design, t, strip))
			result = result || p.IsPossible(strip)
		}
		if result {
			return
		}
	}
	return
}

func (p *Problem) CountPossible(design string) (result int) {
	if len(design) == 0 {
		return 1
	}
	for _, t := range p.Towels {
		if strip, could := strings.CutPrefix(design, t); could {
			log.Debug(fmt.Sprintf("Checking if %s is possible after removing %s [%s]", design, t, strip))
			if ways, cached := p.Cache[strip]; cached {
				result += ways
				continue
			}
			p.Cache[strip] = p.CountPossible(strip)
			result += p.Cache[strip]
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	for i, d := range p.Designs {
		if p.IsPossible(d) {
			result += 1
			log.Debug(fmt.Sprintf("[%03d] Design %s is possible", i, d))
		} else {
			log.Debug(fmt.Sprintf("[%03d] Design %s is not possible", i, d))
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	p.Cache = make(map[string]int)
	for i, d := range p.Designs {
		ways := p.CountPossible(d)
		log.Debug(fmt.Sprintf("[%03d] Design %s is possible in %d ways", i, d, ways))
		result += ways
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
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		if i == 0 {
			for _, c := range strings.Split(strings.TrimSpace(l), ", ") {
				p.Towels = append(p.Towels, c)
			}
		}
		if i > 1 {
			p.Designs = append(p.Designs, l)
		}
	}
	slices.SortFunc(p.Towels, func(a, b string) int {
		if len(a) < len(b) {
			return 1
		}
		return -1
	})
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
