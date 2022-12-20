package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Problem struct {
	Original, Positions []int
	Zero                int
	LList               *Node
}

type Node struct {
	Value      int
	Next, Prev *Node
}

func (n *Node) Print() string {
	l := make([]string, 1)
	l[0] = strconv.Itoa(n.Value)
	current := n.Next
	for current != n {
		l = append(l, strconv.Itoa(current.Value))
		current = current.Next
	}

	return strings.Join(l, ", ")
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (p *Problem) MixLList() {
	l := len(p.Original)
	for i, n := range p.Original {
		currentI := -1
		current := p.LList
		for j := 0; ; j++ {
			if current.Value == i {
				currentI = j
				break
			}
			if current.Next == p.LList {
				break
			}

			current = current.Next
		}

		nCurrent := p.LList
		log.Debug(fmt.Sprintf("[%03d] %d (%v) Replacing %d by %d", i, n, p.LList.Print(), currentI, mod(currentI+n, l-1)))
		if currentI == mod(currentI+n, l-1) {
			continue
		}

		current.Prev.Next = current.Next
		current.Next.Prev = current.Prev
		nCurrent = current
		for j := 0; j < mod(n, l-1); j++ {
			nCurrent = nCurrent.Next
		}
		nCurrent.Next.Prev = current
		current.Next = nCurrent.Next
		current.Prev = nCurrent
		nCurrent.Next = current
		log.Debug(fmt.Sprintf("[%03d] %s", i, p.LList.Print()))
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

func atoiSafe(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
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
	data.Original = make([]int, 0)
	prev := new(Node)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		data.Original = append(data.Original, atoiSafe(l))
		n := Node{Value: i}
		if i == 0 {
			data.LList = &n
		} else {
			prev.Next = &n
			n.Prev = prev
		}
		if atoiSafe(l) == 0 {
			data.Zero = i
		}
		prev = &n
	}
	prev.Next = data.LList
	data.LList.Prev = prev
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
	log.Debug(fmt.Sprintf("Linked List: %s", problem.LList.Print()))
	problem.MixLList()
	current := problem.LList
	for problem.Original[current.Value] != 0 {
		current = current.Next
	}
	log.Info(fmt.Sprintf("0 found at %d (%d)", current.Value, problem.Original[current.Value]))
	for i := 0; i <= 3000; i++ {
		if i == 1000 || i == 2000 || i == 3000 {
			log.Info(fmt.Sprintf("[%04d] %d %d", i, current.Value, problem.Original[current.Value]))
			result += problem.Original[current.Value]
		}
		current = current.Next
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
			start := time.Now()
			echo(fmt.Sprintf("Solution is %v in %s", solution(c), time.Since(start)), c.String("output"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
