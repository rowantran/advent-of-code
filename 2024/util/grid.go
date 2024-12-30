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

func (g Grid[T]) Set(pos Vec2[int], val T) {
	r, c := pos.Parts()
	g[r][c] = val
}

func (g Grid[T]) InBounds(pos Vec2[int]) bool {
	r, c := pos.Parts()
	return r >= 0 && r < len(g) && c >= 0 && c < len(g[0])
}

func (g Grid[T]) Copy() Grid[T] {
	newGrid := make(Grid[T], len(g))
	for i, row := range g {
		newGrid[i] = make([]T, len(row))
		copy(newGrid[i], row)
	}
	return newGrid
}
