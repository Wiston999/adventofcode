package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Node struct {
	Name     string
	Size     int
	Children map[string]*Node
	Parent   *Node
}

func (n *Node) Usage() (result int) {
	if n.Size > 0 {
		result = n.Size
	} else {
		for _, v := range n.Children {
			result += v.Usage()
		}
	}
	return
}

func (n *Node) Print(depth int) (result string) {
	for i := 0; i < depth*2; i++ {
		result += "-"
	}
	if n.Size > 0 {
		result += fmt.Sprintf("%s [size=%05d]\n", n.Name, n.Size)
	} else {
		result += fmt.Sprintf("%s [directory]\n", n.Name)
		for _, v := range n.Children {
			result += v.Print(depth + 1)
		}
	}
	return
}

func (n *Node) FindSmaller(limit int) (result int) {
	if n.Size == 0 { // Is directory
		usage := n.Usage()
		if usage < limit {
			result += usage
		}
		for _, v := range n.Children {
			result += v.FindSmaller(limit)
		}
	}

	return
}

type Problem struct {
	Tree Node
}

func (p *Problem) Print() string {
	return p.Tree.Print(0)
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
	data.Tree.Name = "/"
	data.Tree.Children = make(map[string]*Node)
	current := &data.Tree
	strData := string(byteData)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		parts := strings.Split(l, " ")
		if strings.HasPrefix(l, "$ cd") {
			switch parts[2] {
			case "/":
				current = &data.Tree
			case "..":
				current = current.Parent
			default:
				current = current.Children[parts[2]]
			}
		} else if strings.HasPrefix(l, "$ ls") {
			continue
		} else {
			n := new(Node)
			n.Children = make(map[string]*Node)
			n.Name = parts[1]
			n.Parent = current
			if i, err := strconv.Atoi(parts[0]); err == nil {
				n.Size = i
			}
			current.Children[parts[1]] = n
		}
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	maxSize := 100000
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	fmt.Println(problem.Print())
	result = problem.Tree.FindSmaller(maxSize)

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
