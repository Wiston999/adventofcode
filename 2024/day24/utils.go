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

type Gate struct {
	A, B, Output string
	Operation    string
}

const (
	AND = "AND"
	OR  = "OR"
	XOR = "XOR"
)

type Problem struct {
	Wires map[string]int
	Gates []Gate
	MaxZ  int
}

func (p *Problem) Part1() (result int) {
	processed := make(map[Gate]void)
	for len(processed) != len(p.Gates) {
		for _, g := range p.Gates {
			if _, ok := processed[g]; ok {
				continue
			}
			vA, okA := p.Wires[g.A]
			vB, okB := p.Wires[g.B]
			if okA && okB {
				switch g.Operation {
				case AND:
					p.Wires[g.Output] = vA * vB
				case OR:
					p.Wires[g.Output] = 0
					if vA == 1 || vB == 1 {
						p.Wires[g.Output] = 1
					}
				case XOR:
					p.Wires[g.Output] = 0
					if vA != vB {
						p.Wires[g.Output] = 1
					}
				}
				processed[g] = null
			}
		}
	}
	for i := p.MaxZ; i >= 1; i-- {
		result += p.Wires[fmt.Sprintf("z%02d", i)]
		result = result << 1
	}
	return
}

func (p *Problem) Part2() (result string) {
	wrong := make(map[string]void)
	for _, g := range p.Gates {
		if g.Output[0] == 'z' && g.Operation != XOR && atoi(g.Output[1:]) != p.MaxZ {
			wrong[g.Output] = null
		}
		if g.Operation == XOR {
			if g.A[0] != 'x' && g.A[0] != 'y' && g.A[0] != 'z' {
				if g.B[0] != 'x' && g.B[0] != 'y' && g.B[0] != 'z' {
					if g.Output[0] != 'x' && g.Output[0] != 'y' && g.Output[0] != 'z' {
						wrong[g.Output] = null
					}
				}
			}
			for _, other := range p.Gates {
				if other.Operation == OR && (g.Output == other.A || g.Output == other.B) {
					wrong[g.Output] = null
				}
			}
		}
		if g.Operation == AND && g.A != "x00" && g.B != "x00" {
			for _, other := range p.Gates {
				if other.Operation != OR && (g.Output == other.A || g.Output == other.B) {
					wrong[g.Output] = null
				}
			}
		}
	}
	wrongStr := []string{}
	for k := range wrong {
		wrongStr = append(wrongStr, k)
	}
	sort.Strings(wrongStr)
	result = strings.Join(wrongStr, ",")
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
	p.Wires = make(map[string]int)
	strData := string(byteData)
	wires := true
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		if len(l) == 0 {
			wires = false
			continue
		}
		if wires {
			parts := strings.Split(l, ": ")
			p.Wires[parts[0]] = atoi(parts[1])
		} else {
			g := Gate{}
			fmt.Sscanf(l, "%s %s %s -> %s", &g.A, &g.Operation, &g.B, &g.Output)
			p.Gates = append(p.Gates, g)
			if strings.HasPrefix(g.Output, "z") {
				z := atoi(g.Output[1:])
				if z > p.MaxZ {
					p.MaxZ = z
				}
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
