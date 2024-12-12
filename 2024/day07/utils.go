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

type Equation struct {
	Result   int
	Operands []int
}

func (e *Equation) Evaluate(mask string) (result int) {
	result = e.Operands[0]
	log.Debug(fmt.Sprintf("Testing equation %#v with mask %s", e, mask))
	for k := 1; k < len(e.Operands); k += 1 {
		if mask[k-1] == '0' {
			result += e.Operands[k]
		} else if mask[k-1] == '1' {
			result *= e.Operands[k]
		} else {
			result = atoi(fmt.Sprintf("%d%d", result, e.Operands[k]))
		}
	}
	return
}

func (e *Equation) Check(operands int) bool {
	for j := 0; j < int(math.Pow(float64(operands), float64(len(e.Operands)-1))); j += 1 {
		format := fmt.Sprintf("%%0%ds", len(e.Operands)-1)
		mask := fmt.Sprintf(format, strconv.FormatInt(int64(j), operands))
		tmp := e.Evaluate(mask)
		if tmp == e.Result {
			log.Info(fmt.Sprintf("Found solution for %#v with %s: %d", e, mask, tmp))
			return true
		}
	}
	return false
}

type Problem struct {
	Equations []Equation
}

func (p *Problem) Part1() (result int) {
	for _, equation := range p.Equations {
		if equation.Check(2) {
			result += equation.Result
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	for _, equation := range p.Equations {
		if equation.Check(3) {
			result += equation.Result
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
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		eq := Equation{}
		fmt.Sscanf(l, "%d: ", &eq.Result)
		for j, op := range strings.Split(strings.TrimSpace(l), " ") {
			if j > 0 {
				eq.Operands = append(eq.Operands, atoi(op))
			}
		}
		p.Equations = append(p.Equations, eq)
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