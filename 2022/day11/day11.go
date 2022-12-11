package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	PLUS = iota
	MINUS
	MULT
	DIV
)

type Op struct {
	Operator, Value int
}

func (o *Op) Compute(v int) int {
	value := o.Value
	if o.Value == -1 {
		value = v
	}
	switch o.Operator {
	case PLUS:
		return v + value
	case MINUS:
		return v - value
	case MULT:
		return v * value
	case DIV:
		return v / value
	}
	return v
}

type Action struct {
	Test, ThrowTrue, ThrowFalse int
}

func (a *Action) Decide(v int) int {
	if v%a.Test == 0 {
		return a.ThrowTrue
	}
	return a.ThrowFalse
}

type Monkey struct {
	Items     []int
	Operation Op
	Test      Action
	Actions   int
}

func (m *Monkey) HasItems() bool {
	return len(m.Items) > 0
}

func (m *Monkey) InspectItem(factor int) (item, destination int) {
	m.Actions++
	log.Debug(fmt.Sprintf("Inspecting item on Monkey: %#v", m))
	item = m.Items[0]
	log.Debug(fmt.Sprintf("Inspecting item %02d", item))
	m.Items = m.Items[1:]
	log.Debug(fmt.Sprintf("New items %02d", len(m.Items)))
	item = m.Operation.Compute(item)
	if factor == 3 {
		item = int(math.Floor(float64(item) / 3.0))
	} else {
		item = item % factor
	}
	destination = m.Test.Decide(item)
	log.Debug(fmt.Sprintf("Throwing item %02d to %d", item, destination))
	return
}

type Problem struct {
	Monkeys []Monkey
	Divide3 bool
	LCM     int
}

func (p *Problem) Top(n int) (v int) {
	v = 1
	prev := 9999999
	for i := 0; i < n; i++ {
		tmp := 0
		for _, m := range p.Monkeys {
			if m.Actions > tmp && m.Actions < prev {
				tmp = m.Actions
			}
		}
		prev = tmp
		v *= tmp
	}
	return
}

func (p *Problem) Run(rounds int) {
	for i := 0; i < rounds; i++ {
		log.Debug(fmt.Sprintf("Round %02d", i))
		p.Round()
	}
}

func (p *Problem) Round() {
	factor := 3
	if !p.Divide3 {
		factor = p.LCM
	}
	for i := range p.Monkeys {
		for p.Monkeys[i].HasItems() {
			item, destination := p.Monkeys[i].InspectItem(factor)
			p.Monkeys[destination].Items = append(p.Monkeys[destination].Items, item)
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
	m := Monkey{}
	data.LCM = 1
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		if i == 0 {
			continue
		}
		if strings.HasPrefix(l, "Monkey") {
			data.Monkeys = append(data.Monkeys, m)
			m = Monkey{}
		} else if strings.HasPrefix(l, "  Starting items:") {
			parts := strings.Split(l, " ")
			log.Debug(fmt.Sprintf("[%03d] Parsing items: %v", i, parts[4:]))
			for _, n := range parts[4:] {
				n = strings.Trim(n, " ,")
				v, _ := strconv.Atoi(n)
				m.Items = append(m.Items, v)
			}
		} else if strings.HasPrefix(l, "  Operation:") {
			parts := strings.Split(l, " ")
			log.Debug(fmt.Sprintf("[%03d] Parsing Operation: %v", i, parts[5:]))
			switch parts[6] {
			case "+":
				m.Operation.Operator = PLUS
			case "-":
				m.Operation.Operator = MINUS
			case "*":
				m.Operation.Operator = MULT
			case "/":
				m.Operation.Operator = DIV
			}
			v, err := strconv.Atoi(parts[7])
			m.Operation.Value = v
			if err != nil {
				m.Operation.Value = -1
			}
		} else if strings.HasPrefix(l, "  Test:") {
			parts := strings.Split(l, " ")
			log.Debug(fmt.Sprintf("[%03d] Parsing Test: %v", i, parts[5:]))
			v, _ := strconv.Atoi(parts[5])
			m.Test.Test = v
			data.LCM *= v
		} else if strings.HasPrefix(l, "    If true:") {
			parts := strings.Split(l, " ")
			log.Debug(fmt.Sprintf("[%03d] Parsing Test if true: %v", i, parts[9:]))
			v, _ := strconv.Atoi(parts[9])
			m.Test.ThrowTrue = v
		} else if strings.HasPrefix(l, "    If false:") {
			parts := strings.Split(l, " ")
			log.Debug(fmt.Sprintf("[%03d] Parsing Test if false: %v", i, parts[9:]))
			v, _ := strconv.Atoi(parts[9])
			m.Test.ThrowFalse = v
		}
	}
	data.Monkeys = append(data.Monkeys, m)
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var rounds = context.Int("rounds")
	problem, err := parseInput(input)
	problem.Divide3 = context.Bool("div3")
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	problem.Run(rounds)
	for i, m := range problem.Monkeys {
		log.Info(fmt.Sprintf("Monkey %d thrown %03d times", i, m.Actions))
	}
	result = problem.Top(2)

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
				Name:    "rounds",
				Aliases: []string{"r"},
				Value:   20,
				Usage:   "Number of rounds, use 20 for part 1 and 10000 for part 2",
			},
			&cli.BoolFlag{
				Name:    "div3",
				Aliases: []string{"d"},
				Value:   true,
				Usage: "Divide by 3	worry level, use true for part 1 and false for part 2",
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
