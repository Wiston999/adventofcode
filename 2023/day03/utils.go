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
	Map     map[Coord]rune
	Numbers map[Coord]int
	MaxX    int
	MaxY    int
}

func (p *Problem) Part1() (result int) {
	for i := 0; i <= p.MaxX; i += 1 {
		for j := 0; j <= p.MaxY; j += 1 {
			buf := []rune{}
			t := j
			for ; p.Map[Coord{i, t}] >= '0' && p.Map[Coord{i, t}] <= '9'; t += 1 {
				buf = append(buf, p.Map[Coord{i, t}])
			}
			if len(buf) > 0 {
				t -= 1
				log.Debug(fmt.Sprintf("Found number %v at (%d, %d) - (%d, %d)", string(buf), i, j, i, t))
				isPart := false
				coords := []Coord{}

				for tj := j - 1; tj <= t+1; tj += 1 {
					coords = append(coords, Coord{i - 1, tj})
					coords = append(coords, Coord{i + 1, tj})
				}
				coords = append(coords, Coord{i, j - 1})
				coords = append(coords, Coord{i, t + 1})
				for _, c := range coords {
					log.Debug(fmt.Sprintf("Testing (%d, %d) %c", c.X, c.Y, p.Map[c]))
					if v, ok := p.Map[c]; ok {
						if v != '.' && (v > '9' || v < '0') {
							isPart = true
							break
						}
					}
				}
				if isPart {
					log.Info(fmt.Sprintf("Number %v at (%d, %d) is part", string(buf), i, j))
					result += atoi(string(buf))
				}
				j = t
			}
		}
	}
	return
}

func (p *Problem) Part2() (result int) {
	for i := 0; i <= p.MaxX; i += 1 {
		for j := 0; j <= p.MaxY; j += 1 {
			if v, ok := p.Map[Coord{i, j}]; ok && v == '*' {
				log.Debug(fmt.Sprintf("Found engine at (%d, %d)", i, j))
				numbers := []int{}
				visited := map[Coord]bool{}
				for ti := i - 1; ti <= i+1; ti += 1 {
					for tj := j - 1; tj <= j+1; tj += 1 {
						if _, visit := visited[Coord{ti, tj}]; !visit {
							log.Debug(fmt.Sprintf("Testing (%d, %d) - %v", ti, tj, p.Map[Coord{ti, tj}]))
							if v2, ok2 := p.Map[Coord{ti, tj}]; ok2 && v2 >= '0' && v2 <= '9' {
								k := tj - 1
								for ; p.Map[Coord{ti, k}] >= '0' && p.Map[Coord{ti, k}] <= '9'; k -= 1 {
								}
								l := tj + 1
								for ; p.Map[Coord{ti, l}] >= '0' && p.Map[Coord{ti, l}] <= '9'; l += 1 {
								}
								buf := []rune{}
								for z := k + 1; z < l; z++ {
									visited[Coord{ti, z}] = true
									buf = append(buf, p.Map[Coord{ti, z}])
								}
								log.Info(fmt.Sprintf("Found number %v near engine at (%d, %d)", string(buf), i, j))
								numbers = append(numbers, atoi(string(buf)))
							}
						}
					}
				}
				if len(numbers) >= 2 {
					tmp := 1
					for _, n := range numbers {
						tmp *= n
					}
					log.Info(fmt.Sprintf("Adding %d to result by engine at (%d, %d)", tmp, i, j))

					result += tmp
				}
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
	p = new(Problem)
	p.Map = make(map[Coord]rune)
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			p.Map[Coord{i, j}] = c
			if j > p.MaxY {
				p.MaxY = j
			}
		}
		if i > p.MaxX {
			p.MaxX = i
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
