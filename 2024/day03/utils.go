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

type void struct{}

var null void

type Problem struct {
	Input string
}

func (p *Problem) Part1() (result int) {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	operations := re.FindAllString(p.Input, -1)
	log.Debug(operations)
	for _, op := range operations {
		a, b := 0, 0
		fmt.Sscanf(op, "mul(%d,%d)", &a, &b)
		result += a * b
	}
	return
}

func (p *Problem) Part2() (result int) {
	opRE := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	doRE := regexp.MustCompile(`do\(\)`)
	dnRE := regexp.MustCompile(`don't\(\)`)
	operations := opRE.FindAllStringIndex(p.Input, -1)
	dos := doRE.FindAllStringIndex(p.Input, -1)
	donts := dnRE.FindAllStringIndex(p.Input, -1)
	log.Debug(operations)
	log.Debug(dos)
	log.Debug(donts)
	compute := true
	for op, do, dn := 0, 0, 0; op < len(operations) || do < len(dos) || dn < len(donts); {
		cOperation, cDo, cDont := len(p.Input)+1, len(p.Input)+1, len(p.Input)+1
		if op < len(operations) {
			cOperation = operations[op][0]
		}
		if do < len(dos) {
			cDo = dos[do][0]
		}
		if dn < len(donts) {
			cDont = donts[dn][0]
		}
		if cOperation < cDo && cOperation < cDont {
			if compute {
				a, b := 0, 0
				fmt.Sscanf(p.Input[operations[op][0]:operations[op][1]], "mul(%d,%d)", &a, &b)
				result += a * b
			}
			op += 1
		}
		if cDo < cOperation && cDo < cDont {
			compute = true
			do += 1
		}
		if cDont < cOperation && cDont < cDo {
			compute = false
			dn += 1
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
	p.Input = strData
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
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
