package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"

	deque "github.com/gammazero/deque"
)

type Elf struct {
	Position int
	Next     *Elf
}

type Table struct {
	Elfs []Elf
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

	var left, right deque.Deque[int]

	for i := 1; i <= input; i++ {
		if i < (input/2)+1 {
			left.PushBack(i)
		} else {
			right.PushFront(i)
		}
	}

	for left.Len() > 0 && right.Len() > 0 {
		if left.Len() > right.Len() {
			left.PopBack()
		} else {
			right.PopBack()
		}

		right.PushFront(left.PopFront())
		left.PushBack(right.PopBack())
	}

	if left.Len() > 0 {
		result = left.PopFront()
	} else {
		result = right.PopFront()
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
			&cli.IntFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   3004953,
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
