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
	HUMAN
)

type OP int

type Monkey struct {
	Name               string
	Value              int
	Operation          OP
	Operand1, Operand2 *Monkey
	Result             Result
}

func (m *Monkey) HumanValue(constraint int) (value int) {
	if m.Operation == HUMAN {
		log.Info(fmt.Sprintf("Base case: %#v --> %d", m, constraint))
		return constraint
	}
	var fixed int
	var branch *Monkey
	if m.Operand1.Result.Human {
		fixed = m.Operand2.Result.Value
		branch = m.Operand1
	}
	if m.Operand2.Result.Human {
		fixed = m.Operand1.Result.Value
		branch = m.Operand2
	}
	switch m.Operation {
	case ADD:
		log.Debug(fmt.Sprintf("[%#v] Applying ADD contraint: %d", m, m.Result.Value-fixed))
		value += branch.HumanValue(constraint - fixed)
	case MINUS:
		if m.Operand1.Result.Human {
			log.Debug(fmt.Sprintf("[%#v] Applying MINUS1 contraint: %d", m, m.Result.Value+fixed))
			value += branch.HumanValue(constraint + fixed)
		} else {
			log.Debug(fmt.Sprintf("[%#v] Applying MINUS2 contraint: %d", m, fixed-m.Result.Value))
			value += branch.HumanValue(fixed - constraint)
		}
	case MULT:
		log.Debug(fmt.Sprintf("[%#v] Applying MULT contraint: %d", m, m.Result.Value/fixed))
		value += branch.HumanValue(constraint / fixed)
	case DIV:
		if m.Operand1.Result.Human {
			log.Debug(fmt.Sprintf("[%#v] Applying DIV1 contraint: %d", m, m.Result.Value*fixed))
			value += branch.HumanValue(constraint * fixed)
		} else {
			log.Debug(fmt.Sprintf("[%#v] Applying DIV1 contraint: %d", m, fixed/m.Result.Value))
			value += branch.HumanValue(fixed / m.Result.Value)
		}
	}
	return
}

type Result struct {
	Value int
	Human bool
}

func NewResult(operand OP, op1, op2 Result) (r Result) {
	switch operand {
	case ADD:
		r.Value = op1.Value + op2.Value
	case MINUS:
		r.Value = op1.Value - op2.Value
	case MULT:
		r.Value = op1.Value * op2.Value
	case DIV:
		r.Value = op1.Value / op2.Value
	}
	r.Human = op1.Human || op2.Human
	return r
}

type Problem struct {
	Monkeys map[string]*Monkey
}

func (p *Problem) Compute(m *Monkey) (result Result) {
	switch m.Operation {
	case VALUE:
		result = Result{Value: m.Value}
	case HUMAN:
		result = Result{Value: m.Value, Human: true}
	default:
		result = NewResult(m.Operation, p.Compute(m.Operand1), p.Compute(m.Operand2))
	}
	if result.Human {
		log.Debug(fmt.Sprintf("Human branch detected: %#v (%#v)", m, result))
	}
	m.Result = result
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
		m.Name = parts[0]
		if v, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
			m.Value = v
			m.Operation = VALUE
			if parts[0] == "humn" {
				m.Operation = HUMAN
			}
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
	root := problem.Monkeys["root"]
	// Traverse tree to compute solution
	problem.Compute(root)
	// Reverse traverse matching constraints
	if root.Operand1.Result.Human {
		result = root.Operand1.HumanValue(root.Operand2.Result.Value)
	} else {
		result = root.Operand2.HumanValue(root.Operand1.Result.Value)
	}
	// result = root.HumanValue(0)

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
