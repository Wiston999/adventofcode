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
	Slabs []Slab
	ByZ   map[int]map[int]bool
}

type Space struct {
	Map              map[Coord]int
	MaxX, MaxY, MaxZ int
}

type Slab struct {
	Start, End Coord
}

func (p *Problem) Copy() (result Problem) {
	result.Slabs = make([]Slab, len(p.Slabs))
	for i, s := range p.Slabs {
		result.Slabs[i] = s
	}
	result.ByZ = make(map[int]map[int]bool)
	for z, v := range p.ByZ {
		result.ByZ[z] = make(map[int]bool)
		for k, l := range v {
			result.ByZ[z][k] = l
		}
	}
	return
}

func (p *Problem) ToSpace() (result Space) {
	result.Map = make(map[Coord]int)
	for n, s := range p.Slabs {
		for i := s.Start.X; i <= s.End.X; i++ {
			for j := s.Start.Y; j <= s.End.Y; j++ {
				for k := s.Start.Z; k <= s.End.Z; k++ {
					c := Coord{i, j, k}
					if v, ok := result.Map[c]; ok {
						log.Warn(fmt.Sprintf("Position %v already used by %v", c, v))
					}
					result.Map[c] = n
					if k > result.MaxZ {
						result.MaxZ = k
					}
				}
				if j > result.MaxY {
					result.MaxY = j
				}
			}
			if i > result.MaxX {
				result.MaxX = i
			}
		}
	}
	return
}

func (p *Problem) CanDescend(s Slab) (result bool) {
	result = true
	// log.Debug(fmt.Sprintf("Checking if %v can descend %v", s, p.ByZ))
	for c := range p.ByZ[s.Start.Z-1] {
		other := p.Slabs[c]
		for i := s.Start.X; i <= s.End.X; i++ {
			for j := s.Start.Y; j <= s.End.Y; j++ {
				// log.Debug(fmt.Sprintf("%d <= %d <= %d, %d <= %d <= %d", other.Start.X, i, other.End.X, other.Start.Y, j, other.End.Y))
				if i >= other.Start.X && i <= other.End.X && j >= other.Start.Y && j <= other.End.Y {
					log.Debug(fmt.Sprintf("Slab [%v] can't descend due to %d [%v] (%d, %d)", s, c, other, i, j))
					return false
				}
			}
		}
	}
	return
}

func (p *Problem) UpdateZ(z, slab int) {
	if _, ok := p.ByZ[z]; !ok {
		p.ByZ[z] = make(map[int]bool)
	}
	p.ByZ[z][slab] = true
}

func (p *Problem) Compact() (result map[int]bool) {
	result = make(map[int]bool)
	for n, s := range p.Slabs {
		if s.Start.Z <= 1 {
			continue
		}

		if p.CanDescend(s) {
			log.Debug(fmt.Sprintf("Slab %d [%v] can descend", n, s))
			result[n] = true
			delete(p.ByZ[s.End.Z], n)
			s.Start.Z -= 1
			s.End.Z -= 1
			p.UpdateZ(s.Start.Z, n)
			p.Slabs[n] = s
		}
	}
	return
}

func (p *Problem) Part1() (result int) {
	for len(p.Compact()) > 0 {
	}
	for n, s := range p.Slabs {
		for i := s.Start.Z; i <= s.End.Z; i++ {
			delete(p.ByZ[i], n)
		}
		anyDescend := false
		for other, _ := range p.ByZ[s.End.Z+1] {
			if s.End.Z+1 == p.Slabs[other].Start.Z && p.CanDescend(p.Slabs[other]) {
				log.Debug(fmt.Sprintf("%d [%v] would descend if %d [%v] is disintegrated", other, p.Slabs[other], n, s))
				anyDescend = true
				break
			}
		}
		if !anyDescend {
			log.Info(fmt.Sprintf("Slab %d at %v can be disintegrated", n, s))
			result += 1
		}
		for i := s.Start.Z; i <= s.End.Z; i++ {
			p.ByZ[i][n] = true
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	for n, s := range p.Slabs {
		tmp := p.Copy()
		for i := s.Start.Z; i <= s.End.Z; i++ {
			delete(tmp.ByZ[i], n)
		}
		descended := make(map[int]bool)
		for {
			chain := tmp.Compact()
			if len(chain) == 0 {
				break
			}
			for k, v := range chain {
				descended[k] = v
			}
		}
		if len(descended) > 0 {
			log.Info(fmt.Sprintf("Disintegrating %d moved %d slabs", n, len(descended)))
		}
		result += len(descended)
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
	p.ByZ = make(map[int]map[int]bool)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		s := Slab{}
		fmt.Sscanf(
			l, "%d,%d,%d~%d,%d,%d",
			&s.Start.X, &s.Start.Y, &s.Start.Z,
			&s.End.X, &s.End.Y, &s.End.Z,
		)
		if s.Start.X > s.End.X || s.Start.Y > s.End.Y || s.Start.Z > s.End.Z {
			log.Warn(fmt.Sprintf("Slab %d has flipped start-end %v", i, s))
		}
		for j := s.Start.Z; j <= s.End.Z; j++ {
			if _, ok := p.ByZ[j]; !ok {
				p.ByZ[j] = make(map[int]bool)
			}
			p.ByZ[j][i] = true
		}
		p.Slabs = append(p.Slabs, s)
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
	X, Y, Z int
}

func (c *Coord) Manhattan(oc Coord) float64 {
	return math.Abs(float64(oc.X-c.X)) + math.Abs(float64(oc.Y-c.Y))
}
