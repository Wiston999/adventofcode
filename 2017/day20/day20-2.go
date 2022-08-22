package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Coord struct {
	X, Y, Z int
}

type Particle struct {
	Position, Speed, Acceleration Coord
}

func (p *Particle) Tick() {
	p.Speed.X += p.Acceleration.X
	p.Speed.Y += p.Acceleration.Y
	p.Speed.Z += p.Acceleration.Z

	p.Position.X += p.Speed.X
	p.Position.Y += p.Speed.Y
	p.Position.Z += p.Speed.Z
}

func (p *Particle) Distance() (d float64) {
	d = math.Abs(float64(p.Position.X)) + math.Abs(float64(p.Position.Y)) + math.Abs(float64(p.Position.Z))
	return
}

type Problem struct {
	Particles map[int]*Particle
}

func (p *Problem) Tick() {
	for _, particle := range p.Particles {
		particle.Tick()
	}

	collided := make([]int, 0)
	for i, p0 := range p.Particles {
		for j, p1 := range p.Particles {
			if i != j && p0.Position == p1.Position {
				collided = append(collided, i)
				collided = append(collided, j)
			}
		}
	}
	for _, i := range collided {
		log.Info(fmt.Sprintf("Removing particle that collided: %v", i))
		delete(p.Particles, i)
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
	// p=<-1724,-1700,5620>, v=<44,-10,-107>, a=<2,6,-9>
	regex := *regexp.MustCompile(`p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>`)
	data.Particles = make(map[int]*Particle)
	for i, l := range regex.FindAllStringSubmatch(strData, -1) {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		p := Particle{
			Position:     Coord{atoi(l[1]), atoi(l[2]), atoi(l[3])},
			Speed:        Coord{atoi(l[4]), atoi(l[5]), atoi(l[6])},
			Acceleration: Coord{atoi(l[7]), atoi(l[8]), atoi(l[9])},
		}
		data.Particles[i] = &p
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var limit = context.Int("limit")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	log.Info(fmt.Sprintf("Parsed %d particles", len(problem.Particles)))

	for i := 0; i < limit; i++ {
		problem.Tick()
		log.Debug(fmt.Sprintf("[%04d] %03d particles left", i, len(problem.Particles)))
	}

	result = len(problem.Particles)
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
				Name:  "limit",
				Value: 1000,
				Usage: "Number of iterations to be tested",
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
