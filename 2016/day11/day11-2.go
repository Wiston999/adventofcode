package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"

	astar "github.com/fzipp/astar"
	log "github.com/sirupsen/logrus"

	combinations "github.com/mxschmitt/golang-combinations"

	"github.com/urfave/cli/v2"
)

const FloorSep = "-"

type State struct {
	Floors  string
	Current int
}

func (s *State) Sort() {
	floors := strings.Split(s.Floors, FloorSep)
	for i := range floors {
		f := []rune(floors[i])
		sort.Slice(
			f,
			func(m int, n int) bool { return f[m] < f[n] },
		)
		floors[i] = string(f)
	}
	s.Floors = strings.Join(floors, FloorSep)
}

func (s *State) Valid() (result bool) {
	floors := strings.Split(s.Floors, FloorSep)
	for i := range floors {
		f := []rune(floors[i])
		var gens, chips []rune
		for j := range f {
			c := f[j]
			if c >= 'A' && c <= 'Z' {
				gens = append(gens, c)
			} else {
				chips = append(chips, unicode.ToUpper(c))
			}
		}
		if len(gens) > 0 && len(chips) > 0 {
			// Look for impaired chips
			for j := range chips {
				c := chips[j]
				paired := false
				for k := range gens {
					if c == gens[k] {
						paired = true
						break
					}
				}
				if !paired {
					return false
				}
			}
		}
	}
	return true
}

type void struct{}

type Problem struct {
	Visited  map[State]void
	Elements []rune
}

func (p Problem) Neighbours(s State) (ls []State) {
	floors := strings.Split(s.Floors, FloorSep)
	floorElements := strings.Split(floors[s.Current], "")
	combis := append(combinations.Combinations(floorElements, 2), combinations.Combinations(floorElements, 1)...)
	stateMap := make(map[State]bool)
	worthy := true
	if s.Current > 0 && worthy {
		for i := range combis {
			ns := State{Current: s.Current - 1, Floors: ""}
			ns.Floors = strings.Join(floors[:ns.Current], FloorSep)
			if ns.Current > 0 {
				ns.Floors += FloorSep
			}
			ns.Floors += floors[ns.Current]
			ns.Floors += strings.Join(combis[i], "") + FloorSep
			cFloor := floors[s.Current]
			for j := range combis[i] {
				cFloor = strings.ReplaceAll(cFloor, combis[i][j], "")
			}
			ns.Floors += cFloor
			if s.Current < len(floors)-1 {
				ns.Floors += FloorSep
				ns.Floors += strings.Join(floors[s.Current+1:], FloorSep)
			}
			ns.Sort()
			stateMap[ns] = true
		}
	}
	if s.Current < len(floors)-1 {
		for i := range combis {
			ns := State{Current: s.Current + 1, Floors: ""}
			ns.Floors = strings.Join(floors[:ns.Current-1], FloorSep)
			if s.Current > 0 {
				ns.Floors += FloorSep
			}
			cFloor := floors[s.Current]
			for j := range combis[i] {
				cFloor = strings.ReplaceAll(cFloor, combis[i][j], "")
			}
			ns.Floors += cFloor + FloorSep
			ns.Floors += strings.Join(combis[i], "")
			ns.Floors += strings.Join(floors[ns.Current:], FloorSep)
			ns.Sort()
			stateMap[ns] = true
		}
	}
	for ns := range stateMap {
		if ns.Valid() {
			p.Sort(&ns)
			if _, ok := p.Visited[ns]; !ok {
				p.Visited[ns] = void{}
				ls = append(ls, ns)
			}
		}
	}
	return
}

func swap(a, b, c string) (z string) {
	z = strings.ReplaceAll(a, b, "z")
	z = strings.ReplaceAll(z, c, b)
	z = strings.ReplaceAll(z, "z", c)
	return
}

