package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

type Cache struct {
	Values    map[string]string
	Size      int
	Generator func(string) string
	LRU       []string
}

func (c *Cache) Get(k string) string {
	if _, ok := c.Values[k]; !ok {
		c.Values[k] = c.Generator(k)
		c.UpdateLRU(k)
		if len(c.Values) > c.Size {
			x, LRU := c.LRU[0], c.LRU[1:]
			c.LRU = LRU
			delete(c.Values, x)
		}
	}
	return c.Values[k]
}

func (c *Cache) UpdateLRU(k string) {
	for i, v := range c.LRU {
		if v == k {
			c.LRU = append(c.LRU[:i], c.LRU[i+1:]...)
		}
	}
	c.LRU = append(c.LRU, k)
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

func MD5(i string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(i)))
}

func repeated(s string, count int, c rune) (r rune, result bool) {
	size := len(s)
	for i := range s {
		if !unicode.IsPrint(c) || (unicode.IsPrint(c) && rune(s[i]) == c) {
			r = rune(s[i])
			if (i + count) <= size {
				result = true
				for j := i; j < i+count; j++ {
					if s[i] != s[j] {
						result = false
						break
					}
				}
				if result {
					return
				}
			}
		}
	}
	return
}

func solution(context *cli.Context) (result int) {
	input := context.String("input")
	cache := Cache{
		Values:    make(map[string]string),
		Size:      1000000,
		Generator: MD5,
	}
	keys := 0
	for i := 0; ; i++ {
		v1 := cache.Get(fmt.Sprintf("%s%d", input, i))
		r1, ok := repeated(v1, 3, '\n')
		if ok {
			log.Debug(fmt.Sprintf("Found 3 repetitions on index %d [%s]", i, v1))
			for j := i + 1; j <= i+1000; j++ {
				v2 := cache.Get(fmt.Sprintf("%s%d", input, j))
				if _, ok := repeated(v2, 5, r1); ok {
					log.Debug(fmt.Sprintf("Found 5 repetitions on index %d [%s]", j, v2))
					keys++
					log.Info(fmt.Sprintf("Found new key on index %d: %d", i, keys))
					break
				}
			}
		}
		if keys == 64 {
			result = i
			break
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
				Value:   "jlmsuwbz",
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
