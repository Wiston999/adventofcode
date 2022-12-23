package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	DIRT = iota
	ELF
)

type Ground int

type Coord struct {
	I, J int
}

type Problem struct {
	Grid        map[Coord]Ground
	CurrentMove int
	MinI, MaxI  int
	MinJ, MaxJ  int
}

func (p *Problem) CountDirt() int {
	return (p.MaxI-p.MinI+1)*(p.MaxJ-p.MinJ+1) - len(p.Grid)
}

func (p *Problem) Print() (result string) {
	for i := p.MinI; i <= p.MaxI; i++ {
		for j := p.MinJ; j <= p.MaxJ; j++ {
			if p.Grid[Coord{i, j}] == ELF {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return
}

func (p *Problem) Step() (moved bool) {
	proposed := make(map[Coord][]Coord)
	for k, g := range p.Grid {
		if g == ELF {
			alone := true
			for i := k.I - 1; i <= k.I+1 && alone; i++ {
				for j := k.J - 1; j <= k.J+1; j++ {
					if i == k.I && j == k.J {
						continue
					}
					if p.Grid[Coord{i, j}] == ELF {
						alone = false
						break
					}
				}
			}
			if !alone {
				for i := 0; i < 4; i++ {
					if (i+p.CurrentMove)%4 == 0 && p.Grid[Coord{k.I - 1, k.J - 1}] != ELF && p.Grid[Coord{k.I - 1, k.J}] != ELF && p.Grid[Coord{k.I - 1, k.J + 1}] != ELF {
						log.Debug(fmt.Sprintf("%v proposed to %v", k, Coord{k.I - 1, k.J}))
						proposed[Coord{k.I - 1, k.J}] = append(proposed[Coord{k.I - 1, k.J}], k)
						break
					}
					if (i+p.CurrentMove)%4 == 1 && p.Grid[Coord{k.I + 1, k.J - 1}] != ELF && p.Grid[Coord{k.I + 1, k.J}] != ELF && p.Grid[Coord{k.I + 1, k.J + 1}] != ELF {
						log.Debug(fmt.Sprintf("%v proposed to %v", k, Coord{k.I + 1, k.J}))
						proposed[Coord{k.I + 1, k.J}] = append(proposed[Coord{k.I + 1, k.J}], k)
						break
					}
					if (i+p.CurrentMove)%4 == 2 && p.Grid[Coord{k.I - 1, k.J - 1}] != ELF && p.Grid[Coord{k.I, k.J - 1}] != ELF && p.Grid[Coord{k.I + 1, k.J - 1}] != ELF {
						log.Debug(fmt.Sprintf("%v proposed to %v", k, Coord{k.I, k.J - 1}))
						proposed[Coord{k.I, k.J - 1}] = append(proposed[Coord{k.I, k.J - 1}], k)
						break
					}
					if (i+p.CurrentMove)%4 == 3 && p.Grid[Coord{k.I - 1, k.J + 1}] != ELF && p.Grid[Coord{k.I, k.J + 1}] != ELF && p.Grid[Coord{k.I + 1, k.J + 1}] != ELF {
						log.Debug(fmt.Sprintf("%v proposed to %v", k, Coord{k.I, k.J + 1}))
						proposed[Coord{k.I, k.J + 1}] = append(proposed[Coord{k.I, k.J + 1}], k)
						break
					}
				}
			}
		}
	}

	for d, l := range proposed {
		if len(l) == 1 {
			p.Grid[d] = ELF
			delete(p.Grid, l[0])
			moved = true
		} else {
			log.Debug(fmt.Sprintf("%v proposed by %d elves", d, len(l)))
		}
	}
	p.UpdateBoundaries()
	p.CurrentMove++
	if log.GetLevel() == log.DebugLevel {
		log.Debug("\n" + p.Print())
	}
	return
}

func (p *Problem) UpdateBoundaries() {
	p.MinI, p.MaxI = 999999, 0
	p.MinJ, p.MaxJ = 999999, 0
	for k := range p.Grid {
		if k.I < p.MinI {
			p.MinI = k.I
		}
		if k.I > p.MaxI {
			p.MaxI = k.I
		}
		if k.J < p.MinJ {
			p.MinJ = k.J
		}
		if k.J > p.MaxJ {
			p.MaxJ = k.J
		}
	}
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
	data.Grid = make(map[Coord]Ground)
	data.MinI, data.MinJ = 99999, 99999
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		for j, c := range l {
			if c == '#' {
				data.Grid[Coord{i, j}] = ELF
				if i < data.MinI {
					data.MinI = i
				}
				if i > data.MaxI {
					data.MaxI = i
				}
				if j < data.MinJ {
					data.MinJ = j
				}
				if j > data.MaxJ {
					data.MaxJ = j
				}
			}
		}
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	steps := context.Int("steps")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	i := 0
	for ; i < steps && problem.Step(); i++ {
		log.Info(fmt.Sprintf("[%03d] Step completed", i+1))
	}
	log.Warn(fmt.Sprintf("No more moves after step %d", i+1))
	log.Debug(fmt.Sprintf("Solved problem:\n%s", problem.Print()))
	result = problem.CountDirt()
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
				Name:    "steps",
				Aliases: []string{"s"},
				Value:   10,
				Usage:   "Number of iterations, use a very high number (like 1000000) and solution for part 2 will be printed in warn level",
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
