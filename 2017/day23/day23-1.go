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
	SET = "set"
	SUB = "sub"
	MUL = "mul"
	JNZ = "jnz"
)

type OpType string

type Operand struct {
	A string
	I int
}

func (o *Operand) Value(regs map[string]int) (r int) {
	r = o.I
	if o.A != "" {
		r = regs[o.A]
	}
	return
}

type Operation struct {
	A, B Operand
	Type OpType
}

type Problem struct {
	PC           int
	Instructions []Operation
	Registers    map[string]int
	Counter      map[OpType]int
}

func (p *Problem) Step() bool {
	op := p.Instructions[p.PC]
	switch op.Type {
	case SET:
		p.Registers[op.A.A] = op.B.Value(p.Registers)
	case SUB:
		p.Registers[op.A.A] -= op.B.Value(p.Registers)
	case MUL:
		p.Registers[op.A.A] *= op.B.Value(p.Registers)
	case JNZ:
		if op.A.Value(p.Registers) != 0 {
			p.PC += op.B.Value(p.Registers)
		} else {
			p.PC++
		}
	}
	p.Counter[op.Type]++
	if op.Type != JNZ {
		p.PC++
	}
	return p.PC >= 0 && p.PC < len(p.Instructions)
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
	data.Counter = make(map[OpType]int)
	regex := *regexp.MustCompile(`(\w+) ([a-z]|(?:-?\d+)) ([a-z]|(?:-?\d+))`)
	for i, l := range regex.FindAllStringSubmatch(strData, -1) {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		op := Operation{}
		v, err := strconv.Atoi(l[2])
		if err != nil {
			op.A.A = l[2]
		} else {
			op.A.I = v
		}
		v, err = strconv.Atoi(l[3])
		if err != nil {
			op.B.A = l[3]
		} else {
			op.B.I = v
		}
		switch l[1] {
		case "set":
			op.Type = SET
		case "sub":
			op.Type = SUB
		case "mul":
			op.Type = MUL
		case "jnz":
			op.Type = JNZ
		}
		data.Instructions = append(data.Instructions, op)
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

	var i int
	for problem.Step() {
		log.Debug(fmt.Sprintf("[%04d] PC=[%03d] Registers=%v", i, problem.PC, problem.Registers))
	}
	result = problem.Counter["mul"]
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
