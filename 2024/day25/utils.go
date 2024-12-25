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

type Key [5]int
type Lock [5]int

type Problem struct {
	Keys  []Key
	Locks []Lock
}

func (p *Problem) Fit(key Key, lock Lock) (result bool) {
	result = true
	for i := 0; i < 5; i++ {
		if key[i]+lock[i] > 5 {
			result = false
			break
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	for _, k := range p.Keys {
		for _, l := range p.Locks {
			if p.Fit(k, l) {
				result += 1
			}
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	fmt.Println("NOT TODAY")
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
	tmp := []string{}
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		if len(l) > 0 {
			tmp = append(tmp, l)
		} else {
			if tmp[0] == "....." {
				key := Key{5, 5, 5, 5, 5}
				for i := 1; i <= 5; i++ {
					for j := 0; j < 5; j++ {
						if tmp[i][j] == '.' {
							key[j] -= 1
						}
					}
				}
				p.Keys = append(p.Keys, key)
			} else {
				lock := Lock{0, 0, 0, 0, 0}
				for i := 1; i <= 5; i++ {
					for j := 0; j < 5; j++ {
						if tmp[i][j] == '#' {
							lock[j] += 1
						}
					}
				}
				p.Locks = append(p.Locks, lock)
			}
			tmp = []string{}
		}
	}
	if tmp[0] == "....." {
		key := Key{5, 5, 5, 5, 5}
		for i := 1; i <= 5; i++ {
			for j := 0; j < 5; j++ {
				if tmp[i][j] == '.' {
					key[j] -= 1
				}
			}
		}
		p.Keys = append(p.Keys, key)
	} else {
		lock := Lock{0, 0, 0, 0, 0}
		for i := 1; i <= 5; i++ {
			for j := 0; j < 5; j++ {
				if tmp[i][j] == '.' {
					lock[j] += 1
				}
			}
		}
		p.Locks = append(p.Locks, lock)
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
