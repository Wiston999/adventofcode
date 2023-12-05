package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/oleiade/lane/v2"

	"github.com/urfave/cli/v2"
)

type Problem struct {
	Cards []Card
}

type Card struct {
	Winning map[int]void
	Owned   map[int]void
}

func (c *Card) CountWin() (result int) {
	for k := range c.Winning {
		if _, ok := c.Owned[k]; ok {
			result += 1
		}
	}
	return
}

type void struct{}

func (p *Problem) Part1() (result int) {
	for i, c := range p.Cards {
		count := c.CountWin()
		if count > 0 {
			log.Info(fmt.Sprintf("Found %d winning numbers in card %d", count, i+1))
			result += int(math.Pow(2, float64(count-1)))
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	winCount := make([]int, len(p.Cards))
	accCount := make([]int, len(p.Cards))
	for i, c := range p.Cards {
		winCount[i] = c.CountWin()
	}
	// Reverse
	for i, j := 0, len(winCount)-1; i < j; i, j = i+1, j-1 {
		winCount[i], winCount[j] = winCount[j], winCount[i]
	}

	for i, c := range winCount {
		accCount[i] = 1
		for j := 1; j <= c; j += 1 {
			accCount[i] += accCount[i-j]
		}
	}

	log.Debug(winCount)
	log.Debug(accCount)
	for _, c := range accCount {
		result += c
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
	numbersSplitter := regexp.MustCompile(" +")
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		numbers := strings.Split(strings.Split(l, ": ")[1], " | ")
		winning, owned := numbers[0], numbers[1]
		card := Card{make(map[int]void), make(map[int]void)}
		for _, w := range numbersSplitter.Split(winning, -1) {
			log.Debug(fmt.Sprintf("Adding %d (%s) to winning set", atoi(w), w))
			card.Winning[atoi(w)] = void{}
		}
		for _, w := range numbersSplitter.Split(owned, -1) {
			log.Debug(fmt.Sprintf("Adding %d (%s) to owned set", atoi(w), w))
			card.Owned[atoi(w)] = void{}
		}
		p.Cards = append(p.Cards, card)
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
