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
	VALUE = iota
	ADD
	MINUS
	MULT
	DIV
)

type OP int

type Monkey struct {
	Value              int
	Operation          OP
	Operand1, Operand2 *Monkey
}

type Problem struct {
	Monkeys map[string]*Monkey
}

func (p *Problem) Compute(m *Monkey) (result int) {
	log.Debug(fmt.Sprintf("Computing monkey %#v", *m))
	switch m.Operation {
	case VALUE:
		result = m.Value
	case ADD:
		result = p.Compute(m.Operand1) + p.Compute(m.Operand2)
	case MINUS:
		result = p.Compute(m.Operand1) - p.Compute(m.Operand2)
	case MULT:
		result = p.Compute(m.Operand1) * p.Compute(m.Operand2)
	case DIV:
		result = p.Compute(m.Operand1) / p.Compute(m.Operand2)
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
	data.Monkeys = make(map[string]*Monkey)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		var m *Monkey
		ok := false
		parts := strings.Split(l, ":")
		if m, ok = data.Monkeys[parts[0]]; !ok {
			m = new(Monkey)
			data.Monkeys[parts[0]] = m
		}
		if v, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
			m.Value = v
			m.Operation = VALUE
		} else {
			operation := strings.Split(strings.TrimSpace(parts[1]), " ")
			switch operation[1] {
			case "+":
				m.Operation = ADD
			case "-":
				m.Operation = MINUS
			case "*":
				m.Operation = MULT
			case "/":
				m.Operation = DIV
			}
			if m1, ok := data.Monkeys[operation[0]]; ok {
				m.Operand1 = m1
			} else {
				m1 = new(Monkey)
				data.Monkeys[operation[0]] = m1
				m.Operand1 = m1
			}
			if m2, ok := data.Monkeys[operation[2]]; ok {
				m.Operand2 = m2
			} else {
				m2 = new(Monkey)
				data.Monkeys[operation[2]] = m2
				m.Operand2 = m2
			}
		}
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
	result = problem.Compute(problem.Monkeys["root"])

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
