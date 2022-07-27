package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

var BestScore int
var Best []byte

type Coord struct {
	X, Y int
}

type Map struct {
	Target     Coord
	Path, Salt []byte
}

func sliceCopy(s []byte) (d []byte) {
	for _, c := range s {
		d = append(d, c)
	}
	return
}

func (m *Map) Copy() (n Map) {
	n.Target = m.Target
	n.Path = sliceCopy(m.Path)
	n.Salt = sliceCopy(m.Salt)
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

func MD5(s []byte) (r string) {
	return fmt.Sprintf("%x", md5.Sum(s))
}

func open(s byte) (r bool) {
	switch s {
	case 'b', 'c', 'd', 'e', 'f':
		r = true
	}
	return
}

func doors(s []byte) (result [4]bool) {
	hash := MD5(s)
	for i := 0; i < 4; i++ {
		result[i] = open(hash[i])
	}
	return
}

// Manhattan
func heuristic(s, e Coord) int {
	return int(math.Abs(float64(s.X-e.X)) + math.Abs(float64(s.Y-e.Y)))
}

func indexToStep(i int) (s byte) {
	switch i {
	case 0:
		s = 'U'
	case 1:
		s = 'D'
	case 2:
		s = 'L'
	case 3:
		s = 'R'
	}
	return
}

func find(c Coord, s []byte, m Map) {
	solutionPath := bytes.TrimPrefix(s, m.Salt)
	if c == m.Target {
		if BestScore > len(solutionPath) {
			Best = solutionPath
			BestScore = len(solutionPath)
			log.Info(fmt.Sprintf("Found new best solution: %d (%s)", BestScore, Best))
		}
		return
	}

	cs := []Coord{
		Coord{c.X, c.Y - 1},
		Coord{c.X, c.Y + 1},
		Coord{c.X - 1, c.Y},
		Coord{c.X + 1, c.Y},
	}
	openDoors := doors(s)
	for i, nc := range cs {
		if nc.X < 0 || nc.Y < 0 || nc.X > m.Target.X || nc.Y > m.Target.Y {
			continue
		}
		if openDoors[i] && (len(solutionPath)+heuristic(nc, m.Target)) < BestScore {
			path := append(s, indexToStep(i))
			find(nc, path, m)
		}
	}
}

func solution(context *cli.Context) (result string) {
	var input = context.String("input")

	m := Map{
		Salt:   []byte(input),
		Target: Coord{3, 3},
	}
	BestScore = 10000000
	find(Coord{0, 0}, m.Salt, m)
	result = string(Best)
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
				Value:   "vwbaicqe",
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
			log.Debug(fmt.Sprintf("Input value:  %s", c.String("input")))
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
