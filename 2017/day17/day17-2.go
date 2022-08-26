package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Problem struct {
	Pos0, Pos1            int
	Size, Current, Offset int
}

func NewProblem(size, offset int) (p *Problem) {
	p = new(Problem)
	p.Pos1 = 0
	p.Size = 0
	p.Offset = offset
	return
}

func (p *Problem) Insert(v int) {
	if p.Size == 0 {
		p.Current = 0
		p.Size = 1
		return
	}
	p.Current = (p.Current + p.Offset) % p.Size
	p.Current++
	// 0 doesn't move, just track number that is at position 1
	if p.Current == 1 {
		p.Pos1 = v
	}
	p.Size++
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
	var maxSize = context.Int("max")

	problem := NewProblem(maxSize, input)
	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))

	progress := maxSize / 20
	for i := 0; i < maxSize; i++ {
		problem.Insert(i)
		if i%progress == 0 {
			log.Info(fmt.Sprintf("Processed %d out of %d (%02.2f %%)", i, maxSize, float64(i*100.0/(maxSize*1.0))))
		}
	}

	result = problem.Pos1
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
				Name:    "max",
				Aliases: []string{"m"},
				Value:   50000000,
				Usage:   "Max iterations",
			},
			&cli.IntFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   348,
				Usage:   "Input value",
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