package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Coord struct {
	i, j int
}

type Problem struct {
	Numbers map[Coord]int
}

func (p *Problem) Next(c Coord) (n Coord) {
	if c.i == 0 && c.j == 0 { // Start point
		n.i = 1
	} else {
		_, okT := p.Numbers[Coord{c.i, c.j - 1}]
		_, okB := p.Numbers[Coord{c.i, c.j + 1}]
		_, okL := p.Numbers[Coord{c.i - 1, c.j}]
		_, okR := p.Numbers[Coord{c.i + 1, c.j}]

		log.Debug(fmt.Sprintf("Testing %v, TBLR [%v %v %v %v]", c, okT, okB, okL, okR))
		if !okT && !okB && okL && !okR { // New ring, move up
			n.i = c.i
			n.j = c.j - 1
		} else if !okB && okL && !okR { // Move right
			n.i = c.i + 1
			n.j = c.j
		} else if !okT && !okB && okL && !okR { // Bottom - right corner, move up
			n.i = c.i
			n.j = c.j - 1
		} else if !okT && okB && okL { // Move up
			n.i = c.i
			n.j = c.j - 1
		} else if !okT && okB && !okL && !okR { // Top - right corner, move right
			n.i = c.i - 1
			n.j = c.j
		} else if !okT && okB && !okL && okR { // Move right
			n.i = c.i - 1
			n.j = c.j
		} else if !okT && !okB && !okL && okR { // Top - left corner, move down
			n.i = c.i
			n.j = c.j + 1
		} else if okT && !okB && !okL && okR { // Move down
			n.i = c.i
			n.j = c.j + 1
		} else if okT && !okB && !okL && !okR { // Bottom - left corner, move right
			n.i = c.i + 1
			n.j = c.j
		}
	}
	p.Numbers[n] = p.Numbers[c] + 1
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

func solution(context *cli.Context) (result int) {
	var input = context.Int("input")

	log.Debug(fmt.Sprintf("Parsed input %#v", input))

	problem := Problem{make(map[Coord]int)}
	current := Coord{0, 0}
	problem.Numbers[current] = 1
	for i := 0; i < input-1; i++ {
		current = problem.Next(current)
		log.Debug(fmt.Sprintf("[%06d] Generated %d at %v", i, problem.Numbers[current], current))
	}
	result = int(math.Abs(float64(current.i)) + math.Abs(float64(current.j)))
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
				Value:   361527,
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