// Ensure uppercases are always sorted between them
func (p *Problem) Sort(s *State) {
	var orderUpper, orderLower []int
	for i, c := range s.Floors {
		if c >= 'A' && c <= 'Z' {
			orderUpper = append(orderUpper, i)
			orderLower = append(orderLower, strings.Index(s.Floors, strings.ToLower(string(c))))
		}
	}

	for i, j := range p.Elements {
		s.Floors = s.Floors[:orderUpper[i]] + string(j) + s.Floors[orderUpper[i]+1:]
		s.Floors = s.Floors[:orderLower[i]] + strings.ToLower(string(j)) + s.Floors[orderLower[i]+1:]
	}
	s.Sort()

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

func parseInput(input string) (data State, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	translation := make(map[string]string)
	lines := strings.Split(strings.TrimSpace(string(tmpData)), "\n")
	for i, l := range lines {
		floor := ""
		for j, p := range strings.Split(strings.TrimSpace(string(l)), " a ") {
			if j > 0 { // First part has no useful information, suppossing floors are sorted at input
				words := strings.Split(p, "-")
				words = strings.Split(words[0], ",")
				words = strings.Split(words[0], ".")
				words = strings.Split(words[0], " ")
				item, ok := translation[words[0]]
				log.Debug(fmt.Sprintf("Parsed %s element, already present? %v", words[0], ok))
				if !ok {
					max := '@'
					for _, v := range translation {
						if rune(v[0]) > max {
							max = rune(v[0])
						}
					}
					item = string(max + 1)
					log.Debug(fmt.Sprintf("Assigned letter %v to %s", item, words[0]))
					translation[words[0]] = item
				}
				if strings.Contains(p, "generator") {
					item = strings.ToUpper(item)
				} else {
					item = strings.ToLower(item)
				}
				floor += item
			}
		}
		data.Floors += floor
		if i < len(lines)-1 {
			data.Floors += FloorSep
		}
	}
	return
}

func cost(e, s State) (c float64) {
	// eFloors := strings.Split(e.Floors, FloorSep)
	// c = float64(3 - (len(eFloors[e.Current]) - len(eFloors[s.Current])))
	c = 1
	return
}

func heuristic(s, e State) (c float64) {
	// eFloors := strings.Split(e.Floors, FloorSep)
	// for i, j := 0, len(eFloors)-1; i < len(eFloors); i, j = i+1, j-1 {
	// 	c += float64(len(eFloors[j]) * i)
	// }
	return
}

func addElements(o string, e rune) (c string) {
	c = o
	c += string(e + 1)
	c += string(e + 2)
	c += strings.ToLower(string(e + 1))
	c += strings.ToLower(string(e + 2))
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	start, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	start.Current = 0
	target := start
	target.Current = 3
	target.Floors = strings.Repeat(FloorSep, 3)
	problem := Problem{Visited: make(map[State]void)}
	for _, f := range strings.Split(start.Floors, FloorSep) {
		for _, c := range f {
			if c >= 'A' && c <= 'Z' {
				problem.Elements = append(problem.Elements, c)
			}
		}
		target.Floors += f
	}
	// Add new elements to start and target
	lastRune := problem.Elements[len(problem.Elements)-1]
	floors := strings.Split(start.Floors, FloorSep)
	floors[0] = addElements(floors[0], lastRune)
	start.Floors = strings.Join(floors, FloorSep)

	target.Floors = addElements(target.Floors, lastRune)

	problem.Elements = append(problem.Elements, lastRune+1)
	problem.Elements = append(problem.Elements, lastRune+2)
	problem.Sort(&start)
	problem.Sort(&target)
	log.Info(fmt.Sprintf("Searching path from\n%#v to\n%#v", start, target))
	path := astar.FindPath[State](problem, start, target, cost, heuristic)
	log.Info(fmt.Sprintf("Found solution path: %v", path))
	result = len(path) - 1
	return
}

func main() {
	app := &cli.App{
		Name:  "AOC",
		Usage: "Solve AdventOfCode problem!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "loglevel",
				Aliases: []string{"l"},
				Value:   "info",
				Usage:   "Log level output",
			},
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   "input.txt",
				Usage:   "Input file path",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Value:   "-",
				Usage:   "Output file path, use - for stdout",
			},
		},
		Action: func(c *cli.Context) error {
			setLogLevel(c.String("loglevel"))
			log.Debug("Received parameters:")
			log.Debug(fmt.Sprintf("Input file name:  %s", c.String("input")))
			log.Debug(fmt.Sprintf("Output file name: %s", c.String("output")))
			log.Debug(fmt.Sprintf("Log level:        %s", c.String("loglevel")))
			echo(fmt.Sprintf("Solution is %v", solution(c)), c.String("output"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
