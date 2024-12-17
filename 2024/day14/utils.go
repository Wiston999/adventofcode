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

type Robot struct {
	Start, P, V Coord
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (r *Robot) Move(seconds, MaxX, MaxY int) (result Coord) {
	result.X = mod(r.Start.X+r.V.X*seconds, MaxX)
	result.Y = mod(r.Start.Y+r.V.Y*seconds, MaxY)
	return
}

type Problem struct {
	Robots     []Robot
	Map        map[Coord][]Robot
	MaxX, MaxY int
}

func (p *Problem) Score() (result int) {
	c1, c2, c3, c4 := 0, 0, 0, 0
	for i := 0; i <= p.MaxY; i += 1 {
		for j := 0; j <= p.MaxX; j += 1 {
			if robots, ok := p.Map[Coord{j, i}]; ok {
				// C1
				if i < p.MaxY/2 && j < p.MaxX/2 {
					log.Debug(fmt.Sprintf("Added {%d %d} to cuadrant 1: %#v", j, i, robots))
					c1 += len(robots)
				}
				//C2
				if i > p.MaxY/2 && j < p.MaxX/2 {
					log.Debug(fmt.Sprintf("Added {%d %d} to cuadrant 2: %#v", j, i, robots))
					c2 += len(robots)
				}
				//C3
				if i < p.MaxY/2 && j > p.MaxX/2 {
					log.Debug(fmt.Sprintf("Added {%d %d} to cuadrant 3: %#v", j, i, robots))
					c3 += len(robots)
				}
				//C4
				if i > p.MaxY/2 && j > p.MaxX/2 {
					log.Debug(fmt.Sprintf("Added {%d %d} to cuadrant 4: %#v", j, i, robots))
					c4 += len(robots)
				}
			}
		}
	}
	log.Debug(fmt.Sprintf("%d %d %d %d", c1, c2, c3, c4))
	result = c1 * c2 * c3 * c4
	return
}

func (p *Problem) Part1() (result int) {
	p.Map = make(map[Coord][]Robot)
	for i, r := range p.Robots {
		position := r.Move(100, p.MaxX, p.MaxY)
		p.Map[position] = append(p.Map[position], r)
		log.Debug(fmt.Sprintf("[%02d] Robot %#v is at %v after 100 seconds", i, r, position))
	}
	result = p.Score()
	return
}

func (p *Problem) Part2() (result int) {
	for j := 1; j < 10000; j += 1 {
		p.Map = make(map[Coord][]Robot)
		for _, r := range p.Robots {
			position := r.Move(j, p.MaxX, p.MaxY)
			p.Map[position] = append(p.Map[position], r)
		}
		for k, v := range p.Map {
			if len(v) >= 1 {
				seemsTree := true
				for x := 1; x < 5; x += 1 {
					for y := k.Y - x; y <= k.Y+x; y += 1 {
						log.Debug(fmt.Sprintf("[%02d] %v {%d %d}: %v", j, k, x, y, p.Map[Coord{k.X + x, y}]))
						if robots, ok := p.Map[Coord{k.X + x, y}]; !ok || len(robots) == 0 {
							seemsTree = false
							break
						}
					}
					if !seemsTree {
						break
					}
				}
				if seemsTree {
					result = j
					break
				}
			}
		}
		if result != 0 {
			break
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
	p.MaxX = 101
	p.MaxY = 103
	if strings.Contains(input, "test") {
		p.MaxX = 11
		p.MaxY = 7
	}
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		r := Robot{}
		fmt.Sscanf(l, "p=%d,%d v=%d,%d", &r.Start.X, &r.Start.Y, &r.V.X, &r.V.Y)
		p.Robots = append(p.Robots, r)
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
