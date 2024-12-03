package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

//go:embed input
var input string

func count[T any](list []T, pred func(T) bool) int {
	count := 0
	for _, val := range list {
		if pred(val) {
			count += 1
		}
	}
	return count
}

func parse(input string) [][]int {
	reports := [][]int{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		report := make([]int, len(fields))
		for i, val := range fields {
			report[i] = util.MustAtoi(val)
		}
		reports = append(reports, report)
	}

	return reports
}

func isSafe(report []int) bool {
	prev, curr := 0, 0
	for i, val := range report {
		// report is safe iff:
		// * magnitude of all deltas satisfies 1 <= m <= 3
		// * deltas are same sign <-> product of any two adjacent deltas is > 0
		// first rule should ignore the 1st value since the delta is not defined,
		// similarly the second rule should ignore the first 2 values since we need at least 3
		// values to evaluate a change in delta direction
		magnitude := util.Abs(val - curr)
		badMagnitude := i > 0 && !(1 <= magnitude && magnitude <= 3)
		badDirection := i > 1 && ((curr-prev)*(val-curr) <= 0)
		if badMagnitude || badDirection {
			//fmt.Printf("determined unsafe with prev=%d, curr=%d, val=%d\n", prev, curr, val)
			return false
		}
		prev, curr = curr, val
	}
	return true
}

func isSafeWithDampener(report []int) bool {
	// its 2am so i'm just gonna brute force this, was thinking of a proper O(n) algorithm but
	// didn't have time to complete
	// TODO: think of a better algorithm
	// observations:
	// * if the sequence is already monotonic and unsafe using pt1 rules, it will still be
	//   unsafe with dampener since removing a value can only increase the abs. delta
	//   * UNLESS the invalid magnitude is at one of the boundaries (e.g. 0, 10, 11, ... can be
	//     made safe by removing the 0 at the beginning
	// * if a 3-value sequence forms a "V" shape, we may be able to fix the sequence by removing one
	//   of the values, but we can't always remove the same one (e.g. always remove the 3rd value)
	//   * example: 8 6 3 5 4 forms V shape with (6 3 5) and we need to remove the 3
	//   * example: 8 6 3 7 2 forms V shape with (6 3 7) and we need to remove the 7
	// something should be possible along the lines of choosing the "best" value to keep out of the trio
	if isSafe(report) {
		return true
	} else {
		for i := range report {
			trimmed := make([]int, len(report)-1)
			copy(trimmed, report[:i])
			copy(trimmed[i:], report[i+1:])
			if isSafe(trimmed) {
				return true
			}
		}
		return false
	}
}

func part1() {
	reports := parse(input)
	count := count(reports, isSafe)
	fmt.Println("number of safe reports:", count)
}

func part2() {
	reports := parse(input)
	count := count(reports, isSafeWithDampener)
	fmt.Println("number of safe reports:", count)
}

func main() {
	util.RunChosenPart(part1, part2)
}
