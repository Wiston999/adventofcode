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
	SPIN     = "spin"
	EXCHANGE = "exchange"
	PARTNER  = "partner"
)

type OpType string

type Operand struct {
	I int
	A rune
}

type Instruction struct {
	A, B Operand
	Type OpType
}

func (i *Instruction) Apply(input string) (output string) {
	switch i.Type {
	case SPIN:
		output = input[len(input)-i.A.I:] + input[:len(input)-i.A.I]
	case EXCHANGE:
		inputByte := []byte(input)
		inputByte[i.A.I], inputByte[i.B.I] = inputByte[i.B.I], inputByte[i.A.I]
		output = string(inputByte)
	case PARTNER:
		output = strings.ReplaceAll(input, string(i.A.A), "-")
		output = strings.ReplaceAll(output, string(i.B.A), string(i.A.A))
		output = strings.ReplaceAll(output, "-", string(i.B.A))
	}
	return
}

type Problem struct {
	Instructions []Instruction
	String       string
}

func (p *Problem) StartString(size int) {
	for i := 0; i < size; i++ {
		p.String += string('a' + i)
	}
}

func (p *Problem) Apply() {
	for i, op := range p.Instructions {
		p.String = op.Apply(p.String)
		log.Debug(fmt.Sprintf("[%03d] String is now: %s", i, p.String))
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
	for i, l := range strings.Split(strings.TrimSpace(strData), ",") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		op := Instruction{}
		parts := strings.Split(l[1:], "/")
		switch l[0] {
		case 's':
			op.Type = SPIN
			v, _ := strconv.Atoi(parts[0])
			op.A = Operand{I: v}
		case 'x':
			op.Type = EXCHANGE
			v, _ := strconv.Atoi(parts[0])
			op.A = Operand{I: v}
			v, _ = strconv.Atoi(parts[1])
			op.B = Operand{I: v}
		case 'p':
			op.Type = PARTNER
			op.A = Operand{A: rune(l[1])}
			op.B = Operand{A: rune(l[3])}
		}
		data.Instructions = append(data.Instructions, op)
	}
	return
}

func solution(context *cli.Context) (result string) {
	var input = context.String("input")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	problem.StartString(context.Int("size"))
	iterations := context.Int("iterations")
	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	progress := iterations / 100
	states := make(map[string]int)
	for i := 0; i < iterations; i++ {
		problem.Apply()
		if j, ok := states[problem.String]; ok {
			rest := (iterations - i) % (i - j)
			iterations = i + rest
			log.Info(fmt.Sprintf("[%06d] String %s already found at %06d iteration, fast-forwarding to %d", i, problem.String, j, rest))
		}
		states[problem.String] = i
		if i%progress == 0 {
			log.Info(fmt.Sprintf("Processed %d / %d (%02.2f %%)", i, iterations, float64(i*100.0/(iterations*1.0))))
		}
	}
	result = problem.String
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
				Name:  "iterations",
				Value: 1000000000,
				Usage: "Number of iterations",
			},
			&cli.IntFlag{
				Name:    "size",
				Aliases: []string{"s"},
				Value:   16,
				Usage:   "Number of characters",
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
