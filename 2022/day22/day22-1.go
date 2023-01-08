package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
	NULL
)

const (
	WALL  = '#'
	EMPTY = '.'
)

type Coord struct {
	I, J int
}

type Direction int

type Item rune

type Position struct {
	P Coord
	D Direction
}

type Action struct {
	D     Direction
	Steps int
}

type Pair struct {
	F, L int
}

type Problem struct {
	Grid                    map[Coord]Item
	ColumnLimits, RowLimits []Pair
	Actions                 []Action
	Current                 Position
}

func (p *Problem) Traverse() {
	for i, a := range p.Actions {
		log.Debug(fmt.Sprintf("[%03d] Applying action: %#v. At %#v", i+1, a, p.Current))
		switch a.D {
		case NULL:
			direction := 1
			if p.Current.D == UP || p.Current.D == LEFT {
				direction = -1
			}
			nextC := p.Current.P
			prevC := p.Current.P
			if p.Current.D == UP || p.Current.D == DOWN {
				for j := 0; j < a.Steps; j++ {
					nextC = Coord{nextC.I + direction, p.Current.P.J}
					if _, ok := p.Grid[nextC]; !ok {
						if nextC.I < p.ColumnLimits[nextC.J].F {
							log.Debug(fmt.Sprintf("[%03d] Wrapping columm %d to %d", i+1, nextC.I, p.ColumnLimits[nextC.J].L))
							nextC.I = p.ColumnLimits[nextC.J].L
						} else {
							log.Debug(fmt.Sprintf("[%03d] Wrapping columm %d to %d", i+1, nextC.I, p.ColumnLimits[nextC.J].F))
							nextC.I = p.ColumnLimits[nextC.J].F
						}
					}
					if p.Grid[nextC] == WALL {
						log.Debug(fmt.Sprintf("[%03d] Found wall at %v", i+1, nextC))
						break
					}
					prevC = nextC
				}
			}
			if p.Current.D == RIGHT || p.Current.D == LEFT {
				for j := 0; j < a.Steps; j++ {
					nextC = Coord{p.Current.P.I, nextC.J + direction}
					if _, ok := p.Grid[nextC]; !ok {
						if nextC.J < p.RowLimits[nextC.I].F {
							log.Debug(fmt.Sprintf("[%03d] Wrapping row %d to %d", i+1, nextC.J, p.RowLimits[nextC.I].L))
							nextC.J = p.RowLimits[nextC.I].L
						} else {
							log.Debug(fmt.Sprintf("[%03d] Wrapping row %d to %d", i+1, nextC.J, p.RowLimits[nextC.I].F))
							nextC.J = p.RowLimits[nextC.I].F
						}
					}
					if p.Grid[nextC] == WALL {
						log.Debug(fmt.Sprintf("[%03d] Found wall at %v", i+1, nextC))
						break
					}
					prevC = nextC
				}
			}
			p.Current.P = prevC

		case RIGHT:
			switch p.Current.D {
			case RIGHT:
				p.Current.D = DOWN
			case LEFT:
				p.Current.D = UP
			case UP:
				p.Current.D = RIGHT
			case DOWN:
				p.Current.D = LEFT
			}
		case LEFT:
			switch p.Current.D {
			case RIGHT:
				p.Current.D = UP
			case LEFT:
				p.Current.D = DOWN
			case UP:
				p.Current.D = LEFT
			case DOWN:
				p.Current.D = RIGHT
			}
		}
		log.Debug(fmt.Sprintf("[%03d] Applied action: %#v. Now at %#v", i+1, a, p.Current))
	}
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

func atoiSafe(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
}

func mod(a, b int) int {
	return (a%b + b) % b
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

func parseInput(input string) (data Problem, err error) {
	byteData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	strData := string(byteData)
	data.Grid = make(map[Coord]Item)
	data.Actions = make([]Action, 0)
	actionRegex := regexp.MustCompile(`(\d+)?([LR])(\d+)`)
	rows, columns := 0, 0
	for i, l := range strings.Split(strData, "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		// Last line
		if len(l) > 0 && (l[0] == 'L' || l[0] == 'R' || (l[0] >= '0' && l[0] <= '9')) {
			matches := actionRegex.FindAllSubmatch([]byte(l), -1)
			log.Debug(fmt.Sprintf("Parsing actions: %s", matches))
			for _, m := range matches {
				for j := 1; j < len(m); j++ {
					if len(m[j]) == 0 {
						continue
					}
					v, err := strconv.Atoi(string(m[j]))
					a := Action{}
					if err == nil {
						a.Steps = v
						a.D = NULL
					} else if rune(m[j][0]) == 'L' {
						a.D = LEFT
					} else if rune(m[j][0]) == 'R' {
						a.D = RIGHT
					}
					data.Actions = append(data.Actions, a)
				}
			}
		} else if len(l) > 0 {
			rows = i
			rLimit := Pair{999999, 0}
			for j, c := range l {
				if c == WALL || c == EMPTY {
					if i == 0 && data.Current.P.J == 0 {
						data.Current = Position{
							Coord{i, j},
							RIGHT,
						}
					}
					if j > rLimit.L {
						rLimit.L = j
					}
					if j < rLimit.F {
						rLimit.F = j
					}
					if j > columns {
						columns = j
					}
					data.Grid[Coord{i, j}] = Item(c)
				}
			}
			data.RowLimits = append(data.RowLimits, rLimit)
		}
	}
	for i := 0; i <= columns; i++ {
		cLimit := Pair{999999, 0}
		for j := 0; j <= rows; j++ {
			if _, ok := data.Grid[Coord{j, i}]; ok {
				if j < cLimit.F {
					cLimit.F = j
				}
				if j > cLimit.L {
					cLimit.L = j
				}
			}
		}
		data.ColumnLimits = append(data.ColumnLimits, cLimit)
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	problem.Traverse()
	result = (problem.Current.P.I+1)*1000 + (problem.Current.P.J+1)*4 + int(problem.Current.D)

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
			start := time.Now()
			echo(fmt.Sprintf("Solution is %v in %s", solution(c), time.Since(start)), c.String("output"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
