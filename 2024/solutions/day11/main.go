package main

import (
	"fmt"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

//go:embed input
var input string

type PuzzleInput struct {
	counts map[int64]int
}

func Parse(input string) PuzzleInput {
	var problem PuzzleInput
	problem.counts = make(map[int64]int)
	for _, num := range strings.Fields(input) {
		parsedNum := util.MustAtoiInt64(num)
		problem.counts[parsedNum]++
	}
	return problem
}

func solve(p PuzzleInput, iterations int) int {
	for range iterations {
		// stone transformations are "simultaneous" so we can't mutate the counts map in-place
		deltas := make(map[int64]int)

		for k, v := range p.counts {
			deltas[k] -= v

			switch {
			case k == 0:
				deltas[1] += v
			case util.DigitCountInt64(k)%2 == 0:
				left, right := splitNum(k)
				deltas[left] += v
				deltas[right] += v
			default:
				deltas[2024*k] += v
			}
		}

		for k, v := range deltas {
			p.counts[k] += v
			if p.counts[k] == 0 {
				delete(p.counts, k)
			}
		}

		//fmt.Printf("completed iteration #%d: %v\n", i+1, p.counts)
	}

	sum := 0
	for _, v := range p.counts {
		sum += v
	}
	return sum
}

// split an integer with even number of digits into left and right halves
func splitNum(n int64) (int64, int64) {
	digits := util.DigitCountInt64(n)
	div := util.ExpInt64(10, digits/2)
	return n / div, n % div
}

func part1() {
	problem := Parse(input)
	//fmt.Println(problem)
	ans := solve(problem, 25)
	fmt.Println("answer:", ans)
}

func part2() {
	problem := Parse(input)
	ans := solve(problem, 75)
	fmt.Println("answer:", ans)
}

func main() {
	util.RunChosenPart(part1, part2)
}
