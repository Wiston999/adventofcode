package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
	NULL
)

const (
	WALL  = '#'
	EMPTY = '.'
)

type Coord struct {
	I, J int
}

type Direction int

type Sticker struct {
	Value                 rune
	Up, Down, Left, Right Position
}

type Position struct {
	P Coord
	D Direction
}

type Action struct {
	D     Direction
	Steps int
}

type Pair struct {
	F, L int
}

type Problem struct {
	Grid    map[Coord]*Sticker
	Size    int
	Actions []Action
	Current Position
}

var movements = map[Position][]Position{
	Position{Coord{0, 0}, RIGHT}: []Position{
		Position{Coord{2, 1}, RIGHT},
		Position{Coord{0, 1}, LEFT},
	},
	Position{Coord{0, 0}, LEFT}: []Position{
		Position{Coord{1, -1}, UP},
		Position{Coord{2, -1}, LEFT},
	},
	Position{Coord{0, 0}, UP}: []Position{
		Position{Coord{1, -2}, UP},
		Position{Coord{3, -1}, LEFT},
	},
	Position{Coord{0, 0}, DOWN}: []Position{
		Position{Coord{1, 0}, UP},
	},
	Position{Coord{0, 1}, RIGHT}: []Position{
		Position{Coord{2, 0}, RIGHT},
		Position{Coord{2, 1}, RIGHT},
	},
	Position{Coord{0, 1}, LEFT}: []Position{
		Position{Coord{0, 0}, RIGHT},
	},
	Position{Coord{0, 1}, UP}: []Position{
		Position{Coord{3, -1}, DOWN},
	},
	Position{Coord{0, 1}, DOWN}: []Position{
		Position{Coord{1, 0}, RIGHT},
	},
	Position{Coord{1, 0}, RIGHT}: []Position{
		Position{Coord{0, 1}, DOWN},
		Position{Coord{2, 1}, UP},
	},
	Position{Coord{1, 0}, LEFT}: []Position{
		Position{Coord{2, -1}, UP},
	},
	Position{Coord{1, 0}, UP}: []Position{
		Position{Coord{0, 0}, DOWN},
	},
	Position{Coord{1, 0}, DOWN}: []Position{
		Position{Coord{2, 0}, UP},
	},
	Position{Coord{2, 0}, RIGHT}: []Position{
		Position{Coord{2, 1}, LEFT},
		Position{Coord{0, 1}, RIGHT},
	},
	Position{Coord{2, 0}, LEFT}: []Position{
		Position{Coord{2, -1}, RIGHT},
		Position{Coord{1, -1}, DOWN},
	},
	Position{Coord{2, 0}, UP}: []Position{
		Position{Coord{1, 0}, DOWN},
	},
	Position{Coord{2, 0}, DOWN}: []Position{
		Position{Coord{1, -2}, DOWN},
		Position{Coord{3, -1}, RIGHT},
	},
	Position{Coord{1, -1}, RIGHT}: []Position{
		Position{Coord{1, 0}, LEFT},
	},
	Position{Coord{1, -1}, LEFT}: []Position{
		Position{Coord{1, -2}, RIGHT},
	},
	Position{Coord{1, -1}, UP}: []Position{
		Position{Coord{0, 0}, LEFT},
	},
	Position{Coord{1, -1}, DOWN}: []Position{
		Position{Coord{2, 0}, LEFT},
	},
	Position{Coord{1, -2}, RIGHT}: []Position{
		Position{Coord{1, -1}, LEFT},
	},
	Position{Coord{1, -2}, LEFT}: []Position{
		Position{Coord{2, 1}, DOWN},
	},
	Position{Coord{1, -2}, UP}: []Position{
		Position{Coord{0, 0}, UP},
	},
	Position{Coord{1, -2}, DOWN}: []Position{
		Position{Coord{2, 0}, DOWN},
	},
	Position{Coord{2, -1}, RIGHT}: []Position{
		Position{Coord{2, 0}, LEFT},
	},
	Position{Coord{2, -1}, LEFT}: []Position{
		Position{Coord{0, 0}, LEFT},
	},
	Position{Coord{2, -1}, UP}: []Position{
		Position{Coord{1, 0}, LEFT},
	},
	Position{Coord{2, -1}, DOWN}: []Position{
		Position{Coord{3, -1}, UP},
	},
	Position{Coord{2, 1}, RIGHT}: []Position{
		Position{Coord{0, 0}, RIGHT},
	},
	Position{Coord{2, 1}, LEFT}: []Position{
		Position{Coord{2, 0}, RIGHT},
	},
	Position{Coord{2, 1}, UP}: []Position{
		Position{Coord{1, 0}, RIGHT},
	},
	Position{Coord{2, 1}, DOWN}: []Position{
		Position{Coord{1, -2}, LEFT},
	},
	Position{Coord{3, -1}, RIGHT}: []Position{
		Position{Coord{2, 0}, DOWN},
	},
	Position{Coord{3, -1}, LEFT}: []Position{
		Position{Coord{0, 0}, UP},
	},
	Position{Coord{3, -1}, UP}: []Position{
		Position{Coord{2, -1}, DOWN},
	},
	Position{Coord{3, -1}, DOWN}: []Position{
		Position{Coord{0, 1}, UP},
	},
}

