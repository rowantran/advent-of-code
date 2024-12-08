package main

import (
	"bufio"
	"fmt"
	"math/big"
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
				equation.target = big.NewInt(parsedNum)
			} else {
				equation.vals = append(equation.vals, int64(util.MustAtoi(num)))
			}
		}

		problem.equations = append(problem.equations, equation)
	}
	return problem
}

type Equation struct {
	target *big.Int
	vals []int64
}

func (e Equation) isSatisfiable(allowConcatenation bool) bool {
	return e.isSatisfiableHelper(allowConcatenation, big.NewInt(e.vals[0]), 1)
}

func (e Equation) isSatisfiableHelper(allowConcatenation bool, partialResult *big.Int, nextIndex int) bool {
	// key observation: all vals are positive and only *, + operations are allowed, so
	// our partial sum can only increase as we use more values
	if partialResult.Cmp(e.target) == 0 {
		return true
	} else if partialResult.Cmp(e.target) == 1 || nextIndex >= len(e.vals) {
		return false
	}

	sum := new(big.Int).Add(partialResult, big.NewInt(e.vals[nextIndex]))
	product := new(big.Int).Mul(partialResult, big.NewInt(e.vals[nextIndex]))
	satisfiable := e.isSatisfiableHelper(allowConcatenation, sum, nextIndex+1) ||
	               e.isSatisfiableHelper(allowConcatenation, product, nextIndex+1)
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
func concatenate(a *big.Int, b int64) *big.Int {
	ten := big.NewInt(10)
	res := new(big.Int).Set(a)
	for range digits(b) {
		res.Mul(res, ten)
	}

	return res.Add(res, big.NewInt(b))
}

func part1() {
	problem := Parse(input)
	var ans int64
	for _, eqn := range problem.equations {
		//fmt.Println("checking equation", eqn)
		if eqn.isSatisfiable(false) {
			//fmt.Println("satisfied")
			ans += eqn.target.Int64()
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
			ans += eqn.target.Int64()
		}
	}

	fmt.Println("answer", ans)
}

func main() {
	util.RunChosenPart(part1, part2)
}
