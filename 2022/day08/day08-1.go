package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Coord struct {
	X, Y int
}

type Problem struct {
	Map        map[Coord]int
	Seen       map[Coord]bool
	MaxX, MaxY int
}

func (p *Problem) CountVisible() (result int) {
	for i := 0; i <= p.MaxX; i++ {
		bigger := -1
		for j := 0; j <= p.MaxY; j++ {
			log.Debug(fmt.Sprintf("Testing {%02d, %02d} from top [%d > %d]", i, j, p.Map[Coord{i, j}], bigger))
			if p.Map[Coord{i, j}] > bigger {
				if _, ok := p.Seen[Coord{i, j}]; !ok {
					result++
					log.Debug(fmt.Sprintf("Seen {%02d, %02d} from left [%d > %d]", i, j, p.Map[Coord{i, j}], bigger))
					p.Seen[Coord{i, j}] = true
				}
				bigger = p.Map[Coord{i, j}]
			}
		}
		bigger = -1
		for j := p.MaxY; j >= 0; j-- {
			log.Debug(fmt.Sprintf("Testing {%02d, %02d} from top [%d > %d]", i, j, p.Map[Coord{i, j}], bigger))
			if p.Map[Coord{i, j}] > bigger {
				if _, ok := p.Seen[Coord{i, j}]; !ok {
					result++
					p.Seen[Coord{i, j}] = true
					log.Debug(fmt.Sprintf("Seen {%02d, %02d} from right [%d > %d]", i, j, p.Map[Coord{i, j}], bigger))
				}
				bigger = p.Map[Coord{i, j}]
			}
		}
	}
	for j := 0; j <= p.MaxY; j++ {
		bigger := -1
		for i := 0; i <= p.MaxX; i++ {
			log.Debug(fmt.Sprintf("Testing {%02d, %02d} from top [%d > %d]", i, j, p.Map[Coord{i, j}], bigger))
			if p.Map[Coord{i, j}] > bigger {
				if _, ok := p.Seen[Coord{i, j}]; !ok {
					result++
					log.Debug(fmt.Sprintf("Seen {%02d, %02d} from top [%d > %d]", i, j, p.Map[Coord{i, j}], bigger))
					p.Seen[Coord{i, j}] = true
				}
				bigger = p.Map[Coord{i, j}]
			}
		}
		bigger = -1
		for i := p.MaxX; i >= 0; i-- {
			log.Debug(fmt.Sprintf("Testing {%02d, %02d} from top [%d > %d]", i, j, p.Map[Coord{i, j}], bigger))
			if p.Map[Coord{i, j}] > bigger {
				if _, ok := p.Seen[Coord{i, j}]; !ok {
					result++
					p.Seen[Coord{i, j}] = true
					log.Debug(fmt.Sprintf("Seen {%02d, %02d} from bottom [%d > %d]", i, j, p.Map[Coord{i, j}], bigger))
				}
				bigger = p.Map[Coord{i, j}]
			}
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
	data.Map = make(map[Coord]int)
	data.Seen = make(map[Coord]bool)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			v, _ := strconv.Atoi(string(c))
			data.Map[Coord{i, j}] = v
			if j > data.MaxY {
				data.MaxY = j
			}
		}
		if i > data.MaxX {
			data.MaxX = i
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
	result = problem.CountVisible()

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
			echo(fmt.Sprintf("Solution is %v\n", solution(c)), c.String("output"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
