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

type Problem struct {
	Cache  Cache
	Stones []int
}

type Cache map[CacheIndex]int

type CacheIndex struct {
	Stone, Stage int
}

func (p *Problem) Evolve(stone, stage, maxStage int) (result int) {
	if stage > maxStage {
		return 1
	}
	log.Debug(fmt.Sprintf(
		"At stage %d with stone %d",
		stage, stone,
	))
	if value, ok := p.Cache[CacheIndex{stone, stage}]; ok {
		return value
	}
	if stone == 0 {
		value := p.Evolve(1, stage+1, maxStage)
		p.Cache[CacheIndex{1, stage + 1}] = value
		result += value
	} else {
		strStone := fmt.Sprintf("%d", stone)
		if len(strStone)%2 == 0 {
			left := p.Evolve(atoi(strStone[:len(strStone)/2]), stage+1, maxStage)
			right := p.Evolve(atoi(strStone[len(strStone)/2:]), stage+1, maxStage)
			p.Cache[CacheIndex{atoi(strStone[:len(strStone)/2]), stage + 1}] = left
			p.Cache[CacheIndex{atoi(strStone[len(strStone)/2:]), stage + 1}] = right
			result += left + right
		} else {
			value := p.Evolve(stone*2024, stage+1, maxStage)
			result += value
			p.Cache[CacheIndex{stone * 2024, stage + 1}] = value
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	p.Cache = make(Cache)
	for i, stone := range p.Stones {
		evolved := p.Evolve(stone, 1, 25)
		result += evolved
		log.Info(fmt.Sprintf("[%d] Stone %d evolved to %v stones", i, stone, evolved))
	}
	return
}

func (p *Problem) Part2() (result int) {
	p.Cache = make(Cache)
	for i, stone := range p.Stones {
		evolved := p.Evolve(stone, 1, 75)
		result += evolved
		log.Info(fmt.Sprintf("[%d] Stone %d evolved to %v stones", i, stone, evolved))
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
	for i, l := range strings.Split(strings.TrimSpace(strData), " ") {
		log.Debug(fmt.Sprintf("Parsing number %03d: %s", i, l))
		p.Stones = append(p.Stones, atoi(l))
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
