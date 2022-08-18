package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type KnotHash struct {
	List                        []int
	Lengths                     []int
	Size, Current, Skip, Rounds int
}

func NewKnotHash(input string) (kh *KnotHash) {
	kh = new(KnotHash)
	kh.MakeList(256)
	kh.Lengths = make([]int, len(input))
	for i, c := range input {
		kh.Lengths[i] = int(c)
	}
	kh.Lengths = append(kh.Lengths, []int{17, 31, 73, 47, 23}...)
	kh.Rounds = 64
	return
}

func (p *KnotHash) MakeList(size int) {
	p.Size = size
	p.List = make([]int, p.Size)
	for i := range p.List {
		p.List[i] = i
	}
}

func (p *KnotHash) HashRound() {
	for _, n := range p.Lengths {
		p.Reverse(p.Current, n)
		p.Current += p.Skip + n
		p.Skip++
	}
}

func (p *KnotHash) Hash() (result string) {
	for i := 0; i < p.Rounds; i++ {
		p.HashRound()
	}

	xors := make([]int, 16)
	for i := 0; i < 16; i++ {
		xors[i] = p.List[i*16]
		for j := 1; j < 16; j++ {
			xors[i] = xors[i] ^ p.List[i*16+j]
		}
	}

	for i := range xors {
		result += fmt.Sprintf("%02x", xors[i])
	}

	return
}

func (p *KnotHash) Reverse(start, length int) {
	for i, j := start, start+length-1; i < j; i, j = i+1, j-1 {
		p.List[i%p.Size], p.List[j%p.Size] = p.List[j%p.Size], p.List[i%p.Size]
	}
}

type Coord struct {
	I, J int
}

type Problem struct {
	Grid map[Coord]bool
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

func hex2bin(h string) (b string) {
	for _, c := range h {
		switch c {
		case '0':
			b += "0000"
		case '1':
			b += "0001"
		case '2':
			b += "0010"
		case '3':
			b += "0011"
		case '4':
			b += "0100"
		case '5':
			b += "0101"
		case '6':
			b += "0110"
		case '7':
			b += "0111"
		case '8':
			b += "1000"
		case '9':
			b += "1001"
		case 'a':
			b += "1010"
		case 'b':
			b += "1011"
		case 'c':
			b += "1100"
		case 'd':
			b += "1101"
		case 'e':
			b += "1110"
		case 'f':
			b += "1111"
		}
	}
	return
}

func solution(context *cli.Context) (result int) {
	var input = context.String("input")

	log.Info(fmt.Sprintf("Input value: %s", input))

	problem := Problem{make(map[Coord]bool)}

	for i := 0; i < 128; i++ {
		d := fmt.Sprintf("%s-%d", input, i)
		kh := NewKnotHash(d)
		r := kh.Hash()
		b := hex2bin(r)
		log.Info(fmt.Sprintf("[%03d] Computed KnotHash of %s: [%s] %s", i, d, r, b))
		for j, c := range b {
			problem.Grid[Coord{i, j}] = (c == '1')
		}
	}

	for _, v := range problem.Grid {
		if v {
			result++
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
				Value:   "hfdlxzhv",
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
