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

	start := time.Now()
	if *part == 1 {
		part1Handler()
	} else {
		part2Handler()
	}
	runtime := time.Since(start)

	fmt.Println("took", runtime)
}

func SolveChosenPart(handler func(bool) int64) {
	var ans int64
	RunChosenPart(func() { ans = handler(false) }, func() { ans = handler(true) })
	fmt.Println("answer:", ans)
}
