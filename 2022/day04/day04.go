package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Section struct {
	Start, End int
}

func (s *Section) Parse(data string) {
	edges := strings.Split(data, "-")
	st, _ := strconv.Atoi(edges[0])
	e, _ := strconv.Atoi(edges[1])
	s.Start = st
	s.End = e
}

type Pair struct {
	S1, S2 Section
}

func (p *Pair) Contained() (result bool) {
	result = p.S1.Start >= p.S2.Start && p.S1.End <= p.S2.End
	result = result || p.S2.Start >= p.S1.Start && p.S2.End <= p.S1.End
	return
}

func (p *Pair) Overlap() (result bool) {
	result = p.S1.Start >= p.S2.Start && p.S1.Start <= p.S2.End
	result = result || p.S1.End >= p.S2.Start && p.S1.End <= p.S2.End
	result = result || p.S2.Start >= p.S1.Start && p.S2.Start <= p.S1.End
	result = result || p.S2.End >= p.S1.Start && p.S2.End <= p.S1.End
	return
}

type Problem struct {
	Pairs []Pair
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
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		pairs := strings.Split(l, ",")
		S1, S2 := Section{}, Section{}
		S1.Parse(pairs[0])
		S2.Parse(pairs[1])
		pair := Pair{S1, S2}
		data.Pairs = append(data.Pairs, pair)
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var overlap = context.Bool("overlap")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	for i, p := range problem.Pairs {
		if !overlap && p.Contained() {
			log.Debug(fmt.Sprintf("[%03d] %v is contained", i, p))
			result += 1
		} else if overlap && p.Overlap() {
			log.Debug(fmt.Sprintf("[%03d] %v is overlapped", i, p))
			result += 1

		}
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
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   "input.txt",
				Usage:   "Input file path",
			},
			&cli.BoolFlag{
				Name:    "overlap",
				Aliases: []string{"p"},
				Value:   false,
				Usage:   "Compute overlaps instead of contained (part 2)",
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
