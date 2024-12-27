package util

import (
	"bufio"
	"strings"
)

type Grid[T any] [][]T

func NewGridFromString[T any](input string, mappingFunc func(r rune, pos Vec2[int]) T) Grid[T] {
	var grid Grid[T]
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]T, len(line))
		for i, r := range line {
			row[i] = mappingFunc(r, Vec2[int]{len(grid), i})
		}
		grid = append(grid, row)
	}
	return grid
}

func (g Grid[T]) Get(pos Vec2[int]) T {
	r, c := pos.Parts()
	return g[r][c]
}

func (g Grid[T]) InBounds(pos Vec2[int]) bool {
	r, c := pos.Parts()
	return r >= 0 && r < len(g) && c >= 0 && c < len(g[0])
}
