package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

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

func parseInput(input string) (data []bool, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	for _, c := range tmpData {
		if c == '.' {
			data = append(data, true)
		} else if c == '^' {
			data = append(data, false)
		}
	}
	return
}

func tileType(i int, previous []bool) (r bool) {
	r = true
	var left, center, right bool
	if i > 0 {
		left = previous[i-1]
	} else {
		left = true
	}
	if i < len(previous)-1 {
		right = previous[i+1]
	} else {
		right = true
	}
	center = previous[i]
	if !right && !center && left {
		r = false
	}
	if right && !center && !left {
		r = false
	}
	if right && center && !left {
		r = false
	}
	if !right && center && left {
		r = false
	}
	return
}

func printTiles(tiles [][]bool) (result string) {
	for i := 0; i < len(tiles); i++ {
		result = fmt.Sprintf("[%04d] ", i)
		for j := 0; j < len(tiles[i]); j++ {
			if tiles[i][j] {
				result += "."
			} else {
				result += "^"
			}
		}
		fmt.Println(result)
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	data, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	rows := context.Int("rows")
	tiles := make([][]bool, rows)
	tiles[0] = data
	for j := 0; j < len(tiles[0]); j++ {
		if tiles[0][j] {
			result++
		}
	}
	for i := 1; i < rows; i++ {
		tiles[i] = make([]bool, len(tiles[i-1]))

		log.Info(fmt.Sprintf("Generating row %05d", i))
		for j := 0; j < len(tiles[i]); j++ {
			tiles[i][j] = tileType(j, tiles[i-1])
			if tiles[i][j] {
				result++
			}
		}
	}
	log.Debug(fmt.Sprintf("Generated tiles are:\n%s", printTiles(tiles)))

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
				Name:    "rows",
				Aliases: []string{"r"},
				Value:   40,
				Usage:   "Number of rows to be computed",
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
