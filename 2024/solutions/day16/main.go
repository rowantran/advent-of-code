package main

import (
	"container/heap"
	"math"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Vec2 = util.Vec2[int]

var directions = []Vec2{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type PuzzleInput struct {
	maze  util.Grid[rune]
	start Vec2
	end   Vec2
}

func Parse(input string) PuzzleInput {
	var p PuzzleInput
	p.maze = util.NewGridFromString(input, func(r rune, pos Vec2) rune {
		if r == 'S' {
			p.start = pos
		} else if r == 'E' {
			p.end = pos
		}
		return r
	})
	return p
}

func (p PuzzleInput) Neighbors(pos Vec2) []Vec2 {
	neighbors := make([]Vec2, 0)
	for _, dir := range directions {
		n := pos.Add(dir)
		if p.maze.InBounds(n) && p.maze.Get(n) != '#' {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

type PuzzleNode struct {
	location  Vec2
	direction Vec2
}

type PuzzleNodeHeapItem struct {
	node     PuzzleNode
	distance int
}

func solve(p PuzzleInput, isPart2 bool) int64 {
	dists, prevs := dijkstra(p)

	minDistance := math.MaxInt
	endNodes := []PuzzleNode{}
	for _, dir := range directions {
		node := PuzzleNode{p.end, dir}
		if dists[node] < minDistance {
			minDistance = dists[node]
			endNodes = []PuzzleNode{node}
		} else if dists[node] == minDistance {
			endNodes = append(endNodes, node)
		}
	}

	if !isPart2 {
		return int64(minDistance)
	} else {
		tiles := make(util.Set[Vec2])
		for _, node := range endNodes {
			tracePaths(prevs, node, tiles)
		}
		return int64(tiles.Size())
	}
}

// collect all tiles on path(s) to the target node
func tracePaths(prevs map[PuzzleNode][]PuzzleNode, node PuzzleNode, tiles util.Set[Vec2]) {
	tiles.Add(node.location)
	for _, prev := range prevs[node] {
		tracePaths(prevs, prev, tiles)
	}
}

func dijkstra(p PuzzleInput) (map[PuzzleNode]int, map[PuzzleNode][]PuzzleNode) {
	dists := make(map[PuzzleNode]int)
	prevs := make(map[PuzzleNode][]PuzzleNode)
	pq := util.NewHeap(func(a, b PuzzleNodeHeapItem) bool { return a.distance < b.distance })

	for r := range len(p.maze) {
		for c := range len(p.maze[r]) {
			for _, dir := range directions {
				if p.maze[r][c] == '#' {
					continue
				}

				dist := math.MaxInt
				if (p.maze[r][c] == 'S' && dir == Vec2{0, 1}) {
					dist = 0
				}
				node := PuzzleNode{Vec2{r, c}, dir}
				dists[node] = dist
				heap.Push(&pq, PuzzleNodeHeapItem{node, dist})
			}
		}
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(PuzzleNodeHeapItem).node
		neighbors := []PuzzleNode{}

		// build list of neighbors
		forward := node.location.Add(node.direction)
		if p.maze.InBounds(forward) && p.maze.Get(forward) != '#' {
			neighbors = append(neighbors, PuzzleNode{forward, node.direction})
		}
		for _, dir := range directions {
			if node.direction != dir && node.direction != dir.Mul(-1) {
				neighbors = append(neighbors, PuzzleNode{node.location, dir})
			}
		}

		// check neighbors
		for _, neighbor := range neighbors {
			var edgeCost int
			if neighbor.direction == node.direction {
				// forward move
				edgeCost = 1
			} else {
				// 90 degree turn
				edgeCost = 1000
			}
			totalCost := dists[node] + edgeCost
			if totalCost < dists[neighbor] {
				pq.Update(PuzzleNodeHeapItem{neighbor, dists[neighbor]}, PuzzleNodeHeapItem{neighbor, totalCost})
				dists[neighbor] = totalCost
				prevs[neighbor] = []PuzzleNode{node}
			} else if totalCost == dists[neighbor] {
				prevs[neighbor] = append(prevs[neighbor], node)
			}
		}
	}

	return dists, prevs
}

//go:embed input
var input string

func main() {
	util.SolveChosenPart(func(isPart2 bool) int64 {
		problem := Parse(input)
		//fmt.Println(problem)
		return solve(problem, isPart2)
	})
}
