package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Vec2 = util.Vec2
type Set_Vec2 = util.Set[Vec2]

var dirs = []complex64{-1, 1, -1i, 1i}

//go:embed input
var input string

type PuzzleInput struct {
	heights [][]int
}

func Parse(input string) PuzzleInput {
	var problem PuzzleInput
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		lineHeights := make([]int, len(line))
		for i, r := range line {
			lineHeights[i] = util.RuneToInt(r)
		}
		problem.heights = append(problem.heights, lineHeights)
	}
	return problem
}

func (p PuzzleInput) Solve() int {
	rows := len(p.heights)
	cols := len(p.heights[0])
	reachable := make(map[[2]int]Set_Vec2)
	ans := 0

	for r := range rows {
		for c := range cols {
			pos := complex(float32(r), float32(c))
			if p.getHeight(pos) == 0 {
//				fmt.Printf("starting dfs from %d,%d\n", r, c)
				ans += p.dfs(reachable, pos).Size()
			}
		}
	}

	return ans
}

func (p PuzzleInput) dfs(reachable map[[2]int]Set_Vec2, pos complex64) Set_Vec2 {
//	row, col := int(real(pos)), int(imag(pos))
//	fmt.Printf("dfs: %d,%d\n", row, col)
	if val, ok := reachable[vec2(pos)]; ok {
//		fmt.Printf("dfs: %d,%d = %v (cached)\n", row, col, reachable[vec2(pos)])
		return val
	}

	ans := make(Set_Vec2)
	if p.getHeight(pos) == 9 {
		ans.Add(vec2(pos))
	} else {
		for _, dir := range dirs {
			if p.isInGrid(pos+dir) && p.getHeight(pos)+1 == p.getHeight(pos+dir) {
				for peak := range p.dfs(reachable, pos+dir) {
					ans.Add(peak)
				}
			}
		}
	}
	reachable[vec2(pos)] = ans
//	fmt.Printf("dfs: %d,%d = %v\n", row, col, reachable[vec2(pos)])
	return ans
}

func (p PuzzleInput) getHeight(pos complex64) int {
	return p.heights[int(real(pos))][int(imag(pos))]
}

func (p PuzzleInput) isInGrid(pos complex64) bool {
	row, col := int(real(pos)), int(imag(pos))
	return row >= 0 && row < len(p.heights) && col >= 0 && col < len(p.heights[0])
}

func vec2(pos complex64) Vec2 {
	return Vec2{int(real(pos)), int(imag(pos))}
}

func part1() {
	problem := Parse(input)
	ans := problem.Solve()
	//fmt.Println(problem)
	fmt.Println("answer:", ans)
}

func part2() {
	//problem := Parse(input)
	ans := 0
	fmt.Println("answer:", ans)
}

func main() {
	util.RunChosenPart(part1, part2)
}
