package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

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

func reverse(s []byte) (result []byte) {
	result = make([]byte, len(s))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = s[j], s[i]
	}
	result[len(s)/2] = s[len(s)/2]
	return
}

func flip(s []byte) (result []byte) {
	for _, c := range s {
		var r byte
		if c == '1' {
			r = '0'
		} else {
			r = '1'
		}
		result = append(result, r)
	}
	return
}

func expand(a []byte) (result []byte) {
	result = append(a, '0')
	for i, j := 0, len(a)-1; j >= 0; i, j = i+1, j-1 {
		if a[j] == '0' {
			result = append(result, '1')
		} else {
			result = append(result, '0')
		}
	}
	log.Debug(fmt.Sprintf("%s", result, a))
	return
}

func checksum(s []byte) (result []byte) {
	for i := 0; i < len(s); i += 2 {
		if s[i] == s[i+1] {
			result = append(result, '1')
		} else {
			result = append(result, '0')
		}
	}
	log.Debug(fmt.Sprintf("Computed checksum of %s (%d) %s", s, len(result), result))
	if (len(result) % 2) == 0 {
		result = checksum(result)
	}
	return
}

func solution(context *cli.Context) (result []byte) {
	var input = context.String("input")
	var limit = context.Int("limit")

	str := []byte(input)
	for i := 0; len(str) < limit; i++ {
		str = expand(str)
		// checksum(str[:len(str)-1])
	}
	log.Info(fmt.Sprintf("Generated (%d) %s", len(str), str))
	str = str[:limit]
	result = checksum(str)
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
				Value:   "00111101111101000",
				Usage:   "Input value",
			},
			&cli.IntFlag{
				Name:  "limit",
				Value: 272,
				Usage: "Disk size limit",
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
			echo(fmt.Sprintf("Solution is %s", solution(c)), c.String("output"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
