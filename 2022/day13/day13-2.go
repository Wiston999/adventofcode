package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

func sortedPair(L, R interface{}) (result int) {
	log.Debug(fmt.Sprintf("Comparing %v vs %v", L, R))
	if vL, ok := L.(float64); ok {
		if vR, ok := R.(float64); ok {
			if vL == vR {
				return 0
			} else if vL < vR {
				return 1
			} else {
				return -1
			}
		} else {
			return sortedPair([]interface{}{vL}, R)
		}
	} else {
		if vR, ok := R.(float64); ok {
			return sortedPair(L, []interface{}{vR})
		} else {
			LL := L.([]interface{})
			LR := R.([]interface{})
			for i := range LL {
				if len(LR) > i {
					cmp := sortedPair(LL[i], LR[i])
					if cmp == 1 || cmp == -1 {
						return cmp
					}
				}
			}
			if len(LL) < len(LR) {
				return 1
			} else if len(LL) > len(LR) {
				return -1
			}
		}
	}
	return
}

type Pair struct {
	Left, Right []interface{}
}

func (p *Pair) IsSorted() (result bool) {
	return sortedPair(p.Left, p.Right) == 1
}

type Problem struct {
	Pairs   []Pair
	Packets [][]interface{}
}

func (p *Problem) Sort() {
	n := len(p.Packets)
	for i := 0; i < n; i++ {
		for j := 0; j < (n - i - 1); j++ {
			if sortedPair(p.Packets[j], p.Packets[j+1]) != 1 {
				p.Packets[j], p.Packets[j+1] = p.Packets[j+1], p.Packets[j]
			}
		}
	}
}

func (p *Problem) Marker(packet interface{}) bool {
	if v1, ok := packet.([]interface{}); ok && len(v1) == 1 {
		if v2, ok := v1[0].([]interface{}); ok && len(v2) == 1 {
			if v3, ok := v2[0].(float64); ok {
				return v3 == 2 || v3 == 6
			}
		}
	}
	return false
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

func parseInput(input string) (data Problem, err error) {
	byteData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}
	strData := string(byteData)
	p := Pair{}
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		if i%3 == 2 {
			data.Pairs = append(data.Pairs, p)
			p = Pair{}
		} else if i%3 == 1 {
			json.Unmarshal([]byte(l), &p.Right)
			data.Packets = append(data.Packets, p.Right)
		} else {
			json.Unmarshal([]byte(l), &p.Left)
			data.Packets = append(data.Packets, p.Left)
		}
	}
	data.Pairs = append(data.Pairs, p)
	var tmp1, tmp2 []interface{}
	json.Unmarshal([]byte("[[2]]"), &tmp1)
	data.Packets = append(data.Packets, tmp1)
	json.Unmarshal([]byte("[[6]]"), &tmp2)
	data.Packets = append(data.Packets, tmp2)
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	problem.Sort()
	result = 1
	for i, p := range problem.Packets {
		log.Debug(fmt.Sprintf("[%03d] Packet: %v", i+1, p))
		if problem.Marker(p) {
			result *= i + 1
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
