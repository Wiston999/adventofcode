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
	Games []Game
}

type Game struct {
	ID   int
	Sets []Set
}

type Set struct {
	Plays []Play
}

type Play struct {
	Count int
	Color string
}

func (g *Game) Power() (result int) {
	maxRed, maxGreen, maxBlue := 0, 0, 0
	for _, s := range g.Sets {
		for _, p := range s.Plays {
			if p.Color == "red" && p.Count > maxRed {
				maxRed = p.Count
			}
			if p.Color == "blue" && p.Count > maxBlue {
				maxBlue = p.Count
			}
			if p.Color == "green" && p.Count > maxGreen {
				maxGreen = p.Count
			}
		}
	}
	return maxRed * maxGreen * maxBlue
}

func (p *Problem) Part1() (result int) {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	for _, game := range p.Games {
		possible := true
		for _, s := range game.Sets {
			for _, p := range s.Plays {
				if p.Color == "red" && p.Count > maxRed {
					possible = false
					break
				}
				if p.Color == "blue" && p.Count > maxBlue {
					possible = false
					break
				}
				if p.Color == "green" && p.Count > maxGreen {
					possible = false
					break
				}
			}
		}
		if possible {
			result += game.ID
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	for _, g := range p.Games {
		result += g.Power()
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
	strData := string(byteData)
	p = new(Problem)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		g := Game{}
		fmt.Sscanf(l, "Game %d:", &g.ID)
		sets := strings.Split(l, ": ")[1]
		log.Debug(fmt.Sprintf("Parsing sets %03d: %s", i, sets))
		for _, s := range strings.Split(sets, ";") {
			set := Set{}
			for _, p := range strings.Split(s, ",") {
				play := Play{}
				fmt.Sscanf(p, "%d %s", &play.Count, &play.Color)
				set.Plays = append(set.Plays, play)
			}
			g.Sets = append(g.Sets, set)
		}
		p.Games = append(p.Games, g)
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
