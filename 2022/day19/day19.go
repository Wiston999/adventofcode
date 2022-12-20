package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Blueprint struct {
	OreOreCost        int
	ClayOreCost       int
	ObsidianOreCost   int
	ObsidianClayCost  int
	GeodeOreCost      int
	GeodeObsidianCost int
	MaxOre            int
	MaxClay           int
	MaxObsidian       int
	Best              int
}

func (b *Blueprint) Analyze() {
	b.MaxOre = b.OreOreCost
	b.MaxClay = b.ObsidianClayCost
	b.MaxObsidian = b.GeodeObsidianCost
	if b.ClayOreCost > b.MaxOre {
		b.MaxOre = b.ClayOreCost
	}
	if b.ObsidianOreCost > b.MaxOre {
		b.MaxOre = b.ObsidianOreCost
	}
	if b.GeodeOreCost > b.MaxOre {
		b.MaxOre = b.GeodeOreCost
	}
}

func (b *Blueprint) BestEstimate(turns, geodeBot int) (result int) {
	return (geodeBot+turns-1)*(geodeBot+turns)/2 - ((geodeBot - 1) * geodeBot / 2)
}

func (b *Blueprint) Quality(turns, ore, clay, obsidian, geode, oreBot, clayBot, obsidianBot, geodeBot int) {
	if turns == 0 || b.Best > geode+b.BestEstimate(turns, geodeBot) {
		if geode > b.Best {
			log.Debug(fmt.Sprintf("(base) New best for %#v: %d", b, geode))
			b.Best = geode
		}
		return
	}
	if oreBot >= b.GeodeOreCost && obsidianBot >= b.GeodeObsidianCost {
		if b.BestEstimate(turns, geodeBot) > b.Best {
			log.Debug(fmt.Sprintf("(shortcut) New best for %#v: %d", b, geode))
			b.Best = b.BestEstimate(turns, geodeBot)
		}
		return
	}
	if b.MaxOre > oreBot {
		b.Quality(
			turns-1,
			ore+oreBot,
			clay+clayBot,
			obsidian+obsidianBot,
			geode+geodeBot,
			oreBot,
			clayBot,
			obsidianBot,
			geodeBot,
		)
	}
	if ore >= b.OreOreCost && b.MaxOre > oreBot {
		b.Quality(
			turns-1,
			ore+oreBot-b.OreOreCost,
			clay+clayBot,
			obsidian+obsidianBot,
			geode+geodeBot,
			oreBot+1,
			clayBot,
			obsidianBot,
			geodeBot,
		)
	}
	if ore >= b.ClayOreCost && b.MaxClay > clayBot {
		b.Quality(
			turns-1,
			ore+oreBot-b.ClayOreCost,
			clay+clayBot,
			obsidian+obsidianBot,
			geode+geodeBot,
			oreBot,
			clayBot+1,
			obsidianBot,
			geodeBot,
		)
	}
	if ore >= b.ObsidianOreCost && clay >= b.ObsidianClayCost && b.MaxObsidian > obsidianBot {
		b.Quality(
			turns-1,
			ore+oreBot-b.ObsidianOreCost,
			clay+clayBot-b.ObsidianClayCost,
			obsidian+obsidianBot,
			geode+geodeBot,
			oreBot,
			clayBot,
			obsidianBot+1,
			geodeBot,
		)
	}
	if ore >= b.GeodeOreCost && obsidian >= b.GeodeObsidianCost {
		b.Quality(
			turns-1,
			ore+oreBot-b.GeodeOreCost,
			clay+clayBot,
			obsidian+obsidianBot-b.GeodeObsidianCost,
			geode+geodeBot,
			oreBot,
			clayBot,
			obsidianBot,
			geodeBot+1,
		)
	}
}

type Problem struct {
	Blueprints []Blueprint
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
	data.Blueprints = make([]Blueprint, 0)
	for i, l := range strings.Split(strings.TrimSpace(strData), "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		b := Blueprint{}
		c, err := fmt.Sscanf(
			l,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&i,
			&b.OreOreCost,
			&b.ClayOreCost,
			&b.ObsidianOreCost,
			&b.ObsidianClayCost,
			&b.GeodeOreCost,
			&b.GeodeObsidianCost,
		)
		if err != nil {
			log.Warn(fmt.Sprintf("Error parsing line [%d]: %s", c, err))
		}
		data.Blueprints = append(data.Blueprints, b)
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

	log.Debug(fmt.Sprintf("Parsed problem %#v", problem))
	results := make(chan int)
	var wg sync.WaitGroup
	turns := 24
	part2 := context.Bool("part2")
	if part2 {
		turns = 32
		result = 1
	}

	if part2 && len(problem.Blueprints) > 3 {
		wg.Add(3)
	} else {
		wg.Add(len(problem.Blueprints))
	}
	for i, b := range problem.Blueprints {
		if context.Bool("part2") && i > 2 {
			break
		}
		go func(i int, b Blueprint) {
			defer wg.Done()
			b.Analyze()
			b.Quality(turns, 0, 0, 0, 0, 1, 0, 0, 0)
			log.Info(fmt.Sprintf("[%02d] %#v Quality: %d", i+1, b, b.Best))
			if part2 {
				results <- b.Best
			} else {
				results <- b.Best * (i + 1)
			}
		}(i, b)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	for r := range results {
		if part2 {
			result *= r
		} else {
			result += r
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
			&cli.BoolFlag{
				Name:    "part2",
				Aliases: []string{"2"},
				Value:   false,
				Usage:   "Part 2 mode (32 minutes and 3 blueprints)",
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
