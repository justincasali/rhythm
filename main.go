package main

import (
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 4 {
		os.Stderr.WriteString("invalid number of arguments\n")
		os.Exit(1)
	}

	hits, err := strconv.Atoi(os.Args[1])
	if err != nil || hits < 0 {
		os.Stderr.WriteString("invalid hits\n")
		os.Exit(1)
	}

	steps, err := strconv.Atoi(os.Args[2])
	if err != nil || steps < 0 || steps < hits {
		os.Stderr.WriteString("invalid steps\n")
		os.Exit(1)
	}

	shift, err := strconv.Atoi(os.Args[3])
	if err != nil || shift < 0 || shift > steps {
		os.Stderr.WriteString("invalid shift\n")
		os.Exit(1)
	}

	pattern := rhythm(hits, steps-hits)

	if pattern != nil {
		pattern = pattern.Move(-shift)
	}

	pattern.Do(func(value interface{}) {
		if value.(bool) {
			os.Stdout.WriteString("x")
		} else {
			os.Stdout.WriteString(".")
		}
	})
	os.Stdout.WriteString("\n")

}
