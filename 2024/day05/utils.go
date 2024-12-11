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
	Rules     map[int][]int
	Printings [][]int
}

func (p *Problem) IsValidPrinting(printing []int) int {
	for j, page := range printing {
		for k := j; k < len(printing); k += 1 {
			log.Debug(fmt.Sprintf(
				"Checking printing %v for page %d at element %d (%d) with rules %v",
				printing,
				page,
				k,
				printing[k],
				p.Rules[page],
			))
			if slices.Contains(p.Rules[page], printing[k]) {
				log.Debug(fmt.Sprintf("Discarded %v as %d [%d] is in %v", printing, printing[k], k, p.Rules[page]))
				return k
			}
		}
	}
	return -1
}

func (p *Problem) Part1() (result int) {
	for _, printing := range p.Printings {
		if p.IsValidPrinting(printing) == -1 {
			log.Info(fmt.Sprintf("Printing %v is valid: %d", printing, printing[len(printing)/2]))
			result += printing[len(printing)/2]
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	for _, printing := range p.Printings {
		if p.IsValidPrinting(printing) != -1 {
			newPrinting := make([]int, 0)
			for i := 0; i < len(printing); i++ {
				newPrinting = append(newPrinting, printing[i])
			}
			for {
				invalid := p.IsValidPrinting(newPrinting)
				if invalid == -1 {
					break
				}
				newPrinting[invalid], newPrinting[invalid-1] = newPrinting[invalid-1], newPrinting[invalid]
			}
			log.Info(fmt.Sprintf("Printing %v is now valid: %v", printing, newPrinting))
			result += newPrinting[len(printing)/2]
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
	p.Rules = make(map[int][]int)
	strData := string(byteData)
	initial := true
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		if len(l) == 0 {
			initial = false
			continue
		}
		if initial {
			x, y := 0, 0
			fmt.Sscanf(l, "%d|%d", &x, &y)
			p.Rules[y] = append(p.Rules[y], x)
		} else {
			printing := make([]int, 0)
			for _, a := range strings.Split(strings.TrimSpace(l), ",") {
				printing = append(printing, atoi(a))
			}
			p.Printings = append(p.Printings, printing)
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
