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

	beats, err := strconv.Atoi(os.Args[1])
	if err != nil || beats < 0 {
		os.Stderr.WriteString("invalid beats\n")
		os.Exit(1)
	}

	steps, err := strconv.Atoi(os.Args[2])
	if err != nil || steps < 0 || steps < beats {
		os.Stderr.WriteString("invalid steps\n")
		os.Exit(1)
	}

	shift, err := strconv.Atoi(os.Args[3])
	if err != nil || shift < 0 || shift > steps {
		os.Stderr.WriteString("invalid shift\n")
		os.Exit(1)
	}

	chain := rhythmFast(beats, steps-beats)

	os.Stdout.WriteString("[")

	if len(chain) == 0 {
		os.Stdout.WriteString(" ")
	} else {
		n := len(chain)
		for i := 0; i < n; i++ {
			if chain[(i+shift)%n] {
				os.Stdout.WriteString("x ")
			} else {
				os.Stdout.WriteString(". ")
			}
		}
	}

	os.Stdout.WriteString("\b]\n")

}
