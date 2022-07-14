package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Position struct {
	x, y int
}

type Instruction struct {
	direction string
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

func parseInput(input string) (data [][]Instruction, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	for _, l := range bytes.Split(bytes.TrimSpace(tmpData), []byte("\n")) {
		var instructions []Instruction
		for _, s := range bytes.Split(bytes.TrimSpace(l), []byte("")) {
			instructions = append(instructions, Instruction{string(s)})
		}
		data = append(data, instructions)
	}
	return
}

func positionToNumber(position Position) (result string) {
	if position.x == 0 && position.y == 0 {
		return "1"
	} else if position.x == 1 && position.y == 0 {
		return "2"
	} else if position.x == 2 && position.y == 0 {
		return "3"
	} else if position.x == 0 && position.y == 1 {
		return "4"
	} else if position.x == 1 && position.y == 1 {
		return "5"
	} else if position.x == 2 && position.y == 1 {
		return "6"
	} else if position.x == 0 && position.y == 2 {
		return "7"
	} else if position.x == 1 && position.y == 2 {
		return "8"
	} else if position.x == 2 && position.y == 2 {
		return "9"
	}
	return ""
}

func solution(context *cli.Context) (result string) {
	var input = context.String("input")
	data, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	myPosition := Position{1, 1}
	for i, l := range data {
		for j, step := range l {
			switch step.direction {
			case "U":
				myPosition.y--
			case "D":
				myPosition.y++
			case "R":
				myPosition.x++
			case "L":
				myPosition.x--
			}
			if myPosition.y < 0 {
				myPosition.y = 0
			}
			if myPosition.y > 2 {
				myPosition.y = 2
			}
			if myPosition.x < 0 {
				myPosition.x = 0
			}
			if myPosition.x > 2 {
				myPosition.x = 2
			}
			log.Debug(fmt.Sprintf("New position after (%d, %d) applied %#v movement %#v", i, j, step, myPosition))
		}
		num := positionToNumber(myPosition)
		result += num
		log.Debug(fmt.Sprintf("Finished line %d at number %s. Current solution: %s", i, num, result))
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
