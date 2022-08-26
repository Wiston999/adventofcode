package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	UP    = "up"
	DOWN  = "down"
	LEFT  = "left"
	RIGHT = "right"
)

type direction string

type Coord struct {
	I, J int
}

type Carrier struct {
	Position  Coord
	Direction direction
}

func (c *Carrier) Move() {
	switch c.Direction {
	case UP:
		c.Position.I--
	case DOWN:
		c.Position.I++
	case RIGHT:
		c.Position.J++
	case LEFT:
		c.Position.J--
	}
}

func (c *Carrier) Turn(d direction) {
	switch d {
	case LEFT:
		switch c.Direction {
		case UP:
			c.Direction = LEFT
		case DOWN:
			c.Direction = RIGHT
		case RIGHT:
			c.Direction = UP
		case LEFT:
			c.Direction = DOWN
		}
	case RIGHT:
		switch c.Direction {
		case UP:
			c.Direction = RIGHT
		case DOWN:
			c.Direction = LEFT
		case RIGHT:
			c.Direction = DOWN
		case LEFT:
			c.Direction = UP
		}
	}
}

type Problem struct {
	Map   map[Coord]bool
	Virus Carrier
}

func (p *Problem) Burst() (infection int) {
	_, infected := p.Map[p.Virus.Position]
	if infected {
		p.Virus.Turn(RIGHT)
		delete(p.Map, p.Virus.Position)
	} else {
		p.Virus.Turn(LEFT)
		p.Map[p.Virus.Position] = true
		infection = 1
	}
	p.Virus.Move()
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
	m := make(map[Coord]bool)
	maxI, maxJ := 0, 0
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			if c == '#' {
				m[Coord{i, j}] = true
			}
			if j > maxJ {
				maxJ = j
			}
		}
		if i > maxI {
			maxI = i
		}
	}
	data.Map = make(map[Coord]bool)
	for c := range m {
		data.Map[Coord{c.I - (maxI / 2), c.J - (maxJ / 2)}] = true
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var bursts = context.Int("bursts")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	problem.Virus = Carrier{Coord{0, 0}, UP}
	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	for i := 0; i < bursts; i++ {
		result += problem.Burst()
		log.Debug(fmt.Sprintf("[%05d] Burst, virus (%v), current solution: %d", i, problem.Virus, result))
	}

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
				Name:    "bursts",
				Aliases: []string{"b"},
				Value:   10000,
				Usage:   "Number of virus bursts",
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
