package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
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

	// Linked list solution
	first := Elf{Position: 1}
	current := &first
	for i := 1; i < input; i++ {
		newElf := Elf{Position: i + 1}
		current.Next = &newElf
		current = &newElf
	}
	current.Next = &first

	current = &first
	for current.Next != current {
		current.Next = current.Next.Next
		current = current.Next
	}
	fmt.Println(current.Position)
	//https://en.wikipedia.org/wiki/Josephus_problem#

	log.Debug(fmt.Sprintf("2*(%d - 2 ^( floor( log2(%d) ) ) ) + 1", input, input))
	log.Debug(fmt.Sprintf("2*(%d - 2 ^( floor( %v ) ) ) + 1", input, math.Log2(float64(input))))
	log.Debug(fmt.Sprintf("2*(%d - 2 ^( %v ) ) + 1", input, math.Floor(math.Log2(float64(input)))))
	log.Debug(fmt.Sprintf("2*(%d - %v ) + 1", input, math.Pow(2, math.Floor(math.Log2(float64(input))))))
	log.Debug(fmt.Sprintf("2*(%d) + 1", input-int(math.Pow(2, math.Floor(math.Log2(float64(input)))))))
	result = 2*(input-int(math.Pow(2, math.Floor(math.Log2(float64(input)))))) + 1

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
