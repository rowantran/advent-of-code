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

// returns: (set of positions visited), (if the walker looped)
func walk(problem PuzzleInput) (map[[2]int]util.Set[complex64], bool) {
	visited := make(map[[2]int]util.Set[complex64])

	// represent orientation as complex number w/ real part in row-axis and imaginary part in col-axis
	// i.e. "up" -> reduce row by 1 -> -1
	pos := problem.startPos
	var direction complex64 = -1
	for problem.InBounds(pos) && !visited[pos].Has(direction) {
		if _, ok := visited[pos]; !ok {
			visited[pos] = make(util.Set[complex64])
		}
		visited[pos].Add(direction)

		nextPos := move(pos, direction)
		if problem.InBounds(nextPos) && problem.Get(nextPos) == Obstacle {
			direction *= -1i
		} else {
			pos = nextPos
		}

	}

	return visited, problem.InBounds(pos)
}

func part1() {
	problem := Parse(input)
	path, _ := walk(problem)
	fmt.Println("answer:", len(path))
}

func part2() {
	problem := Parse(input)
	path, _ := walk(problem)

	count := 0
	for pos := range path {
		// check if rotating at this point would have resulted in a loop
		if problem.Get(pos) == StartPos {
			continue
		}
		problem.Set(pos, Obstacle)
		if _, wouldLoop := walk(problem); wouldLoop {
			count++
		}
		problem.Set(pos, Empty)
	}

	fmt.Println("answer:", count)
}

func main() {
	util.RunChosenPart(part1, part2)
}
