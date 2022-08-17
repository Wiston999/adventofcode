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
	List    []int
	Lengths []int
	Size    int
}

func (p *Problem) MakeList() {
	p.List = make([]int, p.Size)
	for i := range p.List {
		p.List[i] = i
	}
}

func (p *Problem) Reverse(start, length int) {
	for i, j := start, start+length-1; i < j; i, j = i+1, j-1 {
		p.List[i%p.Size], p.List[j%p.Size] = p.List[j%p.Size], p.List[i%p.Size]
		log.Debug(fmt.Sprintf("%d <-> %d: %v", i, j, p.List))
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
	for i, l := range strings.Split(strings.TrimSpace(strData), ",") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		i, _ := strconv.Atoi(l)
		data.Lengths = append(data.Lengths, i)
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var size = context.Int("size")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	problem.Size = size
	problem.MakeList()
	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))

	current := 0
	skip := 0
	for i, n := range problem.Lengths {
		problem.Reverse(current, n)
		log.Debug(fmt.Sprintf("[%03d] List after reversal %d, skip %d at %d: %v", i, n, skip, current, problem.List))
		current += skip + n
		skip++
	}

	result = problem.List[0] * problem.List[1]
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
			&cli.IntFlag{
				Name:    "size",
				Aliases: []string{"s"},
				Value:   256,
				Usage:   "List size",
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
