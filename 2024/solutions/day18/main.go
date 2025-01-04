package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Vec2 = util.Vec2[int]

var dirs = []Vec2{{1,0}, {-1,0}, {0,1}, {0,-1}}

type PuzzleInput struct {
	bytes []Vec2
}

// generate the grid by marking byte locations [0, end) as true
func (p PuzzleInput) GenerateGrid(end int) util.Grid[bool] {
	var grid util.Grid[bool]
	for range 71 {
		grid = append(grid, make([]bool, 71))
	}

	for i := 0; i < end; i++ {
		x, y := p.bytes[i].Parts()
		grid[x][y] = true
	}

	return grid
}

func Parse(input string, isPart2 bool) PuzzleInput {
	var p PuzzleInput

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		x, y := util.MustAtoi(coords[0]), util.MustAtoi(coords[1])
		p.bytes = append(p.bytes, Vec2{x,y})
	}

	return p
}

type BfsQueueEntry struct {
	node Vec2
	depth int
}

func bfs(grid util.Grid[bool]) int {
	visited := make(map[Vec2]bool)
	queue := []BfsQueueEntry{{Vec2{0,0},0}}
	for len(queue) > 0 {
		h := queue[0]
		queue = queue[1:]

		if (h.node == Vec2{len(grid)-1, len(grid[0])-1}) {
			return h.depth
		}

		if visited[h.node] {
			continue
		}

		visited[h.node] = true
		for _, dir := range dirs {
			neighbor := h.node.Add(dir)
			if grid.InBounds(neighbor) && !grid.Get(neighbor) {
				queue = append(queue, BfsQueueEntry{neighbor, h.depth+1})
			}
		}
	}

	return -1
}

func findFirstBlockingByte(p PuzzleInput) Vec2 {
	// search for the first value where we cannot find an exit using bytes [0, n)
	// that means blocker n-1 is the first blocking byte
	lo, hi := 0, len(p.bytes)
	for lo < hi {
		mid := (lo+hi)/2
		success := bfs(p.GenerateGrid(mid)) != -1
		if success {
			lo = mid+1
		} else {
			hi = mid
		}
	}
	
	return p.bytes[lo-1]
}

func solve(p PuzzleInput, isPart2 bool) int64 {
	if !isPart2 {
		grid := p.GenerateGrid(1024)
		return int64(bfs(grid))
	} else {
		fmt.Println(findFirstBlockingByte(p))
		return 0
	}
}

//go:embed input
var input string

func main() {
	util.SolveChosenPart(func(isPart2 bool) int64 {
		problem := Parse(input, isPart2)
		return solve(problem, isPart2)
	})
}
