package main

import (
	"bufio"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Vec2 = util.Vec2[int]

type Tile int

const (
	Empty Tile = iota
	Wall
	Box
	Robot
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

func Parse(input string) PuzzleInput {
	var problem PuzzleInput
	parts := strings.Split(input, "\n\n")

	problem.grid = util.NewGridFromString(parts[0], func(r rune, pos Vec2) Tile {
		if runeToTile[r] == Robot {
			problem.robotPos = pos
		}
		return runeToTile[r]
	})
	
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
		move := runeToDirection[moveRune]
		//fmt.Println("move", move)
		cursor := p.robotPos

		// scan in the direction until we either hit a wall or empty space
		for p.grid.Get(cursor) != Wall && p.grid.Get(cursor) != Empty {
			cursor = cursor.Add(move)
		}

		if p.grid.Get(cursor) == Wall {
			// all spots between our current position and the nearest wall have boxes, move is blocked
			continue
		}

		// move in the desired direction and push any boxes in the way back
		p.grid.Set(p.robotPos, Empty)
		p.robotPos = p.robotPos.Add(move)
		for cursor != p.robotPos {
			p.grid.Set(cursor, p.grid.Get(cursor.Sub(move)))
			cursor = cursor.Sub(move)
		}
		p.grid.Set(p.robotPos, Robot)

		//fmt.Println(p)
	}
}

//go:embed input
var input string

func main() {
	util.SolveChosenPart(func(isPart2 bool) int64 {
		problem := Parse(input)
		return solve(problem, isPart2)
	})
}
