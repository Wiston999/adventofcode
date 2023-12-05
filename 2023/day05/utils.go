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
	Seeds               []int
	SeedSoil            Layer
	SoilFertilizer      Layer
	FertilizerWater     Layer
	WaterLight          Layer
	LightTemperature    Layer
	TemperatureHumidity Layer
	HumidityLocation    Layer
}

type Layer struct {
	Transformations []Transformation
}

type Transformation struct {
	SourceStart int
	DestStart   int
	Range       int
}

func (l *Layer) GetDestination(src int) (destination int) {
	destination = src
	for _, t := range l.Transformations {
		if src >= t.DestStart && src <= t.DestStart+t.Range {
			destination = t.SourceStart + (destination - t.DestStart)
			return
		}
	}
	return
}

func (l *Layer) GetSource(dst int) (source int) {
	source = dst
	for _, t := range l.Transformations {
		if dst >= t.SourceStart && dst <= t.SourceStart+t.Range {
			source = t.DestStart + (source - t.SourceStart)
			return
		}
	}
	return
}

func (p *Problem) GetSeedLocation(seed int) (location int) {
	location = p.SeedSoil.GetDestination(seed)
	location = p.SoilFertilizer.GetDestination(location)
	location = p.FertilizerWater.GetDestination(location)
	location = p.WaterLight.GetDestination(location)
	location = p.LightTemperature.GetDestination(location)
	location = p.TemperatureHumidity.GetDestination(location)
	location = p.HumidityLocation.GetDestination(location)
	return
}

func (p *Problem) GetLocationSeed(location int) (seed int) {
	seed = p.HumidityLocation.GetSource(location)
	seed = p.TemperatureHumidity.GetSource(seed)
	seed = p.LightTemperature.GetSource(seed)
	seed = p.WaterLight.GetSource(seed)
	seed = p.FertilizerWater.GetSource(seed)
	seed = p.SoilFertilizer.GetSource(seed)
	seed = p.SeedSoil.GetSource(seed)
	return
}

func (p *Problem) Part1() (result int) {
	result = 9999999999999
	for _, s := range p.Seeds {
		sl := p.GetSeedLocation(s)
		log.Info(fmt.Sprintf("Seed %d is located at %d", s, sl))
		if sl < result {
			result = sl
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	for i := 0; i < 99999999999999; i++ {
		seed := p.GetLocationSeed(i)
		log.Debug(fmt.Sprintf("Location %d from seed %d", i, seed))
		for j := 0; j < len(p.Seeds); j += 2 {
			if seed >= p.Seeds[j] && seed <= p.Seeds[j]+p.Seeds[j+1] {
				return i
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
	strData := string(byteData)
	p = new(Problem)
	var context *Layer
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		if len(l) > 0 {
			log.Debug(fmt.Sprintf("Parsing line %03d: %s (context: %v)", i, l, context))
			if strings.HasPrefix(l, "seeds:") {
				for _, s := range strings.Split(strings.Split(l, ": ")[1], " ") {
					p.Seeds = append(p.Seeds, atoi(s))
				}
				continue
			}
			if strings.HasPrefix(l, "seed-to-soil map") {
				context = &p.SeedSoil
				continue
			}
			if strings.HasPrefix(l, "soil-to-fertilizer map") {
				context = &p.SoilFertilizer
				continue
			}
			if strings.HasPrefix(l, "fertilizer-to-water map") {
				context = &p.FertilizerWater
				continue
			}
			if strings.HasPrefix(l, "water-to-light map") {
				context = &p.WaterLight
				continue
			}
			if strings.HasPrefix(l, "light-to-temperature map") {
				context = &p.LightTemperature
				continue
			}
			if strings.HasPrefix(l, "temperature-to-humidity map") {
				context = &p.TemperatureHumidity
				continue
			}
			if strings.HasPrefix(l, "humidity-to-location map") {
				context = &p.HumidityLocation
				continue
			}
			values := strings.Split(l, " ")
			context.Transformations = append(context.Transformations, Transformation{atoi(values[0]), atoi(values[1]), atoi(values[2])})
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
