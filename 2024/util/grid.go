package util

import (
	"bufio"
	"strings"
)

type Grid [][]rune

func NewGridFromString(input string) Grid {
	var grid Grid
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid
}

func (g Grid) Get(pos Vec2[int]) rune {
	r, c := pos.Parts()
	return g[r][c]
}

func (g Grid) InBounds(pos Vec2[int]) bool {
	r, c := pos.Parts()
	return r >= 0 && r < len(g) && c >= 0 && c < len(g[0])
}
