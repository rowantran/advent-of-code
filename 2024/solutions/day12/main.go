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

// return directions of adjacent tiles with same value, i.e. adjacent tiles in the same region
func (p PuzzleInput) SameValuedNeighborDirections(pos Vec2) []Vec2 {
	var neighborDirs []Vec2
	for _, dir := range directions {
		neighbor := pos.Add(dir)
		if p.InBounds(neighbor) && p.Get(pos) == p.Get(neighbor) {
			neighborDirs = append(neighborDirs, dir)
		}
	}
	return neighborDirs
}

// count corners contributed by the tile to its respective region, using the following scheme:
// * external corners (where one tile has 2 exposed borders) are counted on that tile
// * internal corners (where 3 tiles form an L shape) are counted at the connecting tile
func (p PuzzleInput) Corners(pos Vec2) int {
	corners := 0

	// count external corners
	neighborDirs := p.SameValuedNeighborDirections(pos)
	switch len(neighborDirs) {
	case 0:
		// special case where we just have a lone tile
		corners += 4
	case 1:
		// "peninsula" tile
		corners += 2
	case 2:
		// need to ensure that the neighbors are diagonal from each other and not collinear,
		// i.e. this tile is actually at a corner and not just along a side
		if neighborDirs[0].IsOrthogonal(neighborDirs[1]) {
			corners += 1
		}
	}

	// count internal corners
	for i := range len(neighborDirs) {
		for j := 0; j < i; j++ {
			diagonal := pos.Add(neighborDirs[i]).Add(neighborDirs[j])
			if neighborDirs[i].IsOrthogonal(neighborDirs[j]) && p.Get(pos) != p.Get(diagonal) {
				corners++
			}
		}
	}

	return corners
}

func solve(p PuzzleInput, part2 bool) int {
	visited := make([][]bool, len(p.grid))
	for i := range len(p.grid) {
		visited[i] = make([]bool, len(p.grid[i]))
	}

	result := 0
	for i := range len(p.grid) {
		for j := range len(p.grid[i]) {
			if !visited[i][j] {
				area, perimeter, corners := dfs(p, visited, Vec2{i, j})
				if part2 {
					result += area * corners
				} else {
					result += area * perimeter
				}
			}
		}
	}
	return result
}

// perform a DFS from the given location, returning (area, perimeter, corners) pair of the unvisited portion of the contained region
func dfs(p PuzzleInput, visited [][]bool, pos Vec2) (int, int, int) {
	r, c := pos.Parts()
	visited[r][c] = true

	neighborDirs := p.SameValuedNeighborDirections(pos)
	// perimeter contributed by pos = number of adjacent tiles that aren't neighbors
	area, perimeter, corners := 1, len(directions)-len(neighborDirs), p.Corners(pos)
	for _, dir := range neighborDirs {
		neighbor := pos.Add(dir)
		nr, nc := neighbor.Parts()
		if !visited[nr][nc] {
			subArea, subPerimeter, subCorners := dfs(p, visited, neighbor)
			area += subArea
			perimeter += subPerimeter
			corners += subCorners
		}
	}
	return area, perimeter, corners
}

//go:embed input
var input string

func part1() {
	problem := Parse(input)
	ans := solve(problem, false)
	//fmt.Println(problem)
	fmt.Println("answer:", ans)
}

func part2() {
	problem := Parse(input)
	ans := solve(problem, true)
	fmt.Println("answer:", ans)
}

func main() {
	util.RunChosenPart(part1, part2)
}
