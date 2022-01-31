package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 4 {
		panic("invalid number of arguments")
	}

	hits, err := strconv.Atoi(os.Args[1])
	if err != nil || hits < 0 {
		panic("invalid hits")
	}

	steps, err := strconv.Atoi(os.Args[2])
	if err != nil || steps < 0 || steps < hits {
		panic("invalid steps")
	}

	phase, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic("invalid phase")
	}

	pattern := rhythm(hits, steps-hits).Move(phase)

	pattern.Do(func(value interface{}) {
		if value.(bool) {
			fmt.Print("x")
		} else {
			fmt.Print(".")
		}
	})
	fmt.Println()

}
