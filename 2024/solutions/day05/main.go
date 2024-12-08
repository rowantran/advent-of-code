package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

//go:embed input
var input string

type PuzzleInput struct {
	rules   [][2]int
	updates [][]int
}

func parse(input string) PuzzleInput {
	var rules [][2]int
	var updates [][]int

	doneRules := false
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			doneRules = true
		} else {
			if !doneRules {
				parts := strings.Split(line, "|")
				rules = append(rules, [2]int{util.MustAtoi(parts[0]), util.MustAtoi(parts[1])})
			} else {
				parts := strings.Split(line, ",")
				update := make([]int, len(parts))
				for i := range parts {
					update[i] = util.MustAtoi(parts[i])
				}
				updates = append(updates, update)
			}
		}
	}

	return PuzzleInput{rules, updates}
}

/*
example: given rule 123|456, the map will contain 456 -> [123] meaning that
we should not ever see page 123 after page 456

i.e. maps each number to numbers that it must appear after
*/
func (p PuzzleInput) predecessors() map[int]util.Set[int] {
	res := make(map[int]util.Set[int])
	for _, rule := range p.rules {
		x, y := rule[0], rule[1]
		if _, ok := res[y]; !ok {
			res[y] = util.CreateSet[int]()
		}
		res[y].Add(x)
	}
	return res
}

func isCorrectlyOrdered(update []int, invalidSuccessors map[int]util.Set[int]) bool {
	// set of values that are illegal to see at a given point
	illegal := util.CreateSet[int]()
	for _, val := range update {
		if illegal.Has(val) {
			return false
		}
		for n := range invalidSuccessors[val] {
			illegal.Add(n)
		}
	}
	return true
}

func part1() {
	problem := parse(input)
	invalid := problem.predecessors()

	total := 0
	for _, update := range problem.updates {
		if isCorrectlyOrdered(update, invalid) {
			//fmt.Println("correctly ordered:", update)
			total += update[len(update)/2]
		}
	}
	fmt.Println("total:", total)
}

func part2() {
	problem := parse(input)
	predecessors := problem.predecessors()

	total := 0
	for _, update := range problem.updates {
		sortedUpdate := slices.SortedFunc(slices.Values(update), func(a, b int) int {
			if predecessors[a].Has(b) {
				return 1
			} else if predecessors[b].Has(a) {
				return -1
			} else {
				return 0
			}
		})
		if !slices.Equal(update, sortedUpdate) {
			total += sortedUpdate[len(sortedUpdate)/2]
		}
	}
	fmt.Println("total:", total)
}

func main() {
	util.RunChosenPart(part1, part2)
}
