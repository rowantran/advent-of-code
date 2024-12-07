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

type PuzzleInput struct {
	grid [][]Tile
	startPos [2]int
}

func (i PuzzleInput) InBounds(pos [2]int) bool {
	return pos[0] >= 0 && pos[0] < len(i.grid) && pos[1] >= 0 && pos[1] < len(i.grid[0])
}

func (i PuzzleInput) Get(pos [2]int) Tile {
	return i.grid[pos[0]][pos[1]]
}

func (i PuzzleInput) Set(pos [2]int, val Tile) {
	i.grid[pos[0]][pos[1]] = val
}

func Parse(input string) PuzzleInput {
	var grid [][]Tile
	var startPos [2]int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		tileLine := make([]Tile, len(line))
		for j := range line {
			tileLine[j] = MatchTile(line[j])
			if tileLine[j] == StartPos {
				startPos = [2]int{i, j}
			}
		}
		grid = append(grid, tileLine)
	}
	return PuzzleInput{grid, startPos}
}

type Tile int
const (
	Empty Tile = iota
	Obstacle
	StartPos
)

func MatchTile(ch byte) Tile {
	switch ch {
	case '#':
		return Obstacle
	case '^':
		return StartPos
	default:
		return Empty
	}
}

func move(pos [2]int, dir complex64) [2]int {
	return [2]int{pos[0] + int(real(dir)), pos[1] + int(imag(dir))}
}

func solve(problem PuzzleInput, p2 bool) (int, bool) {
	ans := 0
	visited := make(map[[2]int]util.Set[complex64])
	loopObstacles := make(util.Set[[2]int])

	pos := problem.startPos
	//fmt.Printf("starting at %v\n", pos)
	// represent orientation as complex number w/ real part in row-axis and imaginary part in col-axis
	// i.e. "up" -> reduce row by 1 -> -1
	var direction complex64 = -1
	for !visited[pos].Has(direction) {
		// check if we're visiting a coordinate for the first time
		if _, ok := visited[pos]; !ok {
			if !p2 {
				ans++
			}
			visited[pos] = make(util.Set[complex64])
		}
		visited[pos].Add(direction)

		// move
		nextPos := move(pos, direction)
		if !problem.InBounds(nextPos) {
			//fmt.Printf("going out of bounds to %v\n", nextPos)
			return ans, false
		}

		if problem.Get(nextPos) == Obstacle {
			// turn right 90 degrees
			direction *= -1i
			//fmt.Printf("rotating to face %v\n", direction)
		} else {
			// check if rotating at this point would have resulted in a loop
			if p2 && problem.Get(nextPos) != StartPos {
				problem.Set(nextPos, Obstacle)

				_, wouldLoop := solve(problem, false)
				if wouldLoop && !loopObstacles.Has(nextPos) {
					loopObstacles.Add(nextPos)
					ans++
				}

				problem.Set(nextPos, Empty)
			}

			pos = nextPos
			//fmt.Printf("moving to %v\n", pos)
		}
	}

	return ans, true
}

func part1() {
	problem := Parse(input)
	ans, _ := solve(problem, false)
	fmt.Println("answer:", ans)
}

func part2() {
	problem := Parse(input)
	ans, _ := solve(problem, true)
	fmt.Println("answer:", ans)
}

func main() {
	util.RunChosenPart(part1, part2)
}
