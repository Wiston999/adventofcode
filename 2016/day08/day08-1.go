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
	RectInstruction = iota
	RotColInstruction
	RotRowInstruction
)

type Instruction struct {
	Kind int
	A, B int
}

type Screen [][]bool

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

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func parseInput(input string) (data []Instruction, err error) {
	tmpDataBytes, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	tmpData := string(tmpDataBytes)
	rectRegex := *regexp.MustCompile(`^rect (\d+)x(\d+)$`)
	rrRegex := *regexp.MustCompile(`^rotate row y=(\d+) by (\d+)$`)
	rcRegex := *regexp.MustCompile(`^rotate column x=(\d+) by (\d+)$`)
	for _, l := range strings.Split(strings.TrimSpace(tmpData), "\n") {
		l = strings.TrimSpace(l)
		op := Instruction{}
		if rectMatch := rectRegex.FindStringSubmatch(l); len(rectMatch) > 0 {
			op.Kind = RectInstruction
			op.A = atoi(rectMatch[1])
			op.B = atoi(rectMatch[2])
		}
		if rrMatch := rrRegex.FindStringSubmatch(l); len(rrMatch) > 0 {
			op.Kind = RotRowInstruction
			op.A = atoi(rrMatch[1])
			op.B = atoi(rrMatch[2])
		}
		if rcMatch := rcRegex.FindStringSubmatch(l); len(rcMatch) > 0 {
			op.Kind = RotColInstruction
			op.A = atoi(rcMatch[1])
			op.B = atoi(rcMatch[2])
		}
		data = append(data, op)
	}
	return
}

func printScreen(screen Screen) (result string) {
	for i := range screen {
		for j := range screen[i] {
			if screen[i][j] {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return
}

func mkScreen(a, b int) [][]bool {
	screen := make([][]bool, a)
	for i := range screen {
		screen[i] = make([]bool, b)
	}
	return screen
}

func cpScreen(s Screen) [][]bool {
	screen := mkScreen(len(s), len(s[0]))
	for i := range s {
		copy(screen[i], s[i])
	}
	return screen
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	data, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	screen := mkScreen(6, 50)

	for i, operation := range data {
		if operation.Kind == RectInstruction {
			for j := 0; j < operation.A; j++ {
				for k := 0; k < operation.B; k++ {
					screen[k][j] = true
				}
			}
		}
		if operation.Kind == RotRowInstruction {
			for j := 0; j < operation.B; j++ {
				screenCopy := cpScreen(screen)
				size := len(screen[operation.A])
				for k := 0; k < size-1; k++ {
					screen[operation.A][(k+1)%size] = screenCopy[operation.A][k%size]
				}
				screen[operation.A][0] = screenCopy[operation.A][size-1]
			}
		}
		if operation.Kind == RotColInstruction {
			for j := 0; j < operation.B; j++ {
				screenCopy := cpScreen(screen)
				size := len(screen)
				for k := 0; k < size-1; k++ {
					screen[(k+1)%size][operation.A] = screenCopy[k%size][operation.A]
				}
				screen[0][operation.A] = screenCopy[size-1][operation.A]
			}
		}
		log.Debug(fmt.Sprintf("(%03d) Applied operation: %#v", i, operation))
		log.Debug(fmt.Sprintf("\n%s", printScreen(screen)))
	}

	for i := range screen {
		for j := range screen[i] {
			if screen[i][j] {
				result++
			}
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
