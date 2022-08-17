package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type void struct{}

type Program struct {
	Name       string
	Neighbours []*Program
}

type Problem struct {
	Programs map[string]*Program
	Visited  map[string]void
}

func (p *Problem) Traverse(prog *Program) {
	p.Visited[prog.Name] = void{}
	for _, n := range prog.Neighbours {
		if _, ok := p.Visited[n.Name]; !ok {
			p.Traverse(n)
		}
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

func parseInput(input string) (data Problem, err error) {
	byteData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	strData := string(byteData)
	data.Programs = make(map[string]*Program)
	data.Visited = make(map[string]void)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		parts := strings.Split(l, " <-> ")
		pipes := strings.Split(parts[1], ", ")
		var neighbours []*Program
		for j, pname := range append(pipes, parts[0]) {
			p, ok := data.Programs[pname]
			if !ok {
				p = new(Program)
				p.Name = pname
			}
			data.Programs[pname] = p
			if j < len(pipes) {
				neighbours = append(neighbours, p)
			}
		}

		data.Programs[parts[0]].Neighbours = neighbours
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
	ungrouped := make(map[string]void)
	for _, p := range problem.Programs {
		log.Debug(fmt.Sprintf("Program %s has %d neighbours: %v", p.Name, len(p.Neighbours), p.Neighbours))
		ungrouped[p.Name] = void{}
	}

	for len(ungrouped) > 0 {
		start := ""
		for k := range ungrouped {
			start = k
			break
		}
		problem.Visited = make(map[string]void)
		problem.Traverse(problem.Programs[start])
		log.Debug(fmt.Sprintf("Traversed from program 0 %v", problem.Visited))
		for p := range problem.Visited {
			delete(ungrouped, p)
		}
		result++
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
