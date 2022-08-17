package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Coord struct {
	X, Y int
}

type void struct{}

type Problem struct {
	Current Coord
	Steps   []string
}

func (p *Problem) MoveStep(step string) {
	switch step {
	case "ne":
		p.Current.X++
		p.Current.Y--
	case "se":
		p.Current.X++
	case "nw":
		p.Current.X--
	case "sw":
		p.Current.X--
		p.Current.Y++
	case "n":
		p.Current.Y--
	case "s":
		p.Current.Y++
	}
}

func (p *Problem) Move() Coord {
	for i, s := range p.Steps {
		p.MoveStep(s)
		log.Debug(fmt.Sprintf("[%03d] After step %s: %v", i, s, p.Current))
	}
	return p.Current
}

// Manhattan
func (p *Problem) Distance(c Coord) int {
	return int((math.Abs(float64(c.X-p.Current.X)) +
		math.Abs(float64(c.Y-p.Current.Y)) +
		math.Abs(float64((-c.X-c.Y)-(-p.Current.X-p.Current.Y)))) / 2,
	)
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
	for i, l := range strings.Split(strings.TrimSpace(strData), ",") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		data.Steps = append(data.Steps, l)
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

	start := Coord{0, 0}
	problem.Current = start
	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	problem.Move()
	log.Info(fmt.Sprintf("Final position after %d steps: %v", len(problem.Steps), problem.Current))
	result = int(problem.Distance(start))

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