func (p *Problem) FindNeighbour(c Coord, d Direction, size int) (result Position) {
	baseI, baseJ := p.Current.P.I/size, p.Current.P.J/size
	offsetI, offsetJ := c.I/size-baseI, c.J/size-baseJ
	check := Position{P: c}
	switch d {
	case UP:
		check.P.I--
	case DOWN:
		check.P.I++
	case RIGHT:
		check.P.J++
	case LEFT:
		check.P.J--
	}
	if _, ok := p.Grid[check.P]; ok && check.P != c {
		return check
	}
	for _, m := range movements[Position{Coord{offsetI, offsetJ}, d}] {
		switch d {
		case UP:
			switch m.D {
			case DOWN:
				check.P.I = m.P.I*size + mod(check.P.I, size)
				check.P.J = (baseJ+m.P.J)*size + mod(check.P.J, size)
			case UP:
				check.P.I = m.P.I*size + mod(c.I, size)
				check.P.J = (baseJ+m.P.J)*size + (size - 1 - mod(c.J, size))
				check.D = -2
			case LEFT:
				check.P.I = m.P.I*size + mod(c.J, size)
				check.P.J = (baseJ+m.P.J)*size - mod(c.I, size)
				check.D = 1
			case RIGHT:
				check.P.I = m.P.I*size + mod(c.J, size) + 1
				check.P.J = (baseJ+m.P.J+1)*size - 1
				check.D = 1
			}
		case DOWN:
			switch m.D {
			case DOWN:
				check.P.I = m.P.I*size + size - 1
				check.P.J = (m.P.J+baseJ)*size + (size - 1 - mod(c.J, size))
				check.D = -2
			case UP:
				check.P.I = m.P.I*size + mod(check.P.I, size)
				check.P.J = (baseJ+m.P.J)*size + mod(c.J, size)
			case LEFT:
				check.P.I = m.P.I*size + (size - 1 - mod(c.J, size))
				check.P.J = (m.P.J+baseJ)*size + (size - 1 - mod(c.I, size))
				check.D = 1
			case RIGHT:
				check.P.I = m.P.I*size + mod(c.J, size)
				check.P.J = (baseJ+m.P.J)*size + mod(c.I, size)
				check.D = 1
			}
		case RIGHT:
			switch m.D {
			case DOWN:
				check.P.I = m.P.I*size + mod(c.J, size)
				check.P.J = (baseJ+m.P.J)*size + mod(c.I, size)
				check.D = 3
			case UP:
				check.P.I = m.P.I*size + (size - 1 - mod(c.J, size))
				check.P.J = (baseJ+m.P.J)*size + (size - 1 - mod(c.I, size))
				check.D = 1
			case LEFT:
				check.P.I = m.P.I*size + check.P.I
				check.P.J = m.P.J*size + check.P.J - 1
			case RIGHT:
				check.P.I = m.P.I*size + (size - 1 - mod(c.I, size))
				check.P.J = (baseJ+m.P.J+1)*size + (size - 1 - mod(c.J, size)) - 1
				check.D = 2
			}
		case LEFT:
			switch m.D {
			case DOWN:
				check.P.I = m.P.I*size + (size - 1 - mod(c.J, size))
				check.P.J = (baseJ+m.P.J)*size + (size - 1 - mod(c.I, size))
				check.D = 1
			case UP:
				check.P.I = m.P.I*size + mod(c.J, size)
				check.P.J = (baseJ+m.P.J)*size + mod(c.I, size)
				check.D = -1
			case LEFT:
				check.P.I = m.P.I*size + (size - 1 - mod(c.I, size))
				check.P.J = (baseJ + m.P.J) * size
				check.D = -2
			case RIGHT:
				check.P.I = m.P.I*size + check.P.I
				check.P.J = m.P.J*size + check.P.J + 1
			}
		}
		log.Debug(fmt.Sprintf(
			"[%02d] [%v] Looking for neighbour of %v to %d at %v",
			size, p.Current, c, d, check,
		))
		if _, ok := p.Grid[check.P]; ok {
			return check
		}
	}
	return
}

