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
	StartRobot, Robot         Coord
	StartWarehouse, Warehouse map[Coord]string
	MaxX, MaxY                int
	Movements                 []string
}

func (p *Problem) Print() (result string) {
	for i := 0; i <= p.MaxX; i += 1 {
		for j := 0; j <= p.MaxY; j += 1 {
			if p.Robot.X == i && p.Robot.Y == j {
				result += "@"
			} else {
				result += p.Warehouse[Coord{i, j}]
			}
		}
		result += "\n"
	}
	return
}

func (p *Problem) CanPush(left, right Coord, m string) (result bool) {
	nX := 1
	if m == "^" {
		nX = -1
	}
	vL := p.Warehouse[Coord{left.X + nX, left.Y}]
	vR := p.Warehouse[Coord{right.X + nX, right.Y}]
	result = vL == "." && vR == "."
	return
}

func (p *Problem) Push(start Coord, m string, doubled bool) Coord {
	next := start
	switch m {
	case "<":
		next.Y -= 1
	case ">":
		next.Y += 1
	case "v":
		next.X += 1
	case "^":
		next.X -= 1
	}
	if m == "<" || m == ">" || !doubled {
		if p.Warehouse[next] == "." {
			p.Warehouse[next] = p.Warehouse[start]
			p.Warehouse[start] = "."
			return next
		} else if p.Warehouse[next] == "#" {
			return start
		} else {
			p.Push(next, m, doubled)
			if p.Warehouse[next] == "." {
				p.Warehouse[next] = p.Warehouse[start]
				p.Warehouse[start] = "."
				return next
			}
		}
	} else {
		switch p.Warehouse[next] {
		case "#":
			return start
		case ".":
			p.Warehouse[next] = p.Warehouse[start]
			p.Warehouse[start] = "."
			return next
		case "[":
			right := p.Push(next, m, doubled)
			left := p.Push(Coord{next.X, next.Y + 1}, m, doubled)
			log.Debug(fmt.Sprintf("Pushing right edge of box at %v, %v %v S:%v\n%s", next, right, left, start, p.Print()))
			if p.Warehouse[next] == "." && p.Warehouse[Coord{next.X, next.Y + 1}] == "." {
				p.Warehouse[next] = p.Warehouse[start]
				p.Warehouse[start] = "."
				// p.Warehouse[Coord{start.X, start.Y + 1}] = "."
				log.Debug(fmt.Sprintf("Pushed right edge of box at %v, %v %v S:%v\n%s", next, right, left, start, p.Print()))
				return next
			}
		case "]":
			right := p.Push(Coord{next.X, next.Y - 1}, m, doubled)
			left := p.Push(next, m, doubled)
			log.Debug(fmt.Sprintf("Pushing left edge of box at %v, %v %v S:%v\n%s", next, right, left, start, p.Print()))
			if p.Warehouse[next] == "." && p.Warehouse[Coord{next.X, next.Y - 1}] == "." {
				p.Warehouse[next] = p.Warehouse[start]
				p.Warehouse[start] = "."
				// p.Warehouse[Coord{start.X, start.Y - 1}] = "."
				log.Debug(fmt.Sprintf("Pushed left edge of box at %v, %v %v S:%v\n%s", next, right, left, start, p.Print()))
				return next
			}
		}
	}
	return start
}

func (p *Problem) CopyWarehouse() (result map[Coord]string) {
	result = make(map[Coord]string)
	for k, v := range p.StartWarehouse {
		result[k] = v
	}
	return
}

func (p *Problem) Part1() (result int) {
	p.Robot = p.StartRobot
	p.Warehouse = p.CopyWarehouse()
	for _, m := range p.Movements {
		p.Robot = p.Push(p.Robot, m, false)
	}
	for i := 0; i <= p.MaxX; i += 1 {
		for j := 0; j <= p.MaxY; j += 1 {
			if p.Warehouse[Coord{i, j}] == "O" {
				result += i*100 + j
			}
		}
	}
	fmt.Println(p.Print())
	return
}

func (p *Problem) Part2() (result int) {
	p.Robot = p.StartRobot
	p.Warehouse = make(map[Coord]string)
	for i := 0; i <= p.MaxX; i += 1 {
		for j := 0; j <= p.MaxY; j += 1 {
			switch p.StartWarehouse[Coord{i, j}] {
			case "#":
				p.Warehouse[Coord{i, 2 * j}] = "#"
				p.Warehouse[Coord{i, 2*j + 1}] = "#"
			case ".":
				p.Warehouse[Coord{i, 2 * j}] = "."
				p.Warehouse[Coord{i, 2*j + 1}] = "."
			case "O":
				p.Warehouse[Coord{i, 2 * j}] = "["
				p.Warehouse[Coord{i, 2*j + 1}] = "]"
			}
		}
	}
	p.MaxY *= 2 + 1
	p.Robot.Y *= 2
	for _, m := range p.Movements {
		p.Robot = p.Push(p.Robot, m, true)
		log.Debug(fmt.Sprintf("Movement %s %v:\n%s", m, p.Robot, p.Print()))
	}
	fmt.Println(p.Print())
	for i := 0; i <= p.MaxX; i += 1 {
		for j := 0; j <= p.MaxY; j += 1 {
			if p.Warehouse[Coord{i, j}] == "[" {
				result += i*100 + j
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
	p.StartWarehouse = make(map[Coord]string)
	movements := false
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		if len(l) == 0 {
			movements = true
			continue
		}
		for j, c := range strings.Split(strings.TrimSpace(l), "") {
			if !movements {
				if c == "@" {
					p.StartRobot = Coord{i, j}
					p.StartWarehouse[Coord{i, j}] = "."
				} else {
					p.StartWarehouse[Coord{i, j}] = c
				}
				if j > p.MaxY {
					p.MaxY = j
				}
				if i > p.MaxX {
					p.MaxX = i
				}
			} else {
				p.Movements = append(p.Movements, c)
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
