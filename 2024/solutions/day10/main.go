package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

var dirs = []complex64{-1, 1, -1i, 1i}

//go:embed input
var input string

type PointInfo struct {
	peaks util.Set[complex64]
	paths int
}

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

func (p PuzzleInput) Solve(p2 bool) int {
	reachable := make(map[complex64]PointInfo)
	ans := 0

	for r := range len(p.heights) {
		for c := range len(p.heights[0]) {
			pos := util.Pack(r, c)
			if p.getHeight(pos) == 0 {
				if p2 {
					ans += p.dfs(reachable, pos).paths
				} else {
					ans += p.dfs(reachable, pos).peaks.Size()
				}
			}
		}
	}

	return ans
}

func (p PuzzleInput) dfs(reachable map[complex64]PointInfo, pos complex64) PointInfo {
	if val, ok := reachable[pos]; ok {
		return val
	}

	ans := PointInfo{peaks: make(util.Set[complex64])}
	if p.getHeight(pos) == 9 {
		ans.peaks.Add(pos)
		ans.paths = 1
	} else {
		for _, dir := range dirs {
			if p.isInGrid(pos+dir) && p.getHeight(pos)+1 == p.getHeight(pos+dir) {
				res := p.dfs(reachable, pos+dir)
				for peak := range res.peaks {
					ans.peaks.Add(peak)
				}
				ans.paths += res.paths
			}
		}
	}
	reachable[pos] = ans
	return ans
}

func (p PuzzleInput) getHeight(pos complex64) int {
	row, col := util.Unpack(pos)
	return p.heights[row][col]
}

func (p PuzzleInput) isInGrid(pos complex64) bool {
	row, col := util.Unpack(pos)
	return row >= 0 && row < len(p.heights) && col >= 0 && col < len(p.heights[0])
}

func part1() {
	problem := Parse(input)
	ans := problem.Solve(false)
	//fmt.Println(problem)
	fmt.Println("answer:", ans)
}

func part2() {
	problem := Parse(input)
	ans := problem.Solve(true)
	fmt.Println("answer:", ans)
}

func main() {
	util.RunChosenPart(part1, part2)
}