func (p *Problem) Fold(size int) {
	for c, s := range p.Grid {
		s.Down = p.FindNeighbour(c, DOWN, size)
		s.Up = p.FindNeighbour(c, UP, size)
		s.Right = p.FindNeighbour(c, RIGHT, size)
		s.Left = p.FindNeighbour(c, LEFT, size)
		log.Debug(fmt.Sprintf("[%v] Folded %#v", c, s))
	}
}

func (p *Problem) Traverse() {
	for i, a := range p.Actions {
		log.Debug(fmt.Sprintf("[%03d] Applying action: %#v. At %#v", i+1, a, p.Current))
		p.ApplyAction(a)
		log.Debug(fmt.Sprintf("[%03d] Applied action: %#v. Now at %#v", i+1, a, p.Current))
	}
}

func (p *Problem) ApplyAction(a Action) {
	switch a.D {
	case NULL:
		var nextC, prevC Position
		prevC = p.Current
		nextC = p.Current
		for j := 0; j < a.Steps; j++ {
			log.Debug(fmt.Sprintf("[%03d] Applying step %#v %#v. At %#v", j, nextC, a, p.Grid[nextC.P]))
			switch nextC.D {
			case UP:
				nextC.D += p.Grid[nextC.P].Up.D
				nextC.P = p.Grid[nextC.P].Up.P
			case DOWN:
				nextC.D += p.Grid[nextC.P].Down.D
				nextC.P = p.Grid[nextC.P].Down.P
			case LEFT:
				nextC.D += p.Grid[nextC.P].Left.D
				nextC.P = p.Grid[nextC.P].Left.P
			case RIGHT:
				nextC.D += p.Grid[nextC.P].Right.D
				nextC.P = p.Grid[nextC.P].Right.P
			}
			nextC.D = Direction(mod(int(nextC.D), 4))
			if p.Grid[nextC.P].Value == WALL {
				log.Debug(fmt.Sprintf("[%03d] Found wall at %v", j, nextC))
				break
			}
			log.Debug(fmt.Sprintf("[%03d] Applied step %#v %#v. At %#v", j, nextC, a, p.Grid[nextC.P]))
			prevC = nextC
		}
		p.Current = prevC

	case RIGHT:
		switch p.Current.D {
		case RIGHT:
			p.Current.D = DOWN
		case LEFT:
			p.Current.D = UP
		case UP:
			p.Current.D = RIGHT
		case DOWN:
			p.Current.D = LEFT
		}
	case LEFT:
		switch p.Current.D {
		case RIGHT:
			p.Current.D = UP
		case LEFT:
			p.Current.D = DOWN
		case UP:
			p.Current.D = LEFT
		case DOWN:
			p.Current.D = RIGHT
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

func mod(a, b int) int {
	return (a%b + b) % b
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
	data.Grid = make(map[Coord]*Sticker)
	data.Actions = make([]Action, 0)
	actionRegex := regexp.MustCompile(`(\d+)?([LR])(\d+)`)
	for i, l := range strings.Split(strData, "\n") {
		log.Debug(fmt.Sprintf("Parsing line %03d: %s", i, l))
		// Last line
		if len(l) > 0 && (l[0] == 'L' || l[0] == 'R' || (l[0] >= '0' && l[0] <= '9')) {
			matches := actionRegex.FindAllSubmatch([]byte(l), -1)
			log.Debug(fmt.Sprintf("Parsing actions: %s", matches))
			for _, m := range matches {
				for j := 1; j < len(m); j++ {
					if len(m[j]) == 0 {
						continue
					}
					v, err := strconv.Atoi(string(m[j]))
					a := Action{}
					if err == nil {
						a.Steps = v
						a.D = NULL
					} else if rune(m[j][0]) == 'L' {
						a.D = LEFT
					} else if rune(m[j][0]) == 'R' {
						a.D = RIGHT
					}
					data.Actions = append(data.Actions, a)
				}
			}
		} else if len(l) > 0 {
			for j, c := range l {
				if c == WALL || c == EMPTY {
					if i == 0 && data.Current.P.J == 0 {
						data.Current = Position{
							Coord{i, j},
							RIGHT,
						}
					}
					data.Grid[Coord{i, j}] = &Sticker{Value: c}
				}
			}
		}
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
	problem.Fold(context.Int("size"))
	log.Debug(fmt.Sprintf("Folded problem %#v", problem.Grid))
	problem.Traverse()
	result = (problem.Current.P.I+1)*1000 + (problem.Current.P.J+1)*4 + int(problem.Current.D)

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
				Name:    "size",
				Aliases: []string{"s"},
				Value:   50,
				Usage:   "Map face size",
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
