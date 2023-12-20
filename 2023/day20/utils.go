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
	Devices             map[string]Device
	CountLow, CountHigh int
}

type Device struct {
	Current    bool
	CurrentMap map[string]bool
	Type       DevType
	Name       string
	Inputs     []string
	Outputs    []string
}

type State struct {
	Src     string
	Current string
	Value   bool
}

func (d *Device) Process() (low, high int) {
	switch d.Type {
	case FlipFlop:

	}
	return
}

type DevType string

const (
	FlipFlop    = "flip-flop"
	Conjunction = "conjunction"
	Broadcaster = "broadcaster"
	Other       = "other"
)

func (p *Problem) Push(inspect map[string]bool) (result map[string]bool) {
	result = make(map[string]bool)
	pending := list.New()
	pending.PushBack(State{
		"",
		"broadcaster",
		false,
	})

	for pending.Len() > 0 {
		newPending := list.New()
		for pending.Len() > 0 {
			current := pending.Remove(pending.Front()).(State)
			dev, found := p.Devices[current.Current]
			if _, ins := inspect[current.Current]; ins {
				if !current.Value {
					result[current.Current] = true
				}
			}
			if current.Value {
				p.CountHigh++
			} else {
				p.CountLow++
			}
			if !found {
				continue
			}
			var newValue bool
			switch dev.Type {
			case Broadcaster:
				newValue = false
			case FlipFlop:
				if current.Value {
					continue
				}
				newValue = !dev.Current
				dev.Current = !dev.Current
			case Conjunction:
				dev.CurrentMap[current.Src] = current.Value
				all := true
				for _, k := range dev.Inputs {
					if v := dev.CurrentMap[k]; !v {
						all = false
						break
					}
				}
				newValue = !all
			default:
				newValue = current.Value
			}
			p.Devices[current.Current] = dev
			for _, o := range dev.Outputs {
				newPending.PushBack(State{dev.Name, o, newValue})
			}
		}
		pending = newPending
	}
	return
}

func (p *Problem) Part1() (result int) {
	p.Reset()
	for i := 0; i < 1000; i++ {
		p.Push(make(map[string]bool))
	}
	log.Debug(fmt.Sprintf("Count High: %d, Count Low: %d", p.CountHigh, p.CountLow))
	result = p.CountLow * p.CountHigh
	return
}

func (p *Problem) Part2() (result int) {
	p.Reset()
	pending := make(map[string]bool)
	found := make(map[string]int)
	src := ""
	for k, dev := range p.Devices {
		for _, o := range dev.Outputs {
			if o == "rx" {
				src = k
			}
		}
	}
	for _, k := range p.Devices[src].Inputs {
		pending[k] = true
	}
	log.Debug(fmt.Sprintf("Looking for cycles in %v", pending))
	for i := 1; len(found) < len(pending); i++ {
		result := p.Push(pending)
		for k, v := range result {
			if _, ok := found[k]; !ok && v {
				found[k] = i
			}
		}
	}
	log.Debug(fmt.Sprintf("Cycles found %v", found))
	var values []int
	for _, v := range found {
		values = append(values, v)
	}
	result = LCM(values[0], values[1], values[2:]...)
	return
}

func (p *Problem) Reset() {
	for k, dev := range p.Devices {
		dev.Current = false
		for k := range dev.CurrentMap {
			dev.CurrentMap[k] = false
		}
		p.Devices[k] = dev
	}
}

func NewProblem(ctx *cli.Context) (p *Problem, err error) {
	input := ctx.String("input")
	byteData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	p = new(Problem)
	p.Devices = make(map[string]Device)
	p.Devices["broadcaster"] = Device{
		Type: Broadcaster,
		Name: Broadcaster,
	}
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		order := strings.Split(l, " -> ")
		dev := Device{}
		dev.CurrentMap = make(map[string]bool)
		dev.Name = string(order[0][1:])
		if order[0][0] == '%' {
			dev.Type = FlipFlop
		} else if order[0][0] == '&' {
			dev.Type = Conjunction
		} else if order[0] != Broadcaster {
			dev.Type = Other
			dev.Name = order[0]
		} else if order[0] == Broadcaster {
			dev = p.Devices["broadcaster"]
		}
		for _, o := range strings.Split(order[1], ", ") {
			dev.Outputs = append(dev.Outputs, o)
		}
		p.Devices[dev.Name] = dev
	}
	for k, dev := range p.Devices {
		for _, o := range dev.Outputs {
			if output, ok := p.Devices[o]; ok {
				output.Inputs = append(output.Inputs, k)
				p.Devices[o] = output
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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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
