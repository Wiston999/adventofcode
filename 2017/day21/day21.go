package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Coord struct {
	I int `json:"i"`
	J int `json:"j"`
}

type Pair struct {
	S Coord `json:"s"`
	D Coord `json:"d"`
}

type Transformation []Pair

func (t *Transformation) Apply(block map[Coord]bool) (nb map[Coord]bool) {
	nb = make(map[Coord]bool)
	for _, a := range *t {
		if v, ok := block[a.S]; ok && v {
			nb[a.D] = v
		}
	}
	return
}

type Transformations struct {
	Rotations []Transformation `json:"rotations"`
	Flips     []Transformation `json:"flips"`
}

type Problem struct {
	Canvas map[Coord]bool    `json:"-"`
	Rules  map[string]string `json:"-"`
	Size   int               `json:"-"`
	T2     Transformations   `json:"2"`
	T3     Transformations   `json:"3"`
}

func (p *Problem) ParseCanvas(input string) {
	p.Canvas = make(map[Coord]bool)
	for i, l := range strings.Split(input, "/") {
		for j, c := range l {
			// Use sparse matrix
			if c == '#' {
				p.Canvas[Coord{i, j}] = true
			}
			p.Size = int(math.Max(float64(p.Size), float64(j+1)))
		}
		p.Size = int(math.Max(float64(p.Size), float64(i+1)))
	}
}

func (p *Problem) BuildBlock(i, j, bs int) map[Coord]bool {
	block := make(map[Coord]bool)
	// Not very orthodox, but IDC
	coords := []Coord{
		Coord{i, j},
		Coord{i + 1, j},
		Coord{i, j + 1},
		Coord{i + 1, j + 1},
	}
	if bs == 3 {
		coords = append(coords, Coord{i + 2, j})
		coords = append(coords, Coord{i, j + 2})
		coords = append(coords, Coord{i + 2, j + 1})
		coords = append(coords, Coord{i + 1, j + 2})
		coords = append(coords, Coord{i + 2, j + 2})
	}
	for _, c := range coords {
		if v, ok := p.Canvas[c]; ok {
			nc := Coord{c.I % bs, c.J % bs}
			block[nc] = v
		}
	}
	return block
}

func (p *Problem) ToList(bs int, block map[Coord]bool) (l []string) {
	l = make([]string, bs)
	for i := 0; i < bs; i++ {
		f := ""
		for j := 0; j < bs; j++ {
			if _, ok := block[Coord{i, j}]; ok {
				f += "#"
			} else {
				f += "."
			}
		}
		l[i] = f
	}
	return
}

func (p *Problem) Flatten(bs int, block map[Coord]bool) (r string) {
	return strings.Join(p.ToList(bs, block), "/")
}

func (p *Problem) GetRule(bs int, block map[Coord]bool) (rule string) {
	T := p.T3
	if bs == 2 {
		T = p.T2
	}
	for _, r := range T.Rotations {
		for _, f := range T.Flips {
			var ok bool
			if rule, ok = p.Rules[p.Flatten(bs, r.Apply(f.Apply(block)))]; ok {
				return
			}
		}
	}
	return
}

func (p *Problem) Enhance() {
	bs := 3
	if p.Size%2 == 0 {
		bs = 2
	}

	blocks := p.Size / bs
	nc := make(map[Coord]bool)

	for i := 0; i < blocks; i++ {
		for j := 0; j < blocks; j++ {
			block := p.BuildBlock(i*bs, j*bs, bs)
			rule := p.GetRule(bs, block)
			if rule != "" {
				for x, l := range strings.Split(rule, "/") {
					for y, c := range l {
						if c == '#' {
							nc[Coord{i*(bs+1) + x, j*(bs+1) + y}] = true
						}
					}
				}
			} else {
				log.Warn(fmt.Sprintf("[%02d, %02d] No tranformation rule found for %s", i, j, p.Flatten(bs, block)))
			}
		}
	}
	p.Size = blocks * (bs + 1)
	p.Canvas = nc
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
	data.Rules = make(map[string]string)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		parts := strings.Split(l, " => ")
		data.Rules[parts[0]] = parts[1]
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	var iterations = context.Int("iterations")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	t, _ := os.ReadFile("transformations.json")
	json.Unmarshal(t, &problem)
	problem.ParseCanvas(".#./..#/###")
	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	log.Info(fmt.Sprintf("Applying %d iterations", iterations))
	for i := 0; i < iterations; i++ {
		problem.Enhance()
		log.Info(fmt.Sprintf("[%02d] Iteration: %d", i, len(problem.Canvas)))
	}
	result = len(problem.Canvas) // As it is a sparse matrix, size is also the number of activated pixels

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
				Name:  "iterations",
				Value: 5,
				Usage: "Number of iterations",
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
