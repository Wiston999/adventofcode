package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Packet struct {
	Data    []string
	Control []string
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

func parseInput(input string) (data []Packet, err error) {
	tmpData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	for i, l := range strings.Split(strings.TrimSpace(string(tmpData)), "\n") {
		log.Debug(fmt.Sprintf("Line: %#v", l))
		packet := Packet{}
		for j, s1 := range strings.Split(strings.TrimSpace(l), "[") {
			for k, s2 := range strings.Split(strings.TrimSpace(s1), "]") {
				if j == 0 {
					packet.Data = append(packet.Data, string(s2))
				} else {
					if k == 0 {
						packet.Control = append(packet.Control, string(s2))
					} else {
						packet.Data = append(packet.Data, string(s2))
					}
				}
			}
		}
		data = append(data, packet)
		log.Debug(fmt.Sprintf("Parsed packet %d: %#v", i, packet))
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

	for i, p := range data {
		found := false
		for k, c := range p.Data {
			for j := 0; j < len(c)-2 && !found; j++ {
				if c[j] == c[j+2] && c[j] != c[j+1] {
					palindrome := string(c[j]) + string(c[j+1]) + string(c[j])
					log.Debug(fmt.Sprintf("Found palindrome on data sequence (%d) %s of packet (%d) %#v", k, c, i, palindrome))
					palindrome = string(c[j+1]) + string(c[j]) + string(c[j+1])
					for l, control := range p.Control {
						if strings.Contains(control, palindrome) {
							log.Info(fmt.Sprintf("Found reverse palindrome on control sequence (%d) %s of packet (%d) %#v", l, control, i, palindrome))
							found = true
							break
						}
					}
				}
			}
			if found {
				result++
				break
			}
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
