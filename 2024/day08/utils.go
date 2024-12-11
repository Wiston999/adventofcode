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
	Antenas    map[string][]Coord
	Antinodes  map[Coord]void
	MaxX, MaxY int
}

func (p *Problem) GenerateAntinodes(antinodes int) {
	p.Antinodes = make(map[Coord]void)
	for a, coords := range p.Antenas {
		for i := 0; i < len(coords); i += 1 {
			for j := i + 1; j < len(coords); j += 1 {
				current := coords[i]
				xDiff := coords[i].X - coords[j].X
				yDiff := coords[i].Y - coords[j].Y
				for k := 0; k < antinodes || antinodes == -1; k += 1 {
					current.X += xDiff
					current.Y += yDiff
					log.Debug(fmt.Sprintf(
						"[%s] Possible antinode %v for %v %v [%d %d %d %d]",
						a, current, coords[i], coords[j], xDiff, yDiff, i, j,
					))
					if current.X >= 0 && current.X <= p.MaxX && current.Y >= 0 && current.Y <= p.MaxY {
						log.Info(fmt.Sprintf("Added antinode %v for %v %v", current, coords[i], coords[j]))
						p.Antinodes[current] = null
					} else {
						break
					}
				}
				current = coords[j]
				for k := 0; k < antinodes || antinodes == -1; k += 1 {
					current.X -= xDiff
					current.Y -= yDiff
					log.Debug(fmt.Sprintf(
						"[%s] Possible antinode %v for %v %v [%d %d %d %d]",
						a, current, coords[i], coords[j], xDiff, yDiff, i, j,
					))
					if current.X >= 0 && current.X <= p.MaxX && current.Y >= 0 && current.Y <= p.MaxY {
						log.Info(fmt.Sprintf("Added antinode %v for %v %v", current, coords[i], coords[j]))
						p.Antinodes[current] = null
					} else {
						break
					}
				}
			}
		}
	}
}

func (p *Problem) Part1() (result int) {
	p.GenerateAntinodes(1)
	result = len(p.Antinodes)
	return
}

func (p *Problem) Part2() (result int) {
	p.GenerateAntinodes(-1)
	for a := range p.Antinodes {
		log.Debug(fmt.Sprintf("Antinode %v", a))
	}
	result = len(p.Antinodes)
	for _, antenas := range p.Antenas {
		for i := range antenas {
			if _, ok := p.Antinodes[antenas[i]]; !ok {
				result += 1
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
	p.Antenas = make(map[string][]Coord)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range strings.Split(strings.TrimSpace(l), "") {
			if c != "." {
				p.Antenas[c] = append(p.Antenas[c], Coord{i, j})
			}
			if j > p.MaxY {
				p.MaxY = j
			}
		}
		if i > p.MaxX {
			p.MaxX = i
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
