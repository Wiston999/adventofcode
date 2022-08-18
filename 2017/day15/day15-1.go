package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const MODULUS = 2147483647

type Generator struct {
	Current, Factor uint
}

func (g *Generator) Next() uint {
	g.Current *= g.Factor
	g.Current %= MODULUS
	return g.Current
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

	A, B := Generator{uint(context.Int("a")), 16807}, Generator{uint(context.Int("b")), 48271}
	log.Debug(fmt.Sprintf("Parsed problem A: %v B: %v", A, B))

	limit := context.Int("limit")
	progress := limit / 20
	for i := 0; i < limit; i++ {
		a, b := fmt.Sprintf("%016b", A.Next()), fmt.Sprintf("%016b", B.Next())

		if a[len(a)-16:] == b[len(b)-16:] {
			result++
		}
		if i%progress == 0 {
			log.Info(fmt.Sprintf("Processed %d elements: %d", i, result))
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
			&cli.IntFlag{
				Name:  "a",
				Value: 783,
				Usage: "Generator A seed",
			},
			&cli.IntFlag{
				Name:  "b",
				Value: 325,
				Usage: "Generator B seed",
			},
			&cli.IntFlag{
				Name:  "limit",
				Value: 40000000,
				Usage: "Number of iterations",
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
