package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Problem struct {
	Stream []string
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
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		data.Stream = append(data.Stream, l)
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

	for _, i := range problem.Stream {
		numbers := []rune{}
		for j, c := range i {
			if c >= '0' && c <= '9' {
				numbers = append(numbers, c-'0')
			} else {
				if (j+3) < len(i) && "zero" == fmt.Sprintf("%c%c%c%c", c, i[j+1], i[j+2], i[j+3]) {
					numbers = append(numbers, 0)
				} else if (j+2) < len(i) && "one" == fmt.Sprintf("%c%c%c", c, i[j+1], i[j+2]) {
					numbers = append(numbers, 1)
				} else if (j+2) < len(i) && "two" == fmt.Sprintf("%c%c%c", c, i[j+1], i[j+2]) {
					numbers = append(numbers, 2)
				} else if (j+4) < len(i) && "three" == fmt.Sprintf("%c%c%c%c%c", c, i[j+1], i[j+2], i[j+3], i[j+4]) {
					numbers = append(numbers, 3)
				} else if (j+3) < len(i) && "four" == fmt.Sprintf("%c%c%c%c", c, i[j+1], i[j+2], i[j+3]) {
					numbers = append(numbers, 4)
				} else if (j+3) < len(i) && "five" == fmt.Sprintf("%c%c%c%c", c, i[j+1], i[j+2], i[j+3]) {
					numbers = append(numbers, 5)
				} else if (j+2) < len(i) && "six" == fmt.Sprintf("%c%c%c", c, i[j+1], i[j+2]) {
					numbers = append(numbers, 6)
				} else if (j+4) < len(i) && "seven" == fmt.Sprintf("%c%c%c%c%c", c, i[j+1], i[j+2], i[j+3], i[j+4]) {
					numbers = append(numbers, 7)
				} else if (j+4) < len(i) && "eight" == fmt.Sprintf("%c%c%c%c%c", c, i[j+1], i[j+2], i[j+3], i[j+4]) {
					numbers = append(numbers, 8)
				} else if (j+3) < len(i) && "nine" == fmt.Sprintf("%c%c%c%c", c, i[j+1], i[j+2], i[j+3]) {
					numbers = append(numbers, 9)
				}
			}
		}
		log.Debug(fmt.Sprintf("Adding %v%v - %v", numbers[0], numbers[len(numbers)-1], numbers))
		n, _ := strconv.Atoi(fmt.Sprintf("%v%v", numbers[0], numbers[len(numbers)-1]))
		result += n
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
