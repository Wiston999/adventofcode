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
	Hands []Hand
}

type Hand struct {
	Cards []int
	Bid   int
	Kind  HandKind
}

type HandKind int

const (
	High = iota
	Pair
	TwoPair
	Three
	Full
	Four
	Five
)

func (h *Hand) Parse(s string) {
	for _, c := range s {
		if v, err := strconv.Atoi(string(c)); err == nil {
			if v == 1 {
				v = 15
			}
			h.Cards = append(h.Cards, v)
		} else {
			switch c {
			case 'T':
				h.Cards = append(h.Cards, 10)
			case 'J':
				h.Cards = append(h.Cards, 11)
			case 'Q':
				h.Cards = append(h.Cards, 12)
			case 'K':
				h.Cards = append(h.Cards, 13)
			case 'A':
				h.Cards = append(h.Cards, 14)
			}
		}
	}
}

func (h *Hand) ComputeKind(joker bool) HandKind {
	h.Kind = High
	cards := make([]int, 5)
	jokers := 0
	for i, c := range h.Cards {
		cards[i] = c
		if c == 11 {
			jokers += 1
		}
	}
	sort.Ints(cards)
	for i, j := 0, len(cards)-1; i < j; i, j = i+1, j-1 {
		cards[i], cards[j] = cards[j], cards[i]
	}
	i := 0
	for i < len(cards) {
		j := 1
		for ; (i+j) < len(cards) && cards[i] == cards[i+j]; j += 1 {
			if joker && cards[i+j] == 11 {
				if cards[i+j] == 11 || cards[i] != cards[i+j] {
					break
				}
			} else if !joker && cards[i] != cards[i+j] {
				break
			}
		}
		i += j
		switch j {
		case 2:
			switch h.Kind {
			case High:
				h.Kind = Pair
			case Pair:
				h.Kind = TwoPair
			case Three:
				h.Kind = Full
			}
		case 3:
			switch h.Kind {
			case High:
				h.Kind = Three
			case Pair:
				h.Kind = Full
			}
		case 4:
			h.Kind = Four
		case 5:
			h.Kind = Five
		}
	}
	if joker && jokers > 0 {
		log.Debug(fmt.Sprintf("Upgrading hand %#v using %d jokers", h, jokers))
		switch h.Kind {
		case High:
			switch jokers {
			case 1:
				h.Kind = Pair
			case 2:
				h.Kind = Three
			case 3:
				h.Kind = Four
			case 4:
				h.Kind = Five
			case 5: // JJJJJ
				h.Kind = Five
			}
		case Pair:
			switch jokers {
			case 2:
				h.Kind = Four
			case 3:
				h.Kind = Five
			default:
				h.Kind += 1 + HandKind(jokers) // Skip TwoPairs step
			}
		case TwoPair:
			switch jokers {
			case 1:
				h.Kind = Full
			default:
				h.Kind += 1 + HandKind(jokers) // Skip Three step
			}
		case Three:
			h.Kind += 1 + HandKind(jokers) // Skip Full step
		default:
			h.Kind += HandKind(jokers)
		}
		log.Debug(fmt.Sprintf("Upgraded hand %#v using %d jokers", h, jokers))
	}
	return h.Kind
}

func (h *Hand) Compare(other Hand, joker bool) bool {
	if h.Kind == other.Kind {
		for i, v := range h.Cards {
			if v != other.Cards[i] {
				if joker {
					if v == 11 {
						return true
					} else if other.Cards[i] == 11 {
						return false
					}
				}
				return v < other.Cards[i]
			}
		}
	}
	return h.Kind < other.Kind
}

func (p *Problem) Part1() (result int) {
	for i := range p.Hands {
		p.Hands[i].ComputeKind(false)
	}
	sort.Slice(p.Hands[:], func(i, j int) bool {
		return p.Hands[i].Compare(p.Hands[j], false)
	})
	for i, h := range p.Hands {
		log.Debug(fmt.Sprintf("[%03d] Hand: %v, Bid: %d", i+1, h.Cards, h.Bid))
		result += (i + 1) * h.Bid
	}
	return
}

func (p *Problem) Part2() (result int) {
	for i := range p.Hands {
		p.Hands[i].ComputeKind(true)
	}
	sort.Slice(p.Hands[:], func(i, j int) bool {
		return p.Hands[i].Compare(p.Hands[j], true)
	})
	for i, h := range p.Hands {
		log.Debug(fmt.Sprintf("[%03d] Hand: %v, Bid: %d", i+1, h.Cards, h.Bid))
		result += (i + 1) * h.Bid
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
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		h := Hand{}
		var cards string
		fmt.Sscanf(l, "%s %d", &cards, &h.Bid)
		h.Parse(cards)
		p.Hands = append(p.Hands, h)
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
