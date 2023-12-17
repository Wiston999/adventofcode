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
	Map        map[Coord]int
	MaxX, MaxY int
}

func (p *Problem) Part1() (result int) {
	pathFinder := PathFinder{
		P:     *p,
		Start: State{Coord{0, 0}, 1, Right, 0},
		Goal: func(s State) bool {
			return s.Current.X == p.MaxX && s.Current.Y == p.MaxY
		},
		Cost: func(n, c State) float64 {
			return float64(p.Map[n.Current])
		},
		Heuristic: func(s State) float64 {
			return s.Current.Manhattan(Coord{p.MaxX, p.MaxY})
		},
		Neighbours: func(s State, p Problem) []State {
			return s.NeighboursP1(p)
		},
	}
	path, score := pathFinder.Search()
	log.Debug(fmt.Sprintf("Found Path %v with score %f", path, score))
	result = int(score)
	return
}

func (p *Problem) Part2() (result int) {
	pathFinder := PathFinder{
		P:     *p,
		Start: State{Coord{0, 0}, 1, Right, 1},
		Goal: func(s State) bool {
			return s.Current.X == p.MaxX && s.Current.Y == p.MaxY
		},
		Cost: func(n, c State) float64 {
			return float64(p.Map[n.Current])
		},
		Heuristic: func(s State) float64 {
			return s.Current.Manhattan(Coord{p.MaxX, p.MaxY})
		},
		Neighbours: func(s State, p Problem) []State {
			return s.NeighboursP2(p)
		},
	}
	path, score := pathFinder.Search()
	log.Debug(fmt.Sprintf("Found Path %v with score %f", path, score))
	result = int(score)
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
	p.Map = make(map[Coord]int)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			p.Map[Coord{j, i}] = atoi(string(c))
			if i > p.MaxY {
				p.MaxY = i
			}
			if j > p.MaxX {
				p.MaxX = j
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
	Current   Coord
	Straight  int
	Direction Dir
	Inertia   int
}

type Dir string

const (
	Up    = "^"
	Down  = "v"
	Right = ">"
	Left  = "<"
)

func (s *State) NeighboursP1(p Problem) (ns []State) {
	candidates := []State{}
	switch s.Direction {
	case Up:
		candidates = append(candidates, State{
			Current:   Coord{s.Current.X - 1, s.Current.Y},
			Straight:  1,
			Direction: Left,
		})
		candidates = append(candidates, State{
			Current:   Coord{s.Current.X + 1, s.Current.Y},
			Straight:  1,
			Direction: Right,
		})
		if s.Straight < 3 {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X, s.Current.Y - 1},
				Straight:  s.Straight + 1,
				Direction: Up,
			})
		}
	case Down:
		candidates = append(candidates, State{
			Current:   Coord{s.Current.X - 1, s.Current.Y},
			Straight:  1,
			Direction: Left,
		})
		candidates = append(candidates, State{
			Current:   Coord{s.Current.X + 1, s.Current.Y},
			Straight:  1,
			Direction: Right,
		})
		if s.Straight < 3 {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X, s.Current.Y + 1},
				Straight:  s.Straight + 1,
				Direction: Down,
			})
		}
	case Left:
		candidates = append(candidates, State{
			Current:   Coord{s.Current.X, s.Current.Y - 1},
			Straight:  1,
			Direction: Up,
		})
		candidates = append(candidates, State{
			Current:   Coord{s.Current.X, s.Current.Y + 1},
			Straight:  1,
			Direction: Down,
		})
		if s.Straight < 3 {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X - 1, s.Current.Y},
				Straight:  s.Straight + 1,
				Direction: Left,
			})
		}
	case Right:
		candidates = append(candidates, State{
			Current:   Coord{s.Current.X, s.Current.Y - 1},
			Straight:  1,
			Direction: Up,
		})
		candidates = append(candidates, State{
			Current:   Coord{s.Current.X, s.Current.Y + 1},
			Straight:  1,
			Direction: Down,
		})
		if s.Straight < 3 {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X + 1, s.Current.Y},
				Straight:  s.Straight + 1,
				Direction: Right,
			})
		}
	}
	for _, c := range candidates {
		if _, ok := p.Map[c.Current]; ok {
			ns = append(ns, c)
		}
	}
	return
}

func (s *State) NeighboursP2(p Problem) (ns []State) {
	candidates := []State{}
	switch s.Direction {
	case Up:
		if s.Inertia < 4 {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X, s.Current.Y - 1},
				Straight:  s.Straight + 1,
				Direction: Up,
				Inertia:   s.Inertia + 1,
			})
		} else {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X - 1, s.Current.Y},
				Straight:  1,
				Direction: Left,
				Inertia:   1,
			})
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X + 1, s.Current.Y},
				Straight:  1,
				Direction: Right,
				Inertia:   1,
			})
			if s.Straight < 10 {
				candidates = append(candidates, State{
					Current:   Coord{s.Current.X, s.Current.Y - 1},
					Straight:  s.Straight + 1,
					Direction: Up,
					Inertia:   s.Inertia + 1,
				})
			}
		}
	case Down:
		if s.Inertia < 4 {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X, s.Current.Y + 1},
				Straight:  s.Straight + 1,
				Direction: Down,
				Inertia:   s.Inertia + 1,
			})
		} else {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X - 1, s.Current.Y},
				Straight:  1,
				Direction: Left,
				Inertia:   1,
			})
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X + 1, s.Current.Y},
				Straight:  1,
				Direction: Right,
				Inertia:   1,
			})
			if s.Straight < 10 {
				candidates = append(candidates, State{
					Current:   Coord{s.Current.X, s.Current.Y + 1},
					Straight:  s.Straight + 1,
					Direction: Down,
					Inertia:   s.Inertia + 1,
				})
			}
		}
	case Left:
		if s.Inertia < 4 {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X - 1, s.Current.Y},
				Straight:  s.Straight + 1,
				Direction: Left,
				Inertia:   s.Inertia + 1,
			})
		} else {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X, s.Current.Y - 1},
				Straight:  1,
				Direction: Up,
				Inertia:   1,
			})
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X, s.Current.Y + 1},
				Straight:  1,
				Direction: Down,
				Inertia:   1,
			})
			if s.Straight < 10 {
				candidates = append(candidates, State{
					Current:   Coord{s.Current.X - 1, s.Current.Y},
					Straight:  s.Straight + 1,
					Direction: Left,
					Inertia:   s.Inertia + 1,
				})
			}
		}
	case Right:
		if s.Inertia < 4 {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X + 1, s.Current.Y},
				Straight:  s.Straight + 1,
				Direction: Right,
				Inertia:   s.Inertia + 1,
			})
		} else {
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X, s.Current.Y - 1},
				Straight:  1,
				Direction: Up,
				Inertia:   1,
			})
			candidates = append(candidates, State{
				Current:   Coord{s.Current.X, s.Current.Y + 1},
				Straight:  1,
				Direction: Down,
				Inertia:   1,
			})
			if s.Straight < 10 {
				candidates = append(candidates, State{
					Current:   Coord{s.Current.X + 1, s.Current.Y},
					Straight:  s.Straight + 1,
					Direction: Right,
					Inertia:   s.Inertia + 1,
				})
			}
		}
	}
	for _, c := range candidates {
		if _, ok := p.Map[c.Current]; ok {
			ns = append(ns, c)
		}
	}
	return
}

type PathFinder struct {
	P          Problem
	Start      State
	Goal       func(State) bool
	Cost       func(State, State) float64
	Heuristic  func(State) float64
	Neighbours func(State, Problem) []State
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

		for _, n := range p.Neighbours(current, p.P) {
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
