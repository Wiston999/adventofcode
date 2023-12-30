package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/oleiade/lane/v2"

	"github.com/urfave/cli/v2"
)

type Problem struct {
	Hails []Hail
}

type Hail struct {
	P Coord
	V Coord
}

func (c *Coord) TestArea(min, max float64) (result bool) {
	if c.X < min || c.X > max {
		return
	}
	if c.Y < min || c.Y > max {
		return
	}
	result = true
	return
}

func (h *Hail) Collide(other Hail) (point Coord, result bool) {
	m1 := -(h.V.Y / -h.V.X)
	m2 := -(other.V.Y / -other.V.X)
	b1 := -((h.V.X*h.P.Y - h.V.Y*h.P.X) / -h.V.X)
	b2 := -((other.V.X*other.P.Y - other.V.Y*other.P.X) / -other.V.X)
	if (m2 - m1) == 0 {
		// Parallel
		return
	}
	k := (b1 - b2) / (m2 - m1)
	point.X = k
	point.Y = m1*k + b1
	if ((-h.P.X+point.X)/h.V.X) < 0 || ((-h.P.Y+point.Y)/h.V.Y) < 0 {
		// Intersection in the past for h
		return
	}
	if ((-other.P.X+point.X)/other.V.X) < 0 || ((-other.P.Y+point.Y)/other.V.Y) < 0 {
		// Intersection in the past for other
		return
	}
	result = true
	return
}

func GetSpeeds(distDiff, speedDiff int) (result []int) {
	for i := -1000; i < 1000; i++ {
		if i != speedDiff && distDiff%(i-speedDiff) == 0 {
			result = append(result, i)
		}
	}
	return
}

func SetIntersect(a, b []int) (result []int) {
	for _, vA := range a {
		for _, vB := range b {
			if vA == vB {
				result = append(result, vA)
			}
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	for i, h1 := range p.Hails {
		for j := i + 1; j < len(p.Hails); j++ {
			h2 := p.Hails[j]
			p, r := h1.Collide(h2)
			if r && p.TestArea(200000000000000, 400000000000000) {
				log.Debug(fmt.Sprintf("Points %v %v collide at %v", h1, h2, p))
				result += 1
			}
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	possibleX, possibleY, possibleZ := []int{}, []int{}, []int{}
	for i, h1 := range p.Hails {
		for j := i + 1; j < len(p.Hails); j++ {
			h2 := p.Hails[j]
			if h1.V.X == h2.V.X {
				possible := GetSpeeds(int(h2.P.X-h1.P.X), int(h1.V.X))
				if len(possibleX) == 0 {
					possibleX = possible
				} else {
					possibleX = SetIntersect(possibleX, possible)
				}
			}
			if h1.V.Y == h2.V.Y {
				possible := GetSpeeds(int(h2.P.Y-h1.P.Y), int(h1.V.Y))
				if len(possibleY) == 0 {
					possibleY = possible
				} else {
					possibleY = SetIntersect(possibleY, possible)
				}
			}
			if h1.V.Z == h2.V.Z {
				possible := GetSpeeds(int(h2.P.Z-h1.P.Z), int(h1.V.Z))
				if len(possibleZ) == 0 {
					possibleZ = possible
				} else {
					possibleZ = SetIntersect(possibleZ, possible)
				}
			}
		}
	}
	sort.Slice(p.Hails, func(i, j int) bool {
		h1, h2 := p.Hails[i], p.Hails[j]
		return math.Abs(h1.P.X+h1.P.Y+h1.P.Z) < math.Abs(h2.P.X+h2.P.Y+h2.P.Z)
	})
	rock := Hail{
		V: Coord{float64(possibleX[0]), float64(possibleY[0]), float64(possibleZ[0])},
	}
	h1, h2 := p.Hails[0], p.Hails[1]
	mA := (h1.V.Y - rock.V.Y) / (h1.V.X - rock.V.X)
	mB := (h2.V.Y - rock.V.Y) / (h2.V.X - rock.V.X)
	cA := h1.P.Y - (mA * h1.P.X)
	cB := h2.P.Y - (mB * h2.P.X)
	rock.P.X = (cB - cA) / (mA - mB)
	rock.P.Y = mA*rock.P.X + cA
	time := (rock.P.X - h1.P.X) / (h1.V.X - rock.V.X)
	rock.P.Z += h1.P.Z + (h1.V.Z-rock.V.Z)*time
	result = int(rock.P.X + rock.P.Y + rock.P.Z)

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
		h := Hail{}
		fmt.Sscanf(l, "%f, %f, %f @ %f, %f, %f", &h.P.X, &h.P.Y, &h.P.Z, &h.V.X, &h.V.Y, &h.V.Z)
		p.Hails = append(p.Hails, h)
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
	X, Y, Z float64
}

func (c *Coord) Manhattan(oc Coord) float64 {
	return math.Abs(float64(oc.X-c.X)) + math.Abs(float64(oc.Y-c.Y))
}
