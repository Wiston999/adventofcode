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

type Position string

func JoinPositions(p []Position) (result string) {
	for _, e := range p {
		result += string(e)
	}
	return
}

const (
	UP       = "^"
	DOWN     = "v"
	RIGHT    = ">"
	LEFT     = "<"
	Activate = "A"
	Gap      = "X"
)

var NumPad map[Position]Coord = map[Position]Coord{
	"A": {3, 2},
	"0": {3, 1},
	"1": {2, 0},
	"2": {2, 1},
	"3": {2, 2},
	"4": {1, 0},
	"5": {1, 1},
	"6": {1, 2},
	"7": {0, 0},
	"8": {0, 1},
	"9": {0, 2},
	"X": {3, 0},
}

var DirPad map[Position]Coord = map[Position]Coord{
	"A": {0, 2},
	"<": {1, 0},
	"v": {1, 1},
	">": {1, 2},
	"^": {0, 1},
	"X": {0, 0},
}

type DirRobot struct {
	Current Position
}

type NumRobot struct {
	Current Position
}

func GetMovements(relative int, positive, negative Position) (result []Position) {
	if relative > 0 {
		for i := 0; i < relative; i++ {
			result = append(result, positive)
		}
	} else if relative < 0 {
		relative = -relative
		for i := 0; i < relative; i++ {
			result = append(result, negative)
		}
	}
	return
}

func MoveRobot(current, next Position, pad map[Position]Coord) (result []Position) {
	rLocation := pad[current]
	nextLocation := pad[next]
	gap := pad[Gap]
	vertical := GetMovements(nextLocation.X-rLocation.X, DOWN, UP)
	horizontal := GetMovements(nextLocation.Y-rLocation.Y, RIGHT, LEFT)
	if nextLocation.Y > rLocation.Y && gap != (Coord{nextLocation.X, rLocation.Y}) {
		result = append(result, vertical...)
		result = append(result, horizontal...)
		return
	}
	if gap != (Coord{rLocation.X, nextLocation.Y}) {
		result = append(result, horizontal...)
		result = append(result, vertical...)
		return
	}
	result = append(result, vertical...)
	result = append(result, horizontal...)
	return
}

func (r *DirRobot) Move(next Position) (result []Position) {
	current := r.Current
	r.Current = next
	return MoveRobot(current, next, DirPad)
}

func (r *NumRobot) Move(next Position) (result []Position) {
	current := r.Current
	r.Current = next
	return MoveRobot(current, next, NumPad)
}

type Problem struct {
	Sequences []string
}

func (p *Problem) Numeric(sequence string) (result int) {
	return atoi(strings.ReplaceAll(sequence, "A", ""))
}

func (p *Problem) Solve(numRobots int) (result int) {
	for i, s := range p.Sequences {
		cRobot := NumRobot{"A"}
		RC := make([]Position, 0)
		freqTable := make(map[string]int)
		for _, n := range s {
			RC = append(RC, cRobot.Move(Position(n))...)
			RC = append(RC, "A")
		}
		steps := JoinPositions(RC)
		freqTable[steps] = 1
		log.Debug(fmt.Sprintf("[%d] Initial frequency table: %v", i, freqTable))
		for j := 0; j <= numRobots; j++ {
			tmpTable := make(map[string]int)
			for sequence, count := range freqTable {
				tmpRobot := DirRobot{"A"}
				for _, n := range sequence {
					steps := make([]Position, 0)
					steps = append(steps, tmpRobot.Move(Position(n))...)
					steps = append(steps, "A")
					log.Debug(fmt.Sprintf("Steps for %v [%d]: %v", sequence, count, steps))
					tmpTable[JoinPositions(steps)] += count
				}
			}
			log.Debug(fmt.Sprintf("[%d] Temp table %d: %v", i, j, tmpTable))
			freqTable = tmpTable
		}
		sum := 0
		for _, v := range freqTable {
			sum += v
		}
		log.Info(fmt.Sprintf("[%d] %s %d*%d = %d", i, s, sum, p.Numeric(s), sum*p.Numeric(s)))
		result += sum * p.Numeric(s)
	}
	return
}

func (p *Problem) Part1() (result int) {
	return p.Solve(2)
}

func (p *Problem) Part2() (result int) {
	return p.Solve(25)
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
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		p.Sequences = append(p.Sequences, l)
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
