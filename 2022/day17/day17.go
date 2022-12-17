package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	LEFT = iota
	RIGHT
)

const (
	AIR = iota
	ROCK
)

type JetAction int

type Element int

type Rock map[Coord]Element

type Coord struct {
	X, Y int
}

var Rocks = []Rock{
	Rock{
		Coord{0, 0}: ROCK, Coord{0, 1}: ROCK, Coord{0, 2}: ROCK, Coord{0, 3}: ROCK,
	},
	Rock{
		Coord{0, 0}: AIR, Coord{0, 1}: ROCK, Coord{0, 2}: AIR,
		Coord{-1, 0}: ROCK, Coord{-1, 1}: ROCK, Coord{-1, 2}: ROCK,
		Coord{-2, 0}: AIR, Coord{-2, 1}: ROCK, Coord{-2, 2}: AIR,
	},
	Rock{
		Coord{0, 0}: AIR, Coord{0, 1}: AIR, Coord{0, 2}: ROCK,
		Coord{-1, 0}: AIR, Coord{-1, 1}: AIR, Coord{-1, 2}: ROCK,
		Coord{-2, 0}: ROCK, Coord{-2, 1}: ROCK, Coord{-2, 2}: ROCK,
	},
	Rock{
		Coord{0, 0}:  ROCK,
		Coord{-1, 0}: ROCK,
		Coord{-2, 0}: ROCK,
		Coord{-3, 0}: ROCK,
	},
	Rock{
		Coord{0, 0}: ROCK, Coord{0, 1}: ROCK,
		Coord{-1, 0}: ROCK, Coord{-1, 1}: ROCK,
	},
}

type Problem struct {
	Actions                          []JetAction
	Grid                             map[Coord]Element
	LineTurn, HeightTurn, JetTurn    map[int]int
	RockIndex, JetIndex              int
	MaxX                             int
	PeriodLine, PeriodTurn           int
	PeriodLineStart, PeriodTurnStart int
}

func (p *Problem) RockWidth(rock Rock) (result int) {
	for c := range rock {
		if c.Y > result {
			result = c.Y
		}
	}
	return
}

func (p *Problem) RockHeight(rock Rock) (result int) {
	for c := range rock {
		if c.X < result {
			result = c.X
		}
	}
	return -result
}

func (p *Problem) PrintLine(i int) (result string) {
	for j := 0; j < 7; j++ {
		if p.Grid[Coord{i, j}] == AIR {
			result += "."
		} else {
			result += "#"
		}
	}
	return
}

func (p *Problem) CompareLines(i, j int) bool {
	for k := 0; k < 7; k++ {
		if p.Grid[Coord{i, k}] != p.Grid[Coord{j, k}] {
			return false
		}
	}
	return true
}

func (p *Problem) Print() (result string) {
	for i := p.MaxX + 3; i >= 0; i-- {
		result += fmt.Sprintf("%03d ", i)
		result += p.PrintLine(i) + "\n"
	}
	return
}

func (p *Problem) CanMoveHorizontal(rockCoord Coord, rock Rock, m int) (result bool) {
	result = true
	for c, v := range rock {
		if v == ROCK {
			nextCoord := Coord{rockCoord.X + c.X, rockCoord.Y + c.Y + m}
			if p.Grid[nextCoord] != AIR || (nextCoord.Y < 0) || (nextCoord.Y >= 7) {
				// log.Debug(fmt.Sprintf("Rock %v cannot move %v = %d [%d]", rock, nextCoord, p.Grid[nextCoord], m))
				result = false
				break
			}
		}
	}
	return
}

