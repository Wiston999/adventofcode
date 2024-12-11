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

type Level []int

func (l Level) IsSafe() bool {
	direction := '>'
	if l[1] < l[0] {
		direction = '<'
	}
	for i := 1; i < len(l); i += 1 {
		if l[i] < l[i-1] && direction != '<' {
			return false
		}
		if l[i] > l[i-1] && direction != '>' {
			return false
		}
		if direction == '>' && (l[i]-l[i-1] < 1 || l[i]-l[i-1] > 3) {
			log.Debug(fmt.Sprintf("%d vs %d is unsafe by %c", l[i-1], l[i], direction))
			return false
		}
		if direction == '<' && (l[i-1]-l[i] < 1 || l[i-1]-l[i] > 3) {
			log.Debug(fmt.Sprintf("%d vs %d is unsafe by %c", l[i-1], l[i], direction))
			return false
		}
		log.Debug(fmt.Sprintf("%d vs %d is safe by %c", l[i-1], l[i], direction))
	}
	return true
}

func (l *Level) Mutate(index int) (result Level) {
	for i, element := range *l {
		if i != index {
			result = append(result, element)
		}
	}
	return
}

type Problem struct {
	Levels []Level
}

func (p *Problem) Part1() (result int) {
	for i, l := range p.Levels {
		if l.IsSafe() {
			result += 1
			log.Debug(fmt.Sprintf("Level [%03d] %v is safe", i, l))
		} else {
			log.Debug(fmt.Sprintf("Level [%03d] %v is unsafe", i, l))
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	for i, l := range p.Levels {
		if l.IsSafe() {
			result += 1
			log.Debug(fmt.Sprintf("Level [%03d] %v is safe", i, l))
		} else {
			for j := 0; j < len(l); j++ {
				if l.Mutate(j).IsSafe() {
					result += 1
					log.Debug(fmt.Sprintf("Level [%03d] %v is safe mutating %d", i, l, j))
					break
				}
			}
			log.Debug(fmt.Sprintf("Level [%03d] %v is unsafe", i, l))
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
		level := Level{}
		for _, a := range strings.Split(strings.TrimSpace(l), " ") {
			level = append(level, atoi(a))
		}
		p.Levels = append(p.Levels, level)
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
