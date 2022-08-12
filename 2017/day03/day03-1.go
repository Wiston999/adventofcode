package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Problem struct {
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

	ring := int(math.Ceil(math.Sqrt(float64(input))) / 2)
	log.Debug(fmt.Sprintf("Number is in ring %d", ring))
	top, bottom, left, right := 1, 1, 1, 1

	for i := 0; i < ring; i++ {
		right += i*8 + 1
		top += i*8 + 3
		left += i*8 + 5
		bottom += i*8 + 7
	}

	closest := 100000.0
	if math.Abs(float64(right-input)) < closest {
		closest = math.Abs(float64(right - input))
	}
	if math.Abs(float64(top-input)) < closest {
		closest = math.Abs(float64(top - input))
	}
	if math.Abs(float64(left-input)) < closest {
		closest = math.Abs(float64(left - input))
	}
	if math.Abs(float64(bottom-input)) < closest {
		closest = math.Abs(float64(bottom - input))
	}
	log.Debug(fmt.Sprintf("Borders at ring %d are RTLB [%d %d %d %d] closest is %f", ring, right, top, left, bottom, closest))
	result = int(closest) + ring
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
