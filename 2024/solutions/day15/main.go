package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Vec2 = util.Vec2[int]

type Tile rune

const (
	Empty    Tile = '.'
	Wall          = '#'
	Robot         = '@'
	Box           = '['
	BoxRight      = ']'
)

var runeToTile = map[rune]Tile{
	'#': Wall,
	'O': Box,
	'@': Robot,
	'.': Empty,
}

var runeToDirection = map[rune]Vec2{
	'<': {0, -1},
	'>': {0, 1},
	'v': {1, 0},
	'^': {-1, 0},
}

type PuzzleInput struct {
	grid     util.Grid[Tile]
	robotPos Vec2
	moves    []rune
}

func Parse(input string, isPart2 bool) PuzzleInput {
	var problem PuzzleInput
	parts := strings.Split(input, "\n\n")

	if isPart2 {
		scanner := bufio.NewScanner(strings.NewReader(parts[0]))
		for scanner.Scan() {
			line := scanner.Text()
			row := make([]Tile, 2*len(line))
			for i, r := range line {
				row[i*2] = runeToTile[r]
				switch runeToTile[r] {
				case Robot:
					problem.robotPos = Vec2{len(problem.grid), i * 2}
					row[i*2+1] = Empty
				case Box:
					row[i*2+1] = BoxRight
				default:
					row[i*2+1] = row[i*2]
				}
			}
			problem.grid = append(problem.grid, row)
		}
	} else {
		problem.grid = util.NewGridFromString(parts[0], func(r rune, pos Vec2) Tile {
			if runeToTile[r] == Robot {
				problem.robotPos = pos
			}
			return runeToTile[r]
		})
	}

	scan := bufio.NewScanner(strings.NewReader(parts[1]))
	for scan.Scan() {
		line := scan.Text()
		problem.moves = append(problem.moves, []rune(line)...)
	}

	return problem
}

func solve(p PuzzleInput, isPart2 bool) int64 {
	simulate(&p)

	ans := int64(0)
	for r, row := range p.grid {
		for c, tile := range row {
			if tile == Box {
				ans += int64(100*r + c)
			}
		}
	}
	return ans
}

func simulate(p *PuzzleInput) {
	for _, moveRune := range p.moves {
		//fmt.Printf("move #%d: %c\n", i, moveRune)
		move := runeToDirection[moveRune]
		gridCopy := p.grid.Copy()
		if tryMove(&gridCopy, p.robotPos, move, false) {
			p.grid = gridCopy
			p.robotPos = p.robotPos.Add(move)
		}
		//fmt.Println()
	}
	printGrid(p.grid)
}

func printGrid(grid util.Grid[Tile]) {
	for _, row := range grid {
		for _, tile := range row {
			fmt.Printf("%c", rune(tile))
		}
		fmt.Println()
	}
}

// returns whether the given tile was successfully moved out of the way,
// updating the passed-in grid
// if first val is false, the grid should be considered invalid
func tryMove(grid *util.Grid[Tile], pos Vec2, dir Vec2, moveOnlySingle bool) bool {
	if grid.Get(pos) == Wall {
		return false
	}

	if grid.Get(pos) == Empty {
		return true
	}

	// check if we are moving part of a double-wide box, before moving anything
	tile := grid.Get(pos)
	isDoubleBoxLeft := !moveOnlySingle && (tile == Box) && (grid.Get(pos.Add(Vec2{0, 1})) == BoxRight)
	isDoubleBoxRight := !moveOnlySingle && (tile == BoxRight)

	// set the original spot to empty before recursing to other side of double box, since
	// the other side could end up moving into the original spot and we don't want to override it
	// e.g. if we call tryMove on the left side of a box while moving left
	grid.Set(pos, Empty)

	success := true
	if isDoubleBoxLeft {
		success = tryMove(grid, pos.Add(Vec2{0, 1}), dir, true)
	} else if isDoubleBoxRight {
		success = tryMove(grid, pos.Add(Vec2{0, -1}), dir, true)
	}

	success = success && tryMove(grid, pos.Add(dir), dir, false)
	grid.Set(pos.Add(dir), tile)

	return success
}

//go:embed input
var input string

func main() {
	util.SolveChosenPart(func(isPart2 bool) int64 {
		problem := Parse(input, isPart2)
		printGrid(problem.grid)
		return solve(problem, isPart2)
	})
}
