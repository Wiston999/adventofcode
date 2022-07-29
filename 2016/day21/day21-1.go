package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	LeftRotation = iota
	RightRotation
)

const (
	SwapPosition = iota
	SwapLetter
	RotateSteps
	RotateBased
	Reverse
	Move
)

type Instruction struct {
	Kind int
	X, Y int
}

func (i *Instruction) Apply(v []byte) (r []byte) {
	r = v
	switch i.Kind {
	case SwapPosition:
		r[i.X], r[i.Y] = r[i.Y], r[i.X]
	case SwapLetter:
		a, b := rune(i.X), rune(i.Y)
		r = bytes.ReplaceAll(r, []byte(string(a)), []byte("\n"))
		r = bytes.ReplaceAll(r, []byte(string(b)), []byte(string(a)))
		r = bytes.ReplaceAll(r, []byte("\n"), []byte(string(b)))
	case RotateSteps:
		for j := 0; j < i.X; j++ {
			if i.Y == LeftRotation {
				r = append(r[1:], r[0])
			} else {
				r = append([]byte{r[len(r)-1]}, r[:len(r)-1]...)
			}
		}
	case RotateBased:
		index := bytes.IndexRune(r, rune(i.X))
		if index >= 4 {
			index++
		}
		index++
		for j := 0; j < index; j++ {
			r = append([]byte{r[len(r)-1]}, r[:len(r)-1]...)
		}
	case Reverse:
		for j, k := i.X, i.Y; j < k; j, k = j+1, k-1 {
			r[j], r[k] = r[k], r[j]
		}
	case Move:
		c := r[i.X]
		r = append(v[:i.X], v[i.X+1:]...)
		r = []byte(fmt.Sprintf("%s%c%s", r[:i.Y], c, r[i.Y:]))
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

func atoi(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
}

func parseInput(input string) (data []Instruction, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	swapPosRegex := *regexp.MustCompile(`swap position (\d+) with position (\d+)`)
	swapLetRegex := *regexp.MustCompile(`swap letter (\w+) with letter (\w+)`)
	rotateStepRegex := *regexp.MustCompile(`rotate (left|right) (\d+) steps?`)
	rotatePosRegex := *regexp.MustCompile(`rotate based on position of letter (\w+)`)
	reverseRegex := *regexp.MustCompile(`reverse positions (\d+) through (\d+)`)
	moveRegex := *regexp.MustCompile(`move position (\d+) to position (\d+)`)
	for i, l := range strings.Split(strings.TrimSpace(string(tmpData)), "\n") {
		op := Instruction{}
		if res := swapPosRegex.FindAllStringSubmatch(l, -1); len(res) > 0 {
			op.Kind = SwapPosition
			op.X = atoi(res[0][1])
			op.Y = atoi(res[0][2])
		}
		if res := swapLetRegex.FindAllStringSubmatch(l, -1); len(res) > 0 {
			op.Kind = SwapLetter
			op.X = int(res[0][1][0])
			op.Y = int(res[0][2][0])
		}
		if res := rotateStepRegex.FindAllStringSubmatch(l, -1); len(res) > 0 {
			op.Kind = RotateSteps
			op.X = atoi(res[0][2])
			if res[0][1] == "left" {
				op.Y = LeftRotation
			} else {
				op.Y = RightRotation
			}
		}
		if res := rotatePosRegex.FindAllStringSubmatch(l, -1); len(res) > 0 {
			op.Kind = RotateBased
			op.X = int(res[0][1][0])
		}
		if res := reverseRegex.FindAllStringSubmatch(l, -1); len(res) > 0 {
			op.Kind = Reverse
			op.X = atoi(res[0][1])
			op.Y = atoi(res[0][2])
		}
		if res := moveRegex.FindAllStringSubmatch(l, -1); len(res) > 0 {
			op.Kind = Move
			op.X = atoi(res[0][1])
			op.Y = atoi(res[0][2])
		}
		log.Debug(fmt.Sprintf("Parsed operation [%03d] %#v", i+1, op))
		data = append(data, op)
	}
	return
}

func solution(context *cli.Context) (result string) {
	var input = context.String("input")
	data, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	value := []byte(context.String("value"))
	for i, op := range data {
		before := bytes.Map(func(r rune) rune { return r }, value)
		value = op.Apply(value)
		log.Debug(fmt.Sprintf("Value after op [%03d] %s --> %s", i+1, before, value))
	}
	result = string(value)
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
				Name:    "value",
				Aliases: []string{"v"},
				Value:   "abcdefgh",
				Usage:   "Scrambled value",
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
