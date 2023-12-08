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
	Steps []Step
	Nodes map[string]Node
}

type Step string

const (
	Right = "R"
	Left  = "L"
)

type Node struct {
	Left  string
	Right string
}

func (p *Problem) FindGoal(current, goal string, threshold int) (result int) {
	steps := len(p.Steps)
	for ; current != goal && result < threshold; result += 1 {
		log.Debug(fmt.Sprintf("Step %03d at %s", result, current))
		switch p.Steps[result%steps] {
		case Left:
			current = p.Nodes[current].Left
		case Right:
			current = p.Nodes[current].Right
		}
	}

	return
}

func (p *Problem) Part1() (result int) {
	return p.FindGoal("AAA", "ZZZ", 100000000)
}

func (p *Problem) Part2() (result int) {
	current := []string{}
	goals := []string{}
	for k := range p.Nodes {
		if k[2] == 'A' {
			current = append(current, k)
		}
		if k[2] == 'Z' {
			goals = append(goals, k)
		}
	}
	log.Info(fmt.Sprintf("There are %d starting nodes", len(current)))
	result = 1
	threshold := 100000

	for _, g := range goals {
		for _, c := range current {
			steps := p.FindGoal(c, g, threshold)
			if steps == threshold {
				log.Debug(fmt.Sprintf("There is no path from %s to %s", c, g))
			} else {
				log.Info(fmt.Sprintf("From %s to %s takes %d steps", c, g, steps))
				result = LCM(result, steps)
				break
			}
		}
	}
	return
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func NewProblem(ctx *cli.Context) (p *Problem, err error) {
	input := ctx.String("input")
	byteData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	strData := string(byteData)
	p = new(Problem)
	p.Nodes = make(map[string]Node)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		if i == 0 {
			for _, c := range l {
				p.Steps = append(p.Steps, Step(c))
			}
			continue
		}
		if len(l) > 0 {
			var nodeName string
			n := Node{}
			parts := strings.Split(strings.Split(l, " = ")[1], ", ")
			fmt.Sscanf(l, "%s = ", &nodeName)
			n.Left = strings.Trim(parts[0], "(")
			n.Right = strings.Trim(parts[1], ")")

			p.Nodes[nodeName] = n
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
