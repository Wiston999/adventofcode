package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Action struct {
	V, M int
	N    rune
}

type State struct {
	V0, V1 Action
}

type Problem struct {
	Position int
	Current  State
	States   map[rune]State
	Tape     map[int]int
}

func (p *Problem) Step() {
	v := p.Tape[p.Position]
	a := p.Current.V0
	if v == 1 {
		a = p.Current.V1
	}

	p.Tape[p.Position] = a.V
	p.Position += a.M
	p.Current = p.States[a.N]
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

func solution(context *cli.Context) (result int) {
	var input = context.Int("input")

	problem := Problem{}
	problem.States = make(map[rune]State)
	problem.States['A'] = State{
		Action{1, 1, 'B'},
		Action{0, -1, 'E'},
	}
	problem.States['B'] = State{
		Action{1, -1, 'C'},
		Action{0, 1, 'A'},
	}
	problem.States['C'] = State{
		Action{1, -1, 'D'},
		Action{0, 1, 'C'},
	}
	problem.States['D'] = State{
		Action{1, -1, 'E'},
		Action{0, -1, 'F'},
	}
	problem.States['E'] = State{
		Action{1, -1, 'A'},
		Action{1, -1, 'C'},
	}
	problem.States['F'] = State{
		Action{1, -1, 'E'},
		Action{1, 1, 'A'},
	}
	problem.Current = problem.States['A']
	problem.Tape = make(map[int]int)

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	for i := 0; i < input; i++ {
		problem.Step()
	}
	for _, v := range problem.Tape {
		result += v
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
			&cli.IntFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   12386363,
				Usage:   "Input value",
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
