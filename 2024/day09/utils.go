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
	InitialData string
	Data        []int
	Files       map[int]int
	NumFiles    int
	Frees       []int
}

func (p *Problem) Result() (result int) {
	for i, v := range p.Data {
		if v > 0 {
			result += i * v
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	p.Init()
	for i, j := len(p.Data)-1, 0; i >= 0 && j < len(p.Frees); i = i - 1 {
		if p.Data[i] == -1 {
			continue
		}
		if i < p.Frees[j] {
			break
		}
		log.Debug(fmt.Sprintf("Moving %d [%d] to %d", i, p.Data[i], p.Frees[j]))
		p.Data[p.Frees[j]] = p.Data[i]
		j += 1
		p.Data[i] = -1
	}
	log.Debug(fmt.Sprintf("Reordered files: %v", p.Data))
	result = p.Result()
	return
}

func (p *Problem) Part2() (result int) {
	p.Init()
	log.Debug(fmt.Sprintf("%#v", p))
	for i, j := len(p.Data)-1, 0; i >= 0 && j < len(p.Frees); i = i - 1 {
		if p.Data[i] < 0 {
			continue
		}
		log.Debug(fmt.Sprintf("i=%d j=%d", i, j))
		freeSize := 0
		freeStart := 0
		for k := 0; k < len(p.Data); k += 1 {
			if p.Data[k] == -1 {
				freeSize += 1
			} else {
				freeSize = 0
				freeStart = k + 1
			}
			if freeSize >= p.Files[p.Data[i]] {
				break
			}
		}
		log.Debug(fmt.Sprintf(
			"Found %d free space from %d for %d",
			freeSize,
			freeStart,
			p.Data[i],
		))
		// File can't be moved
		if p.Files[p.Data[i]] > freeSize || freeStart >= i {
			continue
		}
		for k := 0; k < freeSize; k = k + 1 {
			log.Debug(fmt.Sprintf("Moving %d [%d] to %d", i-k, p.Data[i-k], freeStart+k))
			p.Data[freeStart+k] = p.Data[i-k]
			p.Data[i-k] = -1
		}
		i -= freeSize - 1
	}
	result = p.Result()
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
	p.InitialData = string(byteData)
	return
}

func (p *Problem) Init() {
	p.NumFiles = 0
	p.Files = make(map[int]int)
	p.Data = make([]int, 0)
	p.Frees = make([]int, 0)
	index := 0
	for i, a := range strings.Split(strings.TrimSpace(p.InitialData), "") {
		n := atoi(a)
		if i%2 == 1 {
			for j := 0; j < n; j += 1 {
				p.Data = append(p.Data, -1)
				p.Frees = append(p.Frees, index)
				index += 1
			}
		} else {
			p.Files[p.NumFiles] = n
			for j := 0; j < n; j += 1 {
				p.Data = append(p.Data, p.NumFiles)
				index += 1
			}
			p.NumFiles += 1
		}
	}
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
