package main

import (
	"os"
	"strconv"
)

func main() {

	msg := getMessages()

	if len(os.Args) == 1 {
		os.Stdout.WriteString(msg.Help)
		os.Exit(0)
	}

	if len(os.Args) != 4 {
		os.Stderr.WriteString(msg.InvalidArgs)
		os.Exit(1)
	}

	beats, err := strconv.Atoi(os.Args[1])
	if err != nil || beats < 0 {
		os.Stderr.WriteString(msg.InvalidBeats)
		os.Exit(1)
	}

	steps, err := strconv.Atoi(os.Args[2])
	if err != nil || steps < 0 || steps < beats {
		os.Stderr.WriteString(msg.InvalidSteps)
		os.Exit(1)
	}

	shift, err := strconv.Atoi(os.Args[3])
	if err != nil || shift < 0 || shift > steps {
		os.Stderr.WriteString(msg.InvalidShift)
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
