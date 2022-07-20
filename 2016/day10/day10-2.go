package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type ProblemData struct {
	Bots    map[string]*Edge
	Inputs  map[string]*Edge
	Outputs map[string]*Edge
}

var KindNames = []string{
	"Input",
	"Bot",
	"Output",
}

const (
	InputEdge = iota
	BotEdge
	OutputEdge
)

type EdgeKind int

type Edge struct {
	Name                string
	Kind                EdgeKind
	Low, High           *Edge
	LowValue, HighValue int
}

func (e *Edge) SetValue(v int) {
	// Set in empty slot
	if v != e.LowValue && v != e.HighValue {
		if e.LowValue == 0 {
			e.LowValue = v
		} else if e.HighValue == 0 {
			e.HighValue = v
		}
		// Check and reorder
		if e.HighValue < e.LowValue {
			tmp := e.LowValue
			e.LowValue = e.HighValue
			e.HighValue = tmp
		}
	}
	log.Debug(fmt.Sprintf("Updated Edge (%s %s) [%p] %#v with %d", KindNames[e.Kind], e.Name, e, e, v))
}

func (e *Edge) Completed() bool {
	if e.Kind == InputEdge {
		return true
	} else if e.Kind == BotEdge {
		return e.LowValue != 0 && e.HighValue != 0
	} else if e.Kind == OutputEdge {
		return e.HighValue != 0
	}
	return false
}

func (e *Edge) SetOutputs() {
	if e.Completed() {
		e.Low.SetValue(e.LowValue)
		e.High.SetValue(e.HighValue)
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

func atoi(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
}

func getOrCreate(m map[string]*Edge, key string) *Edge {
	if _, ok := m[key]; !ok {
		m[key] = new(Edge)
	}
	e := m[key]
	e.Name = key
	return e
}

func parseInput(input string) (data ProblemData, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	data.Bots = make(map[string]*Edge)
	data.Outputs = make(map[string]*Edge)
	data.Inputs = make(map[string]*Edge)

	valueRegex := *regexp.MustCompile(`^value (\d+) goes to bot (\d+)$`)
	outputRegex := *regexp.MustCompile(`^bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)$`)
	for i, l := range strings.Split(strings.TrimSpace(string(tmpData)), "\n") {
		l = strings.TrimSpace(l)
		if values := valueRegex.FindStringSubmatch(l); len(values) > 0 {
			input := getOrCreate(data.Inputs, values[1])
			bot := getOrCreate(data.Bots, values[2])

			bot.Kind = BotEdge
			input.Kind = InputEdge
			input.HighValue = atoi(values[1])
			input.Low = bot
			input.High = bot

			data.Bots[values[2]] = bot
			data.Inputs[values[1]] = input
			log.Debug(fmt.Sprintf("(%03d) Parsed input %s for bot %s: %#v <-- %#v", i, values[1], values[2], bot, input))
		}
		if outputs := outputRegex.FindStringSubmatch(l); len(outputs) > 0 {
			bot := getOrCreate(data.Bots, outputs[1])
			bot.Kind = BotEdge

			var lowEdge, highEdge *Edge
			if outputs[2] == "bot" {
				lowEdge = getOrCreate(data.Bots, outputs[3])
				lowEdge.Kind = BotEdge
				data.Bots[outputs[3]] = lowEdge
			} else {
				lowEdge = getOrCreate(data.Outputs, outputs[3])
				lowEdge.Kind = OutputEdge
				data.Outputs[outputs[3]] = lowEdge
			}
			if outputs[4] == "bot" {
				highEdge = getOrCreate(data.Bots, outputs[5])
				highEdge.Kind = BotEdge
				data.Bots[outputs[5]] = highEdge
			} else {
				highEdge = getOrCreate(data.Outputs, outputs[5])
				highEdge.Kind = OutputEdge
				data.Outputs[outputs[5]] = highEdge
			}
			bot.High = highEdge
			bot.Low = lowEdge
			data.Bots[outputs[1]] = bot
			log.Debug(fmt.Sprintf("(%03d) Updated Bot %s outputs: %#v", i, outputs[1], bot))
		}
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	data, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	targets := context.StringSlice("target")
	log.Info(fmt.Sprintf("Looking product of outputs %v", targets))
	for k, v := range data.Inputs {
		log.Debug(fmt.Sprintf("Input %s (%p): %#v", k, v, *v))
	}
	for k, v := range data.Bots {
		log.Debug(fmt.Sprintf("Bot %s (%p): %#v", k, v, *v))
	}
	for k, v := range data.Outputs {
		log.Debug(fmt.Sprintf("Output %s (%p): %#v", k, v, *v))
	}
	for _, v := range data.Inputs {
		v.SetOutputs()
	}
	pending := map[string]bool{}
	for k := range data.Bots {
		pending[k] = true
	}
	for len(pending) > 0 {
		for k, v := range data.Bots {
			if _, ok := pending[k]; !ok {
				continue
			}
			if v.Completed() {
				v.SetOutputs()
				delete(pending, k)
			}
		}
	}
	result = 1
	for _, t := range targets {
		result *= data.Outputs[t].HighValue
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
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Value:   "-",
				Usage:   "Output file path, use - for stdout",
			},
			&cli.StringSliceFlag{
				Name:    "target",
				Aliases: []string{"t"},
				Value:   cli.NewStringSlice("0", "1", "2"),
				Usage:   "Output targets to compute their product",
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
