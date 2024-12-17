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

type Instruction struct {
	Operation Op
	Operand   int
}

func (i *Instruction) Value(p Problem) (result int) {
	switch i.Operand {
	case 0, 1, 2, 3:
		result = i.Operand
	case 4:
		result = p.A
	case 5:
		result = p.B
	case 6:
		result = p.C
	}
	return
}

type Problem struct {
	Instructions []Instruction
	A, B, C, PC  int
	Out          []string
	Raw          []string
}

type Op int

const (
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func (p *Problem) Operate(i Instruction) {
	p.PC += 1
	switch i.Operation {
	case adv:
		p.A = p.A / int(math.Pow(2, float64(i.Value(*p))))
	case bxl:
		p.B = p.B ^ i.Operand
	case bst:
		p.B = i.Value(*p) % 8
	case jnz:
		if p.A != 0 {
			p.PC = i.Operand
		}
	case bxc:
		p.B = p.B ^ p.C
	case out:
		p.Out = append(p.Out, fmt.Sprintf("%d", i.Value(*p)%8))
	case bdv:
		p.B = p.A / int(math.Pow(2, float64(i.Value(*p))))
	case cdv:
		p.C = p.A / int(math.Pow(2, float64(i.Value(*p))))
	}
}

func (p *Problem) Compute(a int) (result []string) {
	p.Out = make([]string, 0)
	p.PC = 0
	p.A, p.B, p.C = a, 0, 0
	for p.PC < len(p.Instructions) {
		p.Operate(p.Instructions[p.PC])
	}
	return p.Out
}

func (p *Problem) Part1() (result string) {
	result = strings.Join(p.Compute(p.A), ",")
	return
}

func (p *Problem) Part2() (result int) {
	for i := len(p.Raw) - 1; i >= 0; i -= 1 {
		result <<= 3
		for {
			log.Debug(fmt.Sprintf("Comparing [%d] %v to %v", result, p.Out, p.Raw[i:]))
			p.Compute(result)
			if slices.Equal(p.Out, p.Raw[i:]) {
				break
			}
			result += 1
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
		parts := strings.Split(l, " ")
		if strings.Contains(l, "Register A") {
			p.A = atoi(parts[2])
		}
		if strings.Contains(l, "Register B") {
			p.B = atoi(parts[2])
		}
		if strings.Contains(l, "Register C") {
			p.C = atoi(parts[2])
		}
		if strings.Contains(l, "Program") {
			ops := strings.Split(parts[1], ",")
			p.Raw = ops
			for i := 0; i < len(ops); i += 2 {
				instruction := Instruction{Op(atoi(ops[i])), atoi(ops[i+1])}
				p.Instructions = append(p.Instructions, instruction)
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
