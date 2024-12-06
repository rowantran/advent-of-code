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

var directions [][]int = buildDirections(false)
var diagonalDirections [][]int = buildDirections(true)

func buildDirections(onlyDiagonals bool) [][]int {
	var directions [][]int
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if (dx == 0 && dy == 0) || (onlyDiagonals && (dx == 0 || dy == 0)) {
				continue
			}
			directions = append(directions, []int{dx, dy})
		}
	}
	return directions
}

func parse(input string) [][]rune {
	var result [][]rune
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, []rune(line))
	}
	return result
}

// returns true if there is a match starting at the given position, going in the given direction
func matchesDirection(grid [][]rune, target string, i int, j int, dir []int) bool {
	di, dj := dir[0], dir[1]
	for d := range target {
		ni := i + d*di
		nj := j + d*dj

		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[0]) || grid[ni][nj] != rune(target[d]) {
			return false
		}
	}
	return true
}

// returns: valid matches starting at the given position, as a list of (i, j, di, dj) pairs
func matches(grid [][]rune, target string, i int, j int, directions [][]int) [][]int {
	var matches [][]int
	for _, dir := range directions {
		if matchesDirection(grid, target, i, j, dir) {
			matches = append(matches, []int{i, j, dir[0], dir[1]})
		}
	}
	return matches
}

func part1() {
	grid := parse(input)
	count := 0

	for i := range grid {
		for j := range grid[i] {
			count += len(matches(grid, "XMAS", i, j, directions))
		}
	}

	fmt.Println("number of occurrences:", count)
}

func part2() {
	grid := parse(input)
	xmasCount := 0

	centerCount := make(map[[2]int]int)
	for i := range grid {
		for j := range grid[i] {
			foundMatches := matches(grid, "MAS", i, j, diagonalDirections)
			for _, match := range foundMatches {
				// location of an A within a "MAS"
				center := [2]int{match[0] + match[2], match[1] + match[3]}
				centerCount[center] += 1
				if centerCount[center] > 1 {
					fmt.Println("XMAS at", center)
					xmasCount += 1
				}
			}
		}
	}

	fmt.Println("number of occurrences:", xmasCount)

}

func main() {
	util.RunChosenPart(part1, part2)
}
