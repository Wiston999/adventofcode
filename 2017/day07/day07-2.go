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

type Program struct {
	Weight, DeepWeight int
	Name               string
	Parent             *Program
	Children           []*Program
}

func (p *Program) RecursiveWeight() (w int) {
	w = p.Weight
	for _, c := range p.Children {
		w += c.RecursiveWeight()
	}
	log.Debug(fmt.Sprintf("Weight of %s [%p]: %d (%d)", p.Name, p, p.Weight, w))
	p.DeepWeight = w
	return
}

type Problem struct {
	Programs map[string]*Program
}

func (p *Problem) Parent() *Program {
	for n := range p.Programs {
		if p.Programs[n].Parent == nil {
			return p.Programs[n]
		}
	}
	return nil
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
	// navfz (187) -> jviwcde, wfwor, vpfabxa
	regex := *regexp.MustCompile(`(\w+) \((\d+)\)(?: -> ([\w, ]+))?`)
	for i, l := range regex.FindAllStringSubmatch(strData, -1) {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		p, found := data.Programs[l[1]]
		if !found {
			p = new(Program)
		}
		w, _ := strconv.Atoi(l[2])
		p.Name = l[1]
		p.Weight = w
		if len(l[3]) > 0 {
			for _, top := range strings.Split(l[3], ", ") {
				pTop, found := data.Programs[top]
				if !found {
					pTop = new(Program)
				}
				pTop.Name = top
				pTop.Parent = p
				data.Programs[top] = pTop
				p.Children = append(p.Children, pTop)
			}
		}
		data.Programs[l[1]] = p
	}
	return
}

func (p *Problem) FindUnbalance(program Program) (*Program, int) {
	weights := make(map[int]int)
	for _, c := range program.Children {
		weights[c.DeepWeight]++
	}

	balance, unbalance := 0, 0
	//Different weight will have 1 as count
	for w, c := range weights {
		if c == 1 {
			unbalance = w
		} else {
			balance = w
		}
	}
	var child *Program
	for _, c := range program.Children {
		if c.DeepWeight == unbalance {
			child = c
		}
	}
	return child, balance - unbalance
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	parent := problem.Parent()
	log.Debug(fmt.Sprintf("Parent %#v total weight %d", parent, parent.RecursiveWeight()))

	current := parent
	unbalance := 0
	for {
		unbalanced, tmpUnbalance := problem.FindUnbalance(*current)
		if unbalanced == nil {
			break
		}
		unbalance = tmpUnbalance
		current = unbalanced
	}
	log.Debug(fmt.Sprintf("Found unbalanced %#v by %d", current, unbalance))
	result = current.Weight + unbalance
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
