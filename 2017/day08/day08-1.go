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

const (
	Dec = "dec"
	Inc = "inc"
	LE  = "le"
	LT  = "lt"
	GE  = "ge"
	GT  = "gt"
	EQ  = "eq"
	NE  = "ne"
)

type Operation string

type Register int

type Instruction struct {
	Register  string
	Op        Operation
	Value     int
	Condition *Instruction
}

func (i *Instruction) Process(registers map[string]int) int {
	condition := false

	switch i.Condition.Op {
	case LT:
		condition = registers[i.Condition.Register] < i.Condition.Value
	case LE:
		condition = registers[i.Condition.Register] <= i.Condition.Value
	case GT:
		condition = registers[i.Condition.Register] > i.Condition.Value
	case GE:
		condition = registers[i.Condition.Register] >= i.Condition.Value
	case EQ:
		condition = registers[i.Condition.Register] == i.Condition.Value
	case NE:
		condition = registers[i.Condition.Register] != i.Condition.Value
	}

	log.Debug(fmt.Sprintf("Applying %v %#v if %#v [%d %d]", condition, i, i.Condition, registers[i.Register], registers[i.Condition.Register]))
	if condition {
		if i.Op == Dec {
			return -i.Value
		}
		return i.Value
	}
	return 0
}

type Problem struct {
	Registers    map[string]int
	Instructions []Instruction
	PC           int
}

func (p *Problem) Process() {
	for p.PC = 0; p.PC < len(p.Instructions); p.PC++ {
		v := p.Instructions[p.PC].Process(p.Registers)
		p.Registers[p.Instructions[p.PC].Register] += v
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
	data.Registers = make(map[string]int)
	// b inc 5 if a > 1
	regex := *regexp.MustCompile(`(\w+) (inc|dec) (-?\d+) if (\w+) (..?) (-?\d+)`)
	for i, l := range regex.FindAllStringSubmatch(strData, -1) {
		op, condition := Instruction{}, Instruction{}
		op.Condition = &condition
		data.Registers[l[1]] = 0
		op.Register = l[1]
		condition.Register = l[4]
		switch l[2] {
		case "inc":
			op.Op = Inc
		case "dec":
			op.Op = Dec
		}
		switch l[5] {
		case "<":
			condition.Op = LT
		case "<=":
			condition.Op = LE
		case ">":
			condition.Op = GT
		case ">=":
			condition.Op = GE
		case "==":
			condition.Op = EQ
		case "!=":
			condition.Op = NE
		}
		v, _ := strconv.Atoi(l[3])
		op.Value = v
		v, _ = strconv.Atoi(l[6])
		condition.Value = v
		data.Instructions = append(data.Instructions, op)
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
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
	problem.Process()
	for _, v := range problem.Registers {
		if v > result {
			result = v
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
