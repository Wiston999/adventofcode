package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	SND = "sound"
	SET = "set"
	ADD = "add"
	MUL = "mul"
	MOD = "mod"
	RCV = "recover"
	JGZ = "jgz"
)

type OpType string

type Operand struct {
	I int
	A string
}

func (o *Operand) Value(regs map[string]int) int {
	if o.A != "" {
		return regs[o.A]
	}
	return o.I
}

type Instruction struct {
	Type OpType
	A, B Operand
}

type Problem struct {
	PC                    int
	Instructions          []Instruction
	Registers             map[string]int
	LastSend, LastRecover int
}

func (p *Problem) Apply() bool {
	op := p.Instructions[p.PC]

	switch op.Type {
	case SND:
		p.LastSend = op.A.Value(p.Registers)
	case SET:
		p.Registers[op.A.A] = op.B.Value(p.Registers)
	case ADD:
		p.Registers[op.A.A] += op.B.Value(p.Registers)
	case MUL:
		p.Registers[op.A.A] *= op.B.Value(p.Registers)
	case MOD:
		p.Registers[op.A.A] %= op.B.Value(p.Registers)
	case RCV:
		if op.A.Value(p.Registers) != 0 {
			p.LastRecover = p.LastSend
		}
	case JGZ:
		if op.A.Value(p.Registers) > 0 {
			p.PC += op.B.Value(p.Registers)
		} else {
			p.PC++
		}
	}

	if op.Type != JGZ {
		p.PC++
	}
	return p.PC < len(p.Instructions) && p.LastRecover == 0 // Finished
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
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		parts := strings.Split(l, " ")
		op := Instruction{}
		switch parts[0] {
		case "snd":
			op.Type = SND
		case "set":
			op.Type = SET
		case "add":
			op.Type = ADD
		case "mul":
			op.Type = MUL
		case "mod":
			op.Type = MOD
		case "rcv":
			op.Type = RCV
		case "jgz":
			op.Type = JGZ
		}
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			op.A = Operand{A: parts[1]}
			data.Registers[parts[1]] = 0
		} else {
			op.A = Operand{I: v}
		}
		if len(parts) > 2 {
			v, err = strconv.Atoi(parts[2])
			if err != nil {
				data.Registers[parts[2]] = 0
				op.B = Operand{A: parts[2]}
			} else {
				op.B = Operand{I: v}
			}
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
	i := 0
	for problem.Apply() {
		i++
		log.Debug(fmt.Sprintf("[%03d] Processing instruction PC [%d] : %v", i, problem.PC, problem.Registers))
	}
	result = problem.LastRecover

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
