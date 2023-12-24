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
	Instructions []Instruction
}

type Instruction struct {
	Direction string
	Length    int
	Color     string
}

type void struct{}

const (
	Empty = "."
	Edge  = "#"
)

const (
	R = "R"
	L = "L"
	U = "U"
	D = "D"
)

func (p *Problem) Dig() (result int) {
	current := Coord{0, 0}
	for _, order := range p.Instructions {
		next := current
		log.Debug(fmt.Sprintf("Instruction %v at %v", order, current))
		switch order.Direction {
		case R:
			next.X += order.Length
		case L:
			next.X -= order.Length
		case U:
			next.Y -= order.Length
		case D:
			next.Y += order.Length
		}
		result += ((next.Y*current.X - next.X*current.Y) + order.Length)
		current = next
	}
	result = result/2 + 1
	return
}

func (p *Problem) Part1() (result int) {
	result = p.Dig()
	return
}

func (p *Problem) Part2() (result int) {
	np := Problem{}
	for _, i := range p.Instructions {
		ni := Instruction{}
		tmp, _ := strconv.ParseInt(string(i.Color[:5]), 16, 64)
		ni.Length = int(tmp)
		switch i.Color[5] {
		case '0':
			ni.Direction = R
		case '1':
			ni.Direction = D
		case '2':
			ni.Direction = L
		case '3':
			ni.Direction = U
		}
		log.Debug(fmt.Sprintf("New instruction: %v", ni))
		np.Instructions = append(np.Instructions, ni)
	}

	result = np.Dig()

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
		instruction := Instruction{}
		fmt.Sscanf(l, "%s %d (#%s)", &instruction.Direction, &instruction.Length, &instruction.Color)
		p.Instructions = append(p.Instructions, instruction)
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
