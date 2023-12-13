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
	Patterns []Map
}

type Map struct {
	Data       map[Coord]Spot
	Rows       []string
	Columns    []string
	MaxX, MaxY int
}

type Spot rune
type ReflectionType int

const (
	Ash    = '.'
	Mirror = '#'
)

const (
	Vertical = iota
	Horizontal
)

func FindMirror(data []string) (found bool, index int) {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == data[i+1] {
			fullReflect := true
			var head, tail int
			for head, tail := i, i+1; head >= 0 && tail < len(data); head, tail = head-1, tail+1 {
				if data[head] != data[tail] {
					fullReflect = false
					break
				}
			}
			if head != 0 && tail != len(data)-1 {
				fullReflect = false
			}
			if fullReflect {
				return true, i + 1
			}
		}
	}
	return
}

func (m *Map) FindReflection() (reflection ReflectionType, index int) {
	if found, index := FindMirror(m.Rows); found {
		return Horizontal, index
	}
	if found, index := FindMirror(m.Columns); found {
		return Vertical, index
	}
	return
}

func (m *Map) Flip(i, j int) {
	if m.Rows[i][j] == '.' {
		m.Rows[i] = m.Rows[i][:j] + "#" + m.Rows[i][j+1:]
		m.Columns[j] = m.Columns[j][:i] + "#" + m.Columns[j][i+1:]
	} else {
		m.Rows[i] = m.Rows[i][:j] + "." + m.Rows[i][j+1:]
		m.Columns[j] = m.Columns[j][:i] + "." + m.Columns[j][i+1:]
	}
}

func (p *Problem) Part1() (result int) {
	for i, m := range p.Patterns {
		reflection, index := m.FindReflection()
		log.Debug(fmt.Sprintf("[%02d] Found reflection %v @ %d", i, reflection, index))
		if reflection == Vertical {
			result += index
		} else {
			result += index * 100
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	for i, m := range p.Patterns {
		for j, r := range m.Rows {
			flipped := false
			for k := range r {
				m.Flip(j, k)
				reflection, index := m.FindReflection()
				if index != 0 {
					log.Debug(fmt.Sprintf("[%02d] Found reflection %v @ %d %v (%d,%d)", i, reflection, index, m, j, k))
					if reflection == Vertical {
						result += index
					} else {
						result += index * 100
					}
					flipped = true
				}
				m.Flip(j, k)
				if flipped {
					break
				}
			}
			if flipped {
				break
			}
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
	tmp := Map{Data: make(map[Coord]Spot)}
	x, y := 0, 0
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		if len(l) <= 2 {
			p.Patterns = append(p.Patterns, tmp)
			tmp = Map{Data: make(map[Coord]Spot)}
			x, y = 0, 0
			continue
		}
		x = 0
		for _, c := range l {
			tmp.Data[Coord{x, y}] = Spot(c)
			if x > tmp.MaxX {
				tmp.MaxX = x
			}
			if y > tmp.MaxY {
				tmp.MaxY = y
			}
			x += 1
		}
		tmp.Rows = append(tmp.Rows, l)
		y += 1
	}
	p.Patterns = append(p.Patterns, tmp)
	log.Info(fmt.Sprintf("Found patterns %02d", len(p.Patterns)))
	for k, pattern := range p.Patterns {
		for i := 0; i < len(pattern.Rows[0]); i++ {
			tmpColumn := ""
			for _, r := range pattern.Rows {
				tmpColumn += string(r[i])
			}
			p.Patterns[k].Columns = append(p.Patterns[k].Columns, tmpColumn)
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
