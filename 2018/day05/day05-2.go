package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Problem struct {
	Chain []rune
}

func (p *Problem) Reduce() {
	var currentLength int
	offset := 'a' - 'A'
	for currentLength != len(p.Chain) {
		currentLength = len(p.Chain)
		for i := 0; i < len(p.Chain)-1; i++ {
			if p.Chain[i] == p.Chain[i+1]+offset || p.Chain[i] == p.Chain[i+1]-offset {
				p.Chain = append(p.Chain[:i], p.Chain[i+2:]...)
			}
		}
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
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for _, c := range l {
			data.Chain = append(data.Chain, c)
		}
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
	problem.Reduce()
	polymers := make(map[rune]bool)

	for _, c := range problem.Chain {
		if c >= 'a' && c <= 'z' {
			polymers[c] = true
		}
	}
	result = 9999999999
	for p := range polymers {
		tmp := Problem{}
		tmp.Chain = append(tmp.Chain, problem.Chain...)
		for i := 0; i < len(tmp.Chain); {
			if tmp.Chain[i] == p || tmp.Chain[i] == p-'a'+'A' {
				tmp.Chain = append(tmp.Chain[:i], tmp.Chain[i+1:]...)
			} else {
				i++
			}
		}
		tmp.Reduce()
		log.Info(fmt.Sprintf("Length after removing %c: %d [%d]", p, len(tmp.Chain), result))
		log.Debug(string(tmp.Chain))
		if len(tmp.Chain) < result {
			result = len(tmp.Chain)
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
