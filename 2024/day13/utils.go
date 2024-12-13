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

type Button struct {
	X, Y int
}

type Claw struct {
	A, B, Prize Button
}

func (c *Claw) Score() (result int, found bool) {
	result = 101 * 4
	for i := 1; i <= 100; i += 1 {
		for j := 1; j <= 100; j += 1 {
			if (c.A.X*i+c.B.X*j) == c.Prize.X && (c.A.Y*i+c.B.Y*j) == c.Prize.Y {
				if (i*3 + j) < result {
					result = i*3 + j
					found = true
				}
			}
		}
	}
	return
}

func (c *Claw) SmartScore() (result int, found bool) {
	B := (c.Prize.Y*c.A.X - c.Prize.X*c.A.Y) / (c.A.X*c.B.Y - c.A.Y*c.B.X)
	A := (c.Prize.X*c.B.Y - c.Prize.Y*c.B.X) / (c.A.X*c.B.Y - c.A.Y*c.B.X)

	log.Debug(fmt.Sprintf("%d %d", A, B))
	if c.Prize.X == c.A.X*A+c.B.X*B && c.Prize.Y == c.A.Y*A+c.B.Y*B {
		result = 3*A + B
		found = true
	}
	return
}

func GCD(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	} else {
		gcd, x, y := GCD(b%a, a)
		return gcd, y - (b/a)*x, x
	}
}

type Problem struct {
	Claws []Claw
}

func (p *Problem) Part1() (result int) {
	for _, c := range p.Claws {
		score, found := c.SmartScore()
		if found {
			result += score
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	for _, c := range p.Claws {
		c.Prize.X += 10000000000000
		c.Prize.Y += 10000000000000
		score, found := c.SmartScore()
		if found {
			result += score
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
	current := Claw{}
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		if len(l) > 0 {
			x, y := 0, 0
			var button rune
			if _, err := fmt.Sscanf(l, "Button %c: X+%d, Y+%d", &button, &x, &y); err == nil {
				if button == 'A' {
					current.A = Button{x, y}
				} else {
					current.B = Button{x, y}
				}
			} else {
				fmt.Sscanf(l, "Prize: X=%d, Y=%d", &x, &y)
				current.Prize.X = x
				current.Prize.Y = y
				p.Claws = append(p.Claws, current)
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
