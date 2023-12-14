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
	Rocks Map
}

type Map struct {
	Temp       map[Coord]Spot
	Data       map[Coord]Spot
	MaxX, MaxY int
}

func (m *Map) Add(x, y int, c Spot) {
	if x > m.MaxX {
		m.MaxX = x
	}
	if y > m.MaxY {
		m.MaxY = y
	}
	m.Data[Coord{x, y}] = c
}

func (m *Map) RollNorth() (result bool) {
	for i := 0; i <= m.MaxX; i++ {
		for j := 0; j <= m.MaxY; j++ {
			if p, ok := m.Temp[Coord{i, j}]; ok && p == Round {
				if n, nOk := m.Temp[Coord{i - 1, j}]; nOk && n == Empty {
					m.Temp[Coord{i, j}], m.Temp[Coord{i - 1, j}] = Empty, m.Temp[Coord{i, j}]
					result = true
				}
			}
		}
	}
	return
}

func (m *Map) RollSouth() (result bool) {
	for i := m.MaxX; i >= 0; i-- {
		for j := m.MaxY; j >= 0; j-- {
			if p, ok := m.Temp[Coord{i, j}]; ok && p == Round {
				if n, nOk := m.Temp[Coord{i + 1, j}]; nOk && n == Empty {
					m.Temp[Coord{i, j}], m.Temp[Coord{i + 1, j}] = Empty, m.Temp[Coord{i, j}]
					result = true
				}
			}
		}
	}
	return
}

func (m *Map) RollEast() (result bool) {
	for i := 0; i <= m.MaxX; i++ {
		for j := 0; j <= m.MaxY; j++ {
			if p, ok := m.Temp[Coord{i, j}]; ok && p == Round {
				if n, nOk := m.Temp[Coord{i, j - 1}]; nOk && n == Empty {
					m.Temp[Coord{i, j}], m.Temp[Coord{i, j - 1}] = Empty, m.Temp[Coord{i, j}]
					result = true
				}
			}
		}
	}
	return
}

func (m *Map) RollWest() (result bool) {
	for i := m.MaxX; i >= 0; i-- {
		for j := m.MaxY; j >= 0; j-- {
			if p, ok := m.Temp[Coord{i, j}]; ok && p == Round {
				if n, nOk := m.Temp[Coord{i, j + 1}]; nOk && n == Empty {
					m.Temp[Coord{i, j}], m.Temp[Coord{i, j + 1}] = Empty, m.Temp[Coord{i, j}]
					result = true
				}
			}
		}
	}
	return
}

func (m *Map) Print() (result string) {
	for i := 0; i <= m.MaxX; i++ {
		for j := 0; j <= m.MaxY; j++ {
			result += string(m.Temp[Coord{i, j}])
		}
		result += "\n"
	}
	return
}

func (m *Map) Score() (result int) {
	for i := 0; i <= m.MaxX; i++ {
		for j := 0; j <= m.MaxY; j++ {
			if r, ok := m.Temp[Coord{i, j}]; ok && r == Round {
				result += m.MaxX - i + 1
			}
		}
	}
	return
}

func (m *Map) Reset() {
	for c, d := range m.Data {
		m.Temp[c] = d
	}
}

type Spot rune

const (
	Empty  = '.'
	Round  = 'O'
	Square = '#'
)

func (p *Problem) Part1() (result int) {
	p.Rocks.Reset()
	for p.Rocks.RollNorth() {
	}
	result = p.Rocks.Score()
	return
}

func (p *Problem) Part2() (result int) {
	memory := make(map[string]int)
	p.Rocks.Reset()
	fmt.Println(p.Rocks.Print())
	for i := 0; i < 1000000000; i++ {
		if cycle_count, ok := memory[p.Rocks.Print()]; ok {
			period := i - cycle_count
			for j := 0; j < (1000000000-i)%period; j++ {
				log.Debug(fmt.Sprintf("[%03d] Cycle found at %d to %d", j, i, cycle_count))
				for p.Rocks.RollNorth() {
				}
				for p.Rocks.RollEast() {
				}
				for p.Rocks.RollSouth() {
				}
				for p.Rocks.RollWest() {
				}
			}
			break
		}
		memory[p.Rocks.Print()] = i
		for p.Rocks.RollNorth() {
		}
		for p.Rocks.RollEast() {
		}
		for p.Rocks.RollSouth() {
		}
		for p.Rocks.RollWest() {
		}
	}
	result = p.Rocks.Score()
	fmt.Println(p.Rocks.Print())
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
	p.Rocks.Data = make(map[Coord]Spot)
	p.Rocks.Temp = make(map[Coord]Spot)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			p.Rocks.Add(i, j, Spot(c))
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
