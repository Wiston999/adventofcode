package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Layer struct {
	Range, Current, Direction int
}

type Problem struct {
	Layers map[int]Layer
	Max    int
}

func (p *Problem) Step() {
	for l := range p.Layers {
		layer := p.Layers[l]
		if layer.Current == 0 {
			layer.Direction = 1
		}
		if layer.Current == (layer.Range - 1) {
			layer.Direction = -1
		}
		layer.Current = layer.Current + layer.Direction
		p.Layers[l] = layer
	}
}

func (p *Problem) Copy() (cp Problem) {
	cp.Max = p.Max
	cp.Layers = make(map[int]Layer)
	for k, v := range p.Layers {
		cp.Layers[k] = Layer{
			Range:     v.Range,
			Current:   v.Current,
			Direction: v.Direction,
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
	data.Layers = make(map[int]Layer)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		parts := strings.Split(l, ": ")
		n, _ := strconv.Atoi(parts[1])
		layer := Layer{Range: n, Direction: 1}
		n, _ = strconv.Atoi(parts[0])

		if n > data.Max {
			data.Max = n
		}
		data.Layers[n] = layer
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var delay = context.Int("delay")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	for j := 0; j < delay; j++ {
		current := problem.Copy()
		caught := false
		for i := 0; i <= current.Max; i++ {
			if l, ok := current.Layers[i]; ok {
				if l.Current == 0 {
					log.Debug(fmt.Sprintf("[%03d][%03d] Caught at layer %v", j, i, l))
					caught = true
					break
				}
			}
			current.Step()
		}
		problem.Step()
		if !caught {
			result = j
			break
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
			&cli.IntFlag{
				Name:    "delay",
				Aliases: []string{"d"},
				Value:   10000000,
				Usage:   "Max delay to be checked",
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
