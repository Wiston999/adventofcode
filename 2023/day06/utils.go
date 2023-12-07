package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/oleiade/lane/v2"

	"github.com/urfave/cli/v2"
)

type Problem struct {
	Races []Race
}

type Race struct {
	Time   int
	Record int
}

func (r *Race) GetDistance(j int) int {
	return (r.Time - j) * j
}

func (p *Problem) Part1() (result int) {
	result = 1
	for i, r := range p.Races {
		better := 0
		for j := 0; j <= r.Time; j++ {
			distance := r.GetDistance(j)
			log.Debug(fmt.Sprintf("Race %d holding for %d milliseconds reached %02d", i+1, j, distance))
			if distance > r.Record {
				better += 1
			}
		}
		result *= better
	}
	return
}

func (p *Problem) Part2() (result int) {
	timeStr := ""
	recordStr := ""
	for _, r := range p.Races {
		timeStr += fmt.Sprintf("%d", r.Time)
		recordStr += fmt.Sprintf("%d", r.Record)
	}
	race := Race{
		atoi(timeStr),
		atoi(recordStr),
	}
	// sort.Search implements binary search
	betterStart := sort.Search(race.Time, func(i int) bool { return race.GetDistance(i) > race.Record })
	log.Info(fmt.Sprintf("New race %#v has better results since millisecond %d", race, betterStart))
	result = (race.Time/2-betterStart)*2 + 1
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
	lines := strings.Split(strings.TrimSpace(strData), "\n")
	times := numbersSplitter.Split(strings.Split(lines[0], ":")[1], -1)
	records := numbersSplitter.Split(strings.Split(lines[1], ":")[1], -1)
	for i, t := range times {
		if i > 0 {
			r := Race{
				atoi(strings.TrimSpace(t)),
				atoi(strings.TrimSpace(records[i])),
			}
			p.Races = append(p.Races, r)
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