func (p *Problem) DropStone(turn int) {
	rock := Rocks[p.RockIndex%len(Rocks)]
	p.RockIndex++
	rockHeight := p.RockHeight(rock)
	// rockWidth := p.RockWidth(rock)
	rockCoord := Coord{p.MaxX + rockHeight + 3, 2}
	falling := true

	for falling {
		action := p.Actions[p.JetIndex%len(p.Actions)]
		p.JetIndex++
		if action == LEFT {
			// log.Debug(fmt.Sprintf("Pushing to LEFT"))
			if p.CanMoveHorizontal(rockCoord, rock, -1) {
				// log.Debug(fmt.Sprintf("Rock %v @ %v move to LEFT", rock, rockCoord))
				rockCoord.Y -= 1
			}
		} else {
			// log.Debug(fmt.Sprintf("Pushing to RIGHT"))
			if p.CanMoveHorizontal(rockCoord, rock, 1) {
				// log.Debug(fmt.Sprintf("Rock %v @ %v move to RIGHT", rock, rockCoord))
				rockCoord.Y += 1
			}
		}

		for c, v := range rock {
			if v == ROCK {
				cDown := Coord{rockCoord.X + c.X - 1, rockCoord.Y + c.Y}
				if p.Grid[cDown] != AIR || (rockCoord.X+c.X-1 < -p.MaxX) {
					// log.Debug(fmt.Sprintf("Rock %v cannot fall %v = %d [%d]", rock, cDown, p.Grid[cDown], p.MaxX))
					falling = false
					break
				}
			}
		}

		if falling {
			// log.Debug(fmt.Sprintf("Rock %v falling 1 step", rock))
			rockCoord.X -= 1
		}
	}

	for c, v := range rock {
		if v == ROCK {
			p.Grid[Coord{rockCoord.X + c.X, rockCoord.Y + c.Y}] = ROCK
			p.LineTurn[rockCoord.X+c.X] = turn
			if rockCoord.X+c.X+1 > p.MaxX {
				p.MaxX = rockCoord.X + c.X + 1
			}
		}
	}

	if log.GetLevel() == log.DebugLevel {
		log.Debug(fmt.Sprintf("Rock %v location %v", rock, rockCoord))
		log.Debug(fmt.Sprintf("Current cave [depth=%d]", p.MaxX))
	}

	p.HeightTurn[turn] = p.MaxX
	p.JetTurn[turn] = p.JetIndex % len(p.Actions)
	if p.PeriodLine == 0 {
		p.FindPeriod(turn)
	}
}

func (p *Problem) FindPeriod(turn int) {
	i := p.MaxX - 1
	for j := i - 1; j >= 0; j-- {
		if p.CompareLines(i, j) {
			repeat := true
			for k := 0; k < i-j; k++ {
				if !p.CompareLines(i-k, j-k) {
					repeat = false
					break
				}
			}
			if repeat && p.JetTurn[i] == p.JetTurn[j] && (i-j) > 10 {
				p.PeriodLine = i - j
				p.PeriodLineStart = j - p.PeriodLine + 1
				p.PeriodTurnStart = p.LineTurn[j-p.PeriodLine]
				p.PeriodTurn = p.LineTurn[j] - p.PeriodTurnStart
				log.Info(fmt.Sprintf(
					"[%03d] Found period from line %d to %d [L=%d T=%d] starting at line %d in turn %d",
					turn, i, j, p.PeriodLine, p.PeriodTurn, p.PeriodLineStart, p.PeriodTurnStart,
				))
				return
			}
		}
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
	data.Actions = make([]JetAction, 0)
	data.Grid = make(map[Coord]Element)
	data.LineTurn = make(map[int]int)
	data.HeightTurn = make(map[int]int)
	data.JetTurn = make(map[int]int)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for _, c := range strings.TrimSpace(l) {
			if c == '<' {
				data.Actions = append(data.Actions, LEFT)
			} else {
				data.Actions = append(data.Actions, RIGHT)
			}
		}
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var steps = context.Int("steps")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))

	for i := 1; i <= steps && problem.PeriodLine == 0; i++ {
		problem.DropStone(i)
	}

	if problem.PeriodLine > 0 {
		turns := steps - problem.PeriodTurnStart
		periods := turns / problem.PeriodTurn
		offset := turns % problem.PeriodTurn
		offsetHeight := problem.HeightTurn[offset+problem.PeriodTurnStart] - problem.HeightTurn[problem.PeriodTurnStart]
		problem.MaxX = problem.PeriodLineStart + (periods * problem.PeriodLine) + offsetHeight
	}

	result = problem.MaxX
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
			&cli.IntFlag{
				Name:    "steps",
				Aliases: []string{"s"},
				Value:   2022,
				Usage:   "Number of rocks to be simulated, use 2022 for part 1 and 1000000000000 for part 2",
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
