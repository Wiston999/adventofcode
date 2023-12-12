package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/oleiade/lane/v2"

	"github.com/urfave/cli/v2"
)

type Problem struct {
	Pipes []Pipe
}

type Pipe struct {
	Data     string
	Unknowns []int
	Readings []int
	Memory   map[string]int
}

type Spot string

const (
	Broken      = "#"
	Operational = "."
	Unknown     = "?"
)

func (p *Pipe) Check(p1, p2 string) (result bool) {
	min := len(p1)
	if len(p2) < len(p1) {
		min = len(p2)
	}
	for i := 0; i < min; i++ {
		if p1[i] != p2[i] && p1[i] != '?' {
			return false
		}
	}
	return true
}

func (p *Pipe) GetCombinations(partial string, size int, readings []int) (result int) {
	if r, ok := p.Memory[partial+fmt.Sprintf("-%d-%v", size, readings)]; ok {
		return r
	}
	if len(readings) == 0 {
		for _, c := range partial {
			if c == '#' {
				return 0
			}
		}
		return 1
	}
	reading := readings[0]
	pending := readings[1:]
	left := sum(pending) + len(pending)

	for i := 0; i <= size-left-reading; i++ {
		candidate := strings.Repeat(".", i) + strings.Repeat("#", reading) + "."
		if p.Check(partial, candidate) {
			if len(candidate) < len(partial) {
				result += p.GetCombinations(partial[len(candidate):], size-reading-i-1, pending)
			} else if len(pending) == 0 {
				result += 1
			}
		}
	}
	p.Memory[partial+fmt.Sprintf("-%d-%v", size, readings)] = result
	return
}

func sum(numbers []int) (result int) {
	for _, n := range numbers {
		result += n
	}
	return
}

func (p *Pipe) CountCombinations() (result int) {
	result = p.GetCombinations(p.Data, len(p.Data), p.Readings)
	return
}

func (p *Problem) Part1() (result int) {
	for i, pipe := range p.Pipes {
		c := pipe.CountCombinations()
		log.Info(fmt.Sprintf("[%03d] Found %02d in pipe %v", i, c, pipe))
		result += c
	}
	return
}

func (p *Problem) Part2() (result int) {
	results := make(chan int, len(p.Pipes))
	wg := sync.WaitGroup{}

	for i, pipe := range p.Pipes {
		newPipe := Pipe{Memory: make(map[string]int)}
		newPipe.Data = pipe.Data
		newPipe.Readings = append(newPipe.Readings, pipe.Readings...)
		for j := 0; j < 4; j++ {
			newPipe.Data += "?" + pipe.Data
			newPipe.Readings = append(newPipe.Readings, pipe.Readings...)
		}
		wg.Add(1)
		go func(i int, p Pipe) {
			c := newPipe.CountCombinations()
			log.Info(fmt.Sprintf("[%03d] Found %02d in a pipe", i, c))
			results <- c
			wg.Done()
		}(i, newPipe)
	}

	wg.Wait()
	close(results)
	i := 0
	for c := range results {
		result += c
		i += 1
		log.Info(fmt.Sprintf("[%03d] Read result %02d", i, c))
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
	for _, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		splitted := strings.Split(l, " ")
		pipe := Pipe{Data: splitted[0], Memory: make(map[string]int)}
		for _, r := range strings.Split(splitted[1], ",") {
			pipe.Readings = append(pipe.Readings, atoi(r))
		}
		p.Pipes = append(p.Pipes, pipe)
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
		log.Fatal(fmt.Sprintf("Unknown log level %s", level))
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
