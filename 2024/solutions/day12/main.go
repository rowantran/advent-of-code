package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Vec2 = util.Vec2

var directions = []Vec2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type PuzzleInput struct {
	grid [][]rune
}

func Parse(input string) PuzzleInput {
	var problem PuzzleInput
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		problem.grid = append(problem.grid, []rune(scanner.Text()))
	}
	return problem
}

func (p PuzzleInput) Get(pos Vec2) rune {
	r, c := pos.Parts()
	return p.grid[r][c]
}

func (p PuzzleInput) InBounds(pos Vec2) bool {
	r, c := pos.Parts()
	return r >= 0 && r < len(p.grid) && c >= 0 && c < len(p.grid[0])
}

// return adjacent tiles with same value, i.e. adjacent tiles in the same region
func (p PuzzleInput) SameValuedNeighbors(pos Vec2) []Vec2 {
	var neighbors []Vec2
	for _, dir := range directions {
		neighbor := pos.Add(dir)
		if p.InBounds(neighbor) && p.Get(pos) == p.Get(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func solve(p PuzzleInput) int {
	visited := make([][]bool, len(p.grid))
	for i := range len(p.grid) {
		visited[i] = make([]bool, len(p.grid[i]))
	}
	
	result := 0
	for i := range len(p.grid) {
		for j := range len(p.grid[i]) {
			if !visited[i][j] {
				area, perimeter := dfs(p, visited, Vec2{i, j})
				result += area * perimeter
			}
		}
	}
	return result
}

// perform a DFS from the given location, returning (area, perimeter) pair of the unvisited portion of the contained region
func dfs(p PuzzleInput, visited [][]bool, pos Vec2) (int, int) {
	r, c := pos.Parts()
	visited[r][c] = true

	neighbors := p.SameValuedNeighbors(pos)
	// perimeter contributed by pos = number of adjacent tiles that aren't neighbors
	area, perimeter := 1, len(directions) - len(neighbors)
	for _, neighbor := range neighbors {
		nr, nc := neighbor.Parts()
		if !visited[nr][nc] {
			subArea, subPerimeter := dfs(p, visited, neighbor)
			area += subArea
			perimeter += subPerimeter
		}
	}
	return area, perimeter
}

//go:embed input
var input string

func part1() {
	problem := Parse(input)
	ans := solve(problem)
	//fmt.Println(problem)
	fmt.Println("answer:", ans)
}

func part2() {
	problem := Parse(input)
	ans := solve(problem)
	fmt.Println("answer:", ans)
}

func main() {
	util.RunChosenPart(part1, part2)
}
