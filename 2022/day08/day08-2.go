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
	Score      map[Coord]int
	MaxX, MaxY int
}

func (p *Problem) ComputeScores() {
	for c, h := range p.Map {
		if c.X == p.MaxX || c.Y == p.MaxY || c.X == 0 || c.Y == 0 {
			continue
		}
		tmpScore := 1
		score := 1
		for i := c.X + 1; i < p.MaxX && p.Map[Coord{i, c.Y}] < h; i++ {
			tmpScore++
		}
		log.Debug(fmt.Sprintf("TMP Score {%02d, %02d} = %02d", c.X, c.Y, tmpScore))
		score *= tmpScore
		tmpScore = 1
		for i := c.X - 1; i > 0 && p.Map[Coord{i, c.Y}] < h; i-- {
			tmpScore++
		}
		log.Debug(fmt.Sprintf("TMP Score {%02d, %02d} = %02d", c.X, c.Y, tmpScore))
		score *= tmpScore
		tmpScore = 1
		for j := c.Y + 1; j < p.MaxY && p.Map[Coord{c.X, j}] < h; j++ {
			tmpScore++
		}
		log.Debug(fmt.Sprintf("TMP Score {%02d, %02d} = %02d", c.X, c.Y, tmpScore))
		score *= tmpScore
		tmpScore = 1
		for j := c.Y - 1; j > 0 && p.Map[Coord{c.X, j}] < h; j-- {
			tmpScore++
		}
		score *= tmpScore
		log.Debug(fmt.Sprintf("TMP Score {%02d, %02d} = %02d", c.X, c.Y, tmpScore))
		log.Debug(fmt.Sprintf("Score {%02d, %02d} = %02d", c.X, c.Y, score))
		p.Score[c] = score
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
	data.Map = make(map[Coord]int)
	data.Score = make(map[Coord]int)
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
	problem.ComputeScores()
	for _, v := range problem.Score {
		if v > result {
			result = v
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
			echo(fmt.Sprintf("Solution is %v\n", solution(c)), c.String("output"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
