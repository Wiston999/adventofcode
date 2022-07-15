package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Room struct {
	Sector   int
	Checksum string
	Letters  string
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

func parseInput(input string) (data []Room, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	regex := *regexp.MustCompile(`(?m)^([\w\-]+?)(\d+)\[(\w+)\]$`)
	res := regex.FindAllStringSubmatch(string(tmpData), -1)
	for _, l := range res {
		s, _ := strconv.Atoi(string(l[2]))
		room := Room{
			s,
			l[3],
			strings.Replace(l[1], "-", "", -1),
		}
		data = append(data, room)
	}

	return
}

func countLetters(r Room) (result map[string]int) {
	result = make(map[string]int)
	for _, l := range r.Letters {
		ls := string(l)
		if _, ok := result[ls]; ok {
			result[ls]++
		} else {
			result[ls] = 1
		}
	}
	return
}

func max(m map[string]int, limit int) (max int) {
	for _, v := range m {
		if v > max && v < limit {
			max = v
		}
	}
	return
}

func buildChecksum(counters map[string]int) (c string) {
	limit := 10000
	for {
		current := max(counters, limit)
		if current == 0 {
			break
		}
		letters := []string{}
		for k, v := range counters {
			if v == current {
				letters = append(letters, k)
			}
		}
		sort.Strings(letters)
		c += strings.Join(letters, "")
		limit = current
	}
	return
}

func checksum(r Room, counters map[string]int) bool {
	c := buildChecksum(counters)
	log.Debug(fmt.Sprintf("Built checksum %s from %#v for room %#v", c, counters, r))
	return strings.Contains(c, r.Checksum)
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	data, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	for i, r := range data {
		roomCount := countLetters(r)
		check := checksum(r, roomCount)
		log.Debug(fmt.Sprintf("Checked room (%02d) %#v: %v", i, r, check))
		if check {
			result += r.Sector
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
			echo(fmt.Sprintf("Solution is %v", solution(c)), c.String("output"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
