package main

import (
	"container/list"
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
	Steps []string
	Boxes [256]list.List
}

type Lens struct {
	Name  string
	Power int
}

type Step string

type SequenceStep struct {
	Label string
	Op    Operation
	Power int
}

type Operation rune

const (
	Remove = '-'
	Set    = '='
)

func (s Step) Hash() (result int) {
	for _, c := range s {
		result += int(c)
		result *= 17
		result %= 256
	}
	return
}

func (s Step) ToSequenceStep() (result SequenceStep) {
	if strings.Contains(string(s), "-") {
		parts := strings.Split(string(s), "-")
		result.Label = parts[0]
		result.Op = Remove
	} else {
		parts := strings.Split(string(s), "=")
		result.Label = parts[0]
		result.Op = Set
		result.Power = atoi(parts[1])
	}
	return
}

func (p *Problem) PrintBoxes() (result string) {
	for i, b := range p.Boxes {
		result += fmt.Sprintf("[%02d] ", i)
		for e := b.Front(); e != nil; e = e.Next() {
			result += fmt.Sprintf("[%v] ", e.Value)
		}
		result += "\n"
	}
	return
}

func (p *Problem) Part1() (result int) {
	for _, s := range p.Steps {
		h := Step(s).Hash()
		result += h
		log.Debug(fmt.Sprintf("Hash of %s : %d", s, h))
	}
	return
}

func (p *Problem) Part2() (result int) {
	for _, s := range p.Steps {
		seq := Step(s).ToSequenceStep()
		dest := Step(seq.Label).Hash()
		log.Debug(fmt.Sprintf("Operating (%c) on label %s on box %d", seq.Op, seq.Label, dest))
		var e *list.Element
		var eLens *Lens
		for e = p.Boxes[dest].Front(); e != nil; e = e.Next() {
			eLens = e.Value.(*Lens)
			if eLens.Name == seq.Label {
				break
			}
		}
		if seq.Op == Remove {
			if eLens != nil && eLens.Name == seq.Label {
				p.Boxes[dest].Remove(e)
			}
		} else {
			if eLens != nil && eLens.Name == seq.Label {
				eLens.Power = seq.Power
			} else {
				p.Boxes[dest].PushBack(&Lens{seq.Label, seq.Power})
			}
		}
	}
	for i, b := range p.Boxes {
		for j, e := 1, b.Front(); e != nil; e, j = e.Next(), j+1 {
			result += (i + 1) * j * e.Value.(*Lens).Power
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
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		p.Steps = strings.Split(l, ",")
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
