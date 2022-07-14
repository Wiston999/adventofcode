package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type MoveInstruction struct {
	direction string
	count     int
}

type Position struct {
	x, y int
	face string
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

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
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

func parseInput(input string) (data []MoveInstruction, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	for _, s := range bytes.Split(bytes.TrimSpace(tmpData), []byte(", ")) {
		var instruction MoveInstruction
		instruction.direction = string(s[0])
		instruction.count, err = strconv.Atoi(string(s[1:]))
		if err != nil {
			log.Warn(fmt.Sprintf("Error while parsing instruction %s", s))
		}
		data = append(data, instruction)
		log.Debug(fmt.Sprintf("Read instruction %#v", instruction))
	}
	return
}

func getFace(current, direction string) (face string) {
	face = current
	switch current {
	case "N":
		if direction == "R" {
			face = "E"
		} else if direction == "L" {
			face = "W"
		}
	case "E":
		if direction == "R" {
			face = "S"
		} else if direction == "L" {
			face = "N"
		}
	case "S":
		if direction == "R" {
			face = "W"
		} else if direction == "L" {
			face = "E"
		}
	case "W":
		if direction == "R" {
			face = "N"
		} else if direction == "L" {
			face = "S"
		}
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

	myPosition := Position{x: 0, y: 0, face: "N"}
	for i, d := range data {
		myPosition.face = getFace(myPosition.face, d.direction)
		switch myPosition.face {
		case "N":
			myPosition.y = myPosition.y + d.count
		case "S":
			myPosition.y = myPosition.y - d.count
		case "E":
			myPosition.x = myPosition.x + d.count
		case "W":
			myPosition.x = myPosition.x - d.count
		}
		log.Debug(fmt.Sprintf("New position after %d applied %#v movement %#v", i, d, myPosition))
	}

	result = Abs(myPosition.x) + Abs(myPosition.y)

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
