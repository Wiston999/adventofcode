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
	Cave       map[Coord]Tile
	MaxX, MaxY int
}

type Tile struct {
	Type   rune
	Energy int
}

type Step struct {
	P Coord
	D Direction
}

type Direction string
type void struct{}

const (
	Right = ">"
	Left  = "<"
	Up    = "^"
	Down  = "v"
)

func (p *Problem) PrintCave(current Coord) (result string) {
	for i := 0; i <= p.MaxY; i++ {
		for j := 0; j <= p.MaxX; j++ {
			if j == current.X && i == current.Y {
				result += "C"
			} else if p.Cave[Coord{j, i}].Energy > 0 {
				result += "#"
			} else {
				result += string(p.Cave[Coord{j, i}].Type)
			}
		}
		result += "\n"
	}
	return
}

func (p *Problem) Next(s Step) (n1, n2 *Step) {
	tile := p.Cave[s.P]
	tile.Energy += 1
	p.Cave[s.P] = tile
	switch tile.Type {
	case '|':
		switch s.D {
		case Right, Left:
			n1 = &Step{Coord{s.P.X, s.P.Y + 1}, Down}
			n2 = &Step{Coord{s.P.X, s.P.Y - 1}, Up}
		case Up:
			n1 = &Step{Coord{s.P.X, s.P.Y - 1}, Up}
		case Down:
			n1 = &Step{Coord{s.P.X, s.P.Y + 1}, Down}
		}
	case '\\':
		switch s.D {
		case Up:
			n1 = &Step{Coord{s.P.X - 1, s.P.Y}, Left}
		case Down:
			n1 = &Step{Coord{s.P.X + 1, s.P.Y}, Right}
		case Right:
			n1 = &Step{Coord{s.P.X, s.P.Y + 1}, Down}
		case Left:
			n1 = &Step{Coord{s.P.X, s.P.Y - 1}, Up}
		}
	case '/':
		switch s.D {
		case Up:
			n1 = &Step{Coord{s.P.X + 1, s.P.Y}, Right}
		case Down:
			n1 = &Step{Coord{s.P.X - 1, s.P.Y}, Left}
		case Right:
			n1 = &Step{Coord{s.P.X, s.P.Y - 1}, Up}
		case Left:
			n1 = &Step{Coord{s.P.X, s.P.Y + 1}, Down}
		}
	case '-':
		switch s.D {
		case Up, Down:
			n1 = &Step{Coord{s.P.X + 1, s.P.Y}, Right}
			n2 = &Step{Coord{s.P.X - 1, s.P.Y}, Left}
		case Right:
			n1 = &Step{Coord{s.P.X + 1, s.P.Y}, Right}
		case Left:
			n1 = &Step{Coord{s.P.X - 1, s.P.Y}, Left}
		}
	case '.':
		switch s.D {
		case Up:
			n1 = &Step{Coord{s.P.X, s.P.Y - 1}, Up}
		case Down:
			n1 = &Step{Coord{s.P.X, s.P.Y + 1}, Down}
		case Right:
			n1 = &Step{Coord{s.P.X + 1, s.P.Y}, Right}
		case Left:
			n1 = &Step{Coord{s.P.X - 1, s.P.Y}, Left}
		}
	}
	return
}

func (p *Problem) Ray(s Step) (result int) {
	visited := make(map[Step]void)
	pending := lane.NewMinPriorityQueue[Step, float64]()
	pending.Push(s, 0)
	for pending.Size() > 0 {
		toVisit, _, _ := pending.Pop()
		n1, n2 := p.Next(toVisit)
		log.Debug(fmt.Sprintf("Visiting %v, next values %v %v", toVisit, n1, n2))
		if n1 != nil {
			_, inCave := p.Cave[n1.P]
			if _, ok := visited[*n1]; !ok && inCave {
				pending.Push(*n1, 0)
				visited[*n1] = void{}
			}
		}
		if n2 != nil {
			_, inCave := p.Cave[n2.P]
			if _, ok := visited[*n2]; !ok && inCave {
				pending.Push(*n2, 0)
				visited[*n2] = void{}
			}
		}
	}
	return p.Energy()
}

func (p *Problem) Energy() (result int) {
	for _, v := range p.Cave {
		if v.Energy > 0 {
			result += 1
		}
	}
	return
}

func (p *Problem) Reset() {
	for c, t := range p.Cave {
		t.Energy = 0
		p.Cave[c] = t
	}
}

func (p *Problem) Part1() (result int) {
	result = p.Ray(Step{Coord{0, 0}, Right})
	return
}

func (p *Problem) Part2() (result int) {
	for i := 0; i <= p.MaxX; i++ {
		p.Reset()
		tmp := p.Ray(Step{Coord{i, 0}, Down})
		if tmp > result {
			result = tmp
		}
		p.Reset()
		tmp = p.Ray(Step{Coord{i, p.MaxY}, Up})
		if tmp > result {
			result = tmp
		}
	}
	for j := 0; j <= p.MaxY; j++ {
		p.Reset()
		tmp := p.Ray(Step{Coord{0, j}, Right})
		if tmp > result {
			result = tmp
		}
		p.Reset()
		tmp = p.Ray(Step{Coord{p.MaxX, j}, Left})
		if tmp > result {
			result = tmp
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
	p.Cave = make(map[Coord]Tile)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			p.Cave[Coord{j, i}] = Tile{Type: c}
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
