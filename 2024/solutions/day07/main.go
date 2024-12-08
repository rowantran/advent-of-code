package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

//go:embed input
var input string

type PuzzleInput struct {
	equations []Equation	
}

func Parse(input string) PuzzleInput {
	var problem PuzzleInput
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		var equation Equation

		for i, num := range strings.Fields(line) {
			if i == 0 {
				// strip the colon from the end of the target number, then parse as int64
				parsedNum, err := strconv.ParseInt(num[:len(num)-1], 10, 64)
				if err != nil {
					panic(err)
				}
				equation.target = parsedNum
			} else {
				equation.vals = append(equation.vals, int64(util.MustAtoi(num)))
			}
		}

		problem.equations = append(problem.equations, equation)
	}
	return problem
}

type Equation struct {
	target int64
	vals []int64
}

func (e Equation) isSatisfiable(allowConcatenation bool) bool {
	return e.isSatisfiableHelper(allowConcatenation, e.vals[0], 1)
}

func (e Equation) isSatisfiableHelper(allowConcatenation bool, partialResult int64, nextIndex int) bool {
	// key observation: all vals are positive and only *, + operations are allowed, so
	// our partial sum can only increase as we use more values
	if partialResult == e.target {
		return true
	} else if partialResult > e.target || nextIndex >= len(e.vals) {
		return false
	}

	satisfiable := e.isSatisfiableHelper(allowConcatenation, partialResult + e.vals[nextIndex], nextIndex+1) ||
	               e.isSatisfiableHelper(allowConcatenation, partialResult * e.vals[nextIndex], nextIndex+1)
	if (allowConcatenation) {
		concatenated := concatenate(partialResult, e.vals[nextIndex])
		satisfiable = satisfiable || e.isSatisfiableHelper(allowConcatenation, concatenated, nextIndex+1)
	}
	return satisfiable
}

func digits(a int64) int {
	var p int = 0
	var base int64 = 1

	for base <= a {
		p++
		base *= 10
	}

	// p is the smallest int s.t. 10^p > a
	return p
}

// concatenate(123, 456) = 123456
func concatenate(a int64, b int64) int64 {
	res := a
	for range digits(b) {
		res *= 10
	}

	return res+b
}

func part1() {
	problem := Parse(input)
	var ans int64
	for _, eqn := range problem.equations {
		//fmt.Println("checking equation", eqn)
		if eqn.isSatisfiable(false) {
			//fmt.Println("satisfied")
			ans += eqn.target
		}
	}

	fmt.Println("answer", ans)
}

func part2() {
	problem := Parse(input)
	var ans int64
	for _, eqn := range problem.equations {
		//fmt.Println("checking equation", eqn)
		if eqn.isSatisfiable(true) {
			//fmt.Println("satisfied")
			ans += eqn.target
		}
	}

	fmt.Println("answer", ans)
}

func main() {
	util.RunChosenPart(part1, part2)
}
