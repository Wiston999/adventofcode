// Misserably translated from https://gist.github.com/bluepichu/59c815b132c0e9ad29e4df32c5cddfbd
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Valve struct {
	Name        string
	Flow        int
	Open        bool
	Connections []string
}

type Problem struct {
	AllValves     map[string]Valve
	Paths         map[string]map[string]int
	PathsInt      [][]int
	AllValvesInt  []Valve
	Keys          []string
	DP            [][][]int
	Result, Limit int
}

type Node struct {
	Name         string
	Index, Level int
}

func (p *Problem) GetFlow(m int) (result int) {
	for i := range p.Keys {
		if ((1 << i) & m) != 0 {
			result += p.AllValvesInt[i].Flow
		}
	}
	return
}

func (p *Problem) ComputeFlows() {
	p.DP = make([][][]int, p.Limit+1)
	for n := 0; n < p.Limit+1; n++ {
		p.DP[n] = make([][]int, len(p.Keys))
		for i := range p.Keys {
			p.DP[n][i] = make([]int, 1<<len(p.Keys))
			for k := 0; k < (1 << len(p.Keys)); k++ {
				p.DP[n][i][k] = -9999999999
			}
		}
	}

	for i, v := range p.Keys {
		dist := p.Paths["AA"][v]
		p.DP[dist+1][i][1<<i] = 0
	}

	for i := 1; i <= p.Limit; i++ {
		for j := 0; j < (1 << len(p.Keys)); j++ {
			for k := range p.Keys {
				flow := p.GetFlow(j)

				hold := p.DP[i-1][k][j] + flow
				if hold > p.DP[i][k][j] {
					p.DP[i][k][j] = hold
				}

				if p.Result < p.DP[i][k][j] {
					p.Result = p.DP[i][k][j]
				}

				if ((1 << k) & j) == 0 {
					continue
				}

				for l := range p.Keys {
					if ((1 << l) & j) != 0 {
						continue
					}
					dist := p.PathsInt[k][l]
					if (i + dist) >= p.Limit {
						continue
					}

					value := p.DP[i][k][j] + flow*(dist+1)
					if value > p.DP[i+dist+1][l][j|(1<<l)] {
						p.DP[i+dist+1][l][j|(1<<l)] = value
					}
				}
			}
		}
	}
}

func (p *Problem) ComputePaths() {
	p.Paths = make(map[string]map[string]int)
	for k1, v1 := range p.AllValves {
		if v1.Flow == 0 && k1 != "AA" {
			continue
		}
		p.Paths[k1] = make(map[string]int)
		for k2, v2 := range p.AllValves {
			if v2.Flow == 0 && k2 != "AA" {
				continue
			}
			distance := 999999999
			if k1 == k2 {
				distance = 0
			}
			p.Paths[k1][k2] = distance
		}
	}

	p.Keys = make([]string, len(p.Paths))
	p.PathsInt = make([][]int, len(p.Paths))
	p.AllValvesInt = make([]Valve, len(p.Paths))
	k2i := make(map[string]int)
	i := 0
	for k := range p.Paths {
		p.Keys[i] = k
		k2i[k] = i
		p.PathsInt[i] = make([]int, len(p.Paths))
		i++
	}
	i = 0
	for _, k := range p.Keys {
		visited := make(map[string]bool)
		heap := make([]Node, 0)
		heap = append(heap, Node{k, i, 0})
		p.AllValvesInt[i] = p.AllValves[k]
		for len(heap) > 0 {
			current := heap[0]
			log.Debug(fmt.Sprintf("[%s] Exploring %v (%#v)", k, current.Name, visited))
			visited[current.Name] = true
			heap = heap[1:]
			if current.Level < p.Paths[k][current.Name] {
				p.Paths[k][current.Name] = current.Level
				p.PathsInt[i][current.Index] = current.Level
			}
			for _, c := range p.AllValves[current.Name].Connections {
				if !visited[c] {
					heap = append(heap, Node{c, k2i[c], current.Level + 1})
				}
			}
		}
		for k2 := range p.Paths[k] {
			log.Info(fmt.Sprintf("%s --> %s = %02d", k, k2, p.Paths[k][k2]))
		}
		i++
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

func Atoi(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
}

func parseInput(input string) (data Problem, err error) {
	byteData, err := os.ReadFile(input)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening file %s for reading input: %v", input, err))
		return
	}

	data.AllValves = make(map[string]Valve)
	data.Paths = make(map[string]map[string]int)
	strData := string(byteData)
	regex := *regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)
	res := regex.FindAllStringSubmatch(strData, -1)
	for i, l := range res {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		v := Valve{
			Name: l[1],
			Flow: Atoi(l[2]),
		}
		other := l[3]
		log.Debug(fmt.Sprintf("[%03d] Parsing other %s", i, other))
		for _, o := range strings.Split(other, ", ") {
			ov, _ := data.AllValves[o]
			v.Connections = append(v.Connections, o)
			data.AllValves[o] = ov
		}
		data.AllValves[v.Name] = v
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")
	problem, err := parseInput(input)
	if err != nil {
		log.Error(fmt.Sprintf("Something went wrong while reading input file: %v", err))
		return
	}

	problem.Limit = 30
	problem.ComputePaths()
	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	problem.ComputeFlows()
	result = problem.Result

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
