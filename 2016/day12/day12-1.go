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
	Copy = iota
	Inc
	Dec
	Jnz
)

type InstructionType int

type Operand struct {
	Register string
	Value    int
}

type Instruction struct {
	Type     InstructionType
	Operand1 Operand
	Operand2 Operand
}

type Computer struct {
	Registers    map[string]int
	Instructions []Instruction
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

func parseInput(input string) (data Computer, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	for i, l := range strings.Split(strings.TrimSpace(string(tmpData)), "\n") {
		parts := strings.Split(l, " ")
		op := Instruction{}
		switch parts[0] {
		case "cpy":
			op.Type = Copy
			op.Operand2.Register = parts[2]
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				op.Operand1.Register = parts[1]
			} else {
				op.Operand1.Value = v
			}
		case "inc":
			op.Type = Inc
			op.Operand1.Register = parts[1]
		case "dec":
			op.Type = Dec
			op.Operand1.Register = parts[1]
		case "jnz":
			op.Type = Jnz
			v, _ := strconv.Atoi(parts[2])
			op.Operand2.Value = v
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				op.Operand1.Register = parts[1]
			} else {
				op.Operand1.Value = v
			}
		}
		log.Debug(fmt.Sprintf("Parsed instruction (%d): %#v", i, op))
		data.Instructions = append(data.Instructions, op)
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

	data.Registers = make(map[string]int)
	data.Registers["a"] = 0
	data.Registers["b"] = 0
	data.Registers["c"] = 0
	data.Registers["d"] = 0
	for pc := 0; pc < len(data.Instructions); {
		op := data.Instructions[pc]
		log.Debug(fmt.Sprintf("Applying operation at PC[%d] (%#v) to %#v", pc, op, data.Registers))
		switch op.Type {
		case Copy:
			if op.Operand1.Register != "" {
				data.Registers[op.Operand2.Register] = data.Registers[op.Operand1.Register]
			} else {
				data.Registers[op.Operand2.Register] = op.Operand1.Value
			}
		case Inc:
			data.Registers[op.Operand1.Register]++
		case Dec:
			data.Registers[op.Operand1.Register]--
		case Jnz:
			v := data.Registers[op.Operand1.Register]
			if op.Operand1.Register == "" {
				v = op.Operand1.Value
			}
			if v != 0 {
				pc += op.Operand2.Value
			} else {
				pc++
			}
		}
		if op.Type != Jnz {
			pc++
		}
	}

	result = data.Registers["a"]
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
