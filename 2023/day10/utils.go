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
	Map        map[Coord]Pipe
	Score      map[Coord]int
	Start      Coord
	MaxX, MaxY int
}

type Pipe string

const (
	Ground     = "."
	Vertical   = "|"
	Horizontal = "-"
	NorthWest  = "L"
	NorthEast  = "J"
	SouthWest  = "7"
	SouthEast  = "F"
	Start      = "S"
)

func (p *Problem) Next(c Coord) (n1, n2 Coord) {
	switch p.Map[c] {
	case Vertical:
		n1, n2 = Coord{c.X, c.Y - 1}, Coord{c.X, c.Y + 1}
	case Horizontal:
		n1, n2 = Coord{c.X - 1, c.Y}, Coord{c.X + 1, c.Y}
	case NorthWest:
		n1, n2 = Coord{c.X, c.Y - 1}, Coord{c.X + 1, c.Y}
	case NorthEast:
		n1, n2 = Coord{c.X, c.Y - 1}, Coord{c.X - 1, c.Y}
	case SouthWest:
		n1, n2 = Coord{c.X - 1, c.Y}, Coord{c.X, c.Y + 1}
	case SouthEast:
		n1, n2 = Coord{c.X + 1, c.Y}, Coord{c.X, c.Y + 1}
	}
	return
}

func (p *Problem) GuessStart() {
	top := Coord{p.Start.X, p.Start.Y - 1}
	bottom := Coord{p.Start.X, p.Start.Y + 1}
	left := Coord{p.Start.X - 1, p.Start.Y}
	right := Coord{p.Start.X + 1, p.Start.Y}
	log.Debug(fmt.Sprintf("Guessing start type: %v (%v, %v, %v, %v)", p.Start, p.Map[top], p.Map[left], p.Map[bottom], p.Map[right]))
	switch p.Map[top] {
	case Vertical, SouthWest, SouthEast:
		switch p.Map[bottom] {
		case Vertical, NorthWest, NorthEast:
			p.Map[p.Start] = Vertical
		}
		switch p.Map[left] {
		case Horizontal, NorthWest, SouthEast:
			p.Map[p.Start] = NorthEast
		}
		switch p.Map[right] {
		case Horizontal, NorthEast, SouthWest:
			p.Map[p.Start] = NorthEast
		}
	}
	switch p.Map[left] {
	case Horizontal, NorthWest, SouthEast:
		switch p.Map[right] {
		case Horizontal, NorthEast, SouthWest:
			p.Map[p.Start] = Horizontal
		}
		switch p.Map[bottom] {
		case Vertical, NorthWest, NorthEast:
			p.Map[p.Start] = SouthWest
		}
		switch p.Map[top] {
		case Vertical, NorthEast, SouthWest:
			p.Map[p.Start] = NorthEast
		}
	}
	switch p.Map[bottom] {
	case Vertical, NorthWest, NorthEast:
		switch p.Map[top] {
		case Vertical, SouthWest, SouthEast:
			p.Map[p.Start] = Vertical
		}
		switch p.Map[left] {
		case Horizontal, NorthWest, SouthEast:
			p.Map[p.Start] = SouthWest
		}
		switch p.Map[right] {
		case Horizontal, NorthEast, SouthWest:
			p.Map[p.Start] = SouthEast
		}
	}
	switch p.Map[right] {
	case Horizontal, NorthWest, SouthWest:
		switch p.Map[top] {
		case Vertical, SouthWest, SouthEast:
			p.Map[p.Start] = NorthWest
		}
		switch p.Map[bottom] {
		case Vertical, SouthWest, NorthEast:
			p.Map[p.Start] = SouthEast
		}
		switch p.Map[left] {
		case Horizontal, NorthWest, SouthEast:
			p.Map[p.Start] = Horizontal
		}
	}
	log.Debug(fmt.Sprintf("Guessed start type: %v (%v, %v, %v, %v)", p.Map[p.Start], p.Map[top], p.Map[left], p.Map[bottom], p.Map[right]))
}

func (p *Problem) Traverse(c Coord, depth int) {
	log.Debug(fmt.Sprintf("Exploring %v (%03d)", c, depth))
	if score, ok := p.Score[c]; !ok || (ok && score != 0 && score >= depth) {
		log.Debug(fmt.Sprintf("Updating %v (%03d) with value %03d", c, score, depth))
		p.Score[c] = depth
		n1, n2 := p.Next(c)
		p.Traverse(n1, depth+1)
		p.Traverse(n2, depth+1)
	}
}

func (p *Problem) Part1() (result int) {
	for _, s := range p.Score {
		if s > result {
			result = s
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	for c := range p.Map {
		if _, ok := p.Score[c]; !ok {
			crosses := 0
			for current := c; current.X <= p.MaxX && current.Y <= p.MaxY; {
				if _, found := p.Score[current]; found && p.Map[current] != SouthWest && p.Map[current] != NorthWest {
					crosses += 1
				}
				current.X += 1
				current.Y += 1
			}
			if crosses%2 == 1 {
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
	p.Map = make(map[Coord]Pipe)
	p.Score = make(map[Coord]int)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			p.Map[Coord{j, i}] = Pipe(c)
			if j > p.MaxY {
				p.MaxY = j
			}
			if Pipe(c) == Start {
				p.Start = Coord{j, i}
			}
		}
		if i > p.MaxX {
			p.MaxX = i
		}
	}
	p.GuessStart()
	p.Traverse(p.Start, 0)
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
