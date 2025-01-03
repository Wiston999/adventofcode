package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Problem struct {
	Data []string
}

func decode(s string) (i int) {
	k := 0
	for j := len(s) - 1; j >= 0; j-- {
		v := 0
		switch s[j] {
		case '=':
			v = -2
		case '-':
			v = -1
		case '0':
			v = 0
		case '1':
			v = 1
		case '2':
			v = 2
		}
		i += v * int(math.Pow(5.0, float64(k)))
		k++
	}

	return
}

func digit(i int) (s string) {
	switch i {
	case -2:
		s = "="
	case -1:
		s = "-"
	case 0:
		s = "0"
	case 1:
		s = "1"
	case 2:
		s = "2"
	}
	return
}

func encode(i int) (s string) {
	var tmp string
	for ; i > 0; i /= 5 {
		r := i % 5
		if r > 2 {
			i += r
			tmp += digit(r - 5)
		} else {
			tmp += digit(r)
		}
	}
	for j := len(tmp) - 1; j >= 0; j-- {
		s += string(tmp[j])
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

func atoiSafe(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
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
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		data.Data = append(data.Data, l)
	}
	return
}

func solution(context *cli.Context) (result string) {
	var input = context.String("input")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	value := 0
	for i, n := range problem.Data {
		v := decode(n)
		log.Debug(fmt.Sprintf("[%03d] Decoded value %d", i, v))
		value += v
	}

	log.Info(fmt.Sprintf("Added values: %d", value))
	result = encode(value)

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
			start := time.Now()
			echo(fmt.Sprintf("Solution is %v in %s", solution(c), time.Since(start)), c.String("output"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
