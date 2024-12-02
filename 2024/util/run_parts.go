package util

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func divider() {
	fmt.Println(strings.Repeat("-", 40))
}

func RunChosenPart(part1Handler func(), part2Handler func()) {
	part := flag.Int("part", 1, "either 1 or 2")
	flag.Parse()

	fmt.Println("running part", *part)
	divider()

	start := time.Now()
	if *part == 1 {
		part1Handler()
	} else {
		part2Handler()
	}
	runtime := time.Since(start)

	divider()
	fmt.Println("took", runtime)
}
