package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Guard struct {
	ID     string
	Sleeps [][]int
}

type Problem struct {
	Guards map[string]*Guard
}

func (p *Problem) Sleepiest() (id string, minute int) {
	var max int
	for i, g := range p.Guards {
		var slept int
		for _, s := range g.Sleeps {
			slept += len(s)
		}
		if slept > max {
			max = slept
			id = i
		}
	}
	minutes := make(map[int]int)
	for _, s := range p.Guards[id].Sleeps {
		for _, m := range s {
			minutes[m]++
		}
	}
	for m, v := range minutes {
		if v > minutes[minute] {
			minute = m
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

func atoi(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
}

func parseInput(input string) (data Problem, err error) {
	byteData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	strData := string(byteData)
	// Assume input is sorted
	data.Guards = make(map[string]*Guard)
	var current *Guard
	var currentStart int
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		parts := strings.Split(l, "] ")
		ts := parts[0]
		info := parts[1]
		if strings.HasPrefix(info, "Guard #") {
			id := strings.Split(info, " ")[1]
			if _, ok := data.Guards[id]; !ok {
				data.Guards[id] = new(Guard)
			}
			current = data.Guards[id]
			current.ID = id
		}
		if info == "falls asleep" {
			currentStart = atoi(strings.Split(ts, ":")[1])
		}
		if info == "wakes up" {
			end := atoi(strings.Split(ts, ":")[1])
			var sleep []int
			for i := currentStart; i < end; i++ {
				sleep = append(sleep, i)
			}
			current.Sleeps = append(current.Sleeps, sleep)
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
	id, minute := problem.Sleepiest()
	result = minute * atoi(id[1:])

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