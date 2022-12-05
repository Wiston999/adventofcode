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

type Stack struct {
	Data []byte
}

func (s *Stack) PushBack(data []byte) {
	for i := len(data) - 1; i >= 0; i-- {
		s.Data = append(s.Data, data[i])
	}
}

func (s *Stack) Push(data []byte) {
	for _, c := range data {
		s.Data = append(s.Data, c)
	}
}

func (s *Stack) Pop(count int) []byte {
	if count > len(s.Data) {
		count = len(s.Data)
	}
	last := len(s.Data)
	elements := s.Data[last-count:]
	s.Data = s.Data[:last-count]
	return elements
}

func (s *Stack) Reverse() {
	for i, j := 0, len(s.Data)-1; i < j; i, j = i+1, j-1 {
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
}

type Move struct {
	From, To, Count int
}

type Problem struct {
	Moves  []Move
	Stacks map[int]*Stack
	P2     bool
}

func (p *Problem) Apply(i int) {
	m := p.Moves[i]
	d := p.Stacks[m.From].Pop(m.Count)
	if p.P2 {
		p.Stacks[m.To].Push(d)
	} else {
		p.Stacks[m.To].PushBack(d)
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
	data.Stacks = make(map[int]*Stack)
	for i, l := range strings.Split(strData, "\n") {
		if len(l) == 0 {
			break
		}
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			if c >= 'A' && c <= 'Z' {
				index := ((j - 1) / 4) + 1
				if _, ok := data.Stacks[index]; !ok {
					data.Stacks[index] = &Stack{make([]byte, 0)}
				}
				log.Debug(fmt.Sprintf("Pushing %c (%d) to stack %d", c, c, index))
				data.Stacks[index].Push([]byte{byte(c)})
			}
		}
	}

	for _, s := range data.Stacks {
		s.Reverse()
	}
	regex := *regexp.MustCompile(`(?m)^move (\d+) from (\d+) to (\d+)$`)
	res := regex.FindAllStringSubmatch(string(byteData), -1)
	for _, l := range res {
		m := Move{
			atoi(l[2]),
			atoi(l[3]),
			atoi(l[1]),
		}
		data.Moves = append(data.Moves, m)
	}

	return
}

func solution(context *cli.Context) (result string) {
	var input = context.String("input")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	log.Debug("Parsed stacks")
	for k, s := range problem.Stacks {
		log.Debug(fmt.Sprintf("%02d: %v", k, s))
	}

	problem.P2 = context.Bool("part2")
	for i := range problem.Moves {
		log.Debug(fmt.Sprintf("Applying movement %03d", i))
		problem.Apply(i)
		for k, s := range problem.Stacks {
			log.Debug(fmt.Sprintf("%02d: %v", k, s))
		}
	}

	for i := 1; i <= len(problem.Stacks); i++ {
		last := len(problem.Stacks[i].Data)
		result += string(problem.Stacks[i].Data[last-1])
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
			&cli.BoolFlag{
				Name:    "part2",
				Aliases: []string{"p"},
				Value:   false,
				Usage:   "Solve part 2",
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
