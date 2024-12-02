package util

import (
	"flag"
	"fmt"
)

func RunChosenPart(part1Handler func(), part2Handler func()) {
	part := flag.Int("part", 1, "either 1 (default) or 2")
	flag.Parse()

	fmt.Println("running part", *part)
	if *part == 1 {
		part1Handler()
	} else {
		part2Handler()
	}
}
