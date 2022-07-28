package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Range struct {
	Min, Max int
}

type RangeList []Range

func (rl RangeList) Len() int {
	return len(rl)
}

func (rl RangeList) Less(i, j int) bool {
	return rl[i].Min <= rl[j].Min
}

func (rl RangeList) Swap(i, j int) {
	rl[i], rl[j] = rl[j], rl[i]
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

func parseInput(input string) (data RangeList, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	for _, l := range strings.Split(strings.TrimSpace(string(tmpData)), "\n") {
		l = strings.TrimSuffix(l, "\n")
		ranges := strings.Split(l, "-")
		r := Range{Min: atoi(ranges[0]), Max: atoi(ranges[1])}
		data = append(data, r)
		log.Debug(fmt.Sprintf("Parsed range %#v", r))
	}
	return
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func merge(r Range, ranges []Range) (n Range, merged []int) {
	n.Min = r.Min
	n.Max = r.Max
	for i, ra := range ranges {
		if ra.Min == n.Min && ra.Max == n.Max {
			continue
		}
		if ra.Min >= n.Min && ra.Min <= n.Max && ra.Max > n.Max {
			n.Max = ra.Max
			merged = append(merged, i)
		}
		if ra.Max <= n.Max && ra.Max >= n.Min && ra.Min < n.Min {
			n.Min = ra.Min
			merged = append(merged, i)
		}
	}
	sort.Ints(merged)
	reverse(merged)
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	data, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	sort.Sort(RangeList(data))
	var merged RangeList

	for i := 0; i < len(data); {
		r := data[i]
		nr := Range{Min: r.Min, Max: r.Max}
		j := i + 1
		for ; j < len(data) && data[j].Min <= (nr.Max+1); j++ {
			if nr.Max < data[j].Max {
				nr.Max = data[j].Max
			}
			log.Debug(fmt.Sprintf("Merged %#v into %#v", data[j], nr))
		}
		i = j
		merged = append(merged, nr)
	}

	log.Debug(fmt.Sprintf("Merged %d ranges into %d", len(data), len(merged)))
	for i, r := range merged {
		if i > 0 {
			result += (r.Min - merged[i-1].Max - 1)
		}
		log.Debug(fmt.Sprintf("Final range %#v: %d", r, result))
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
