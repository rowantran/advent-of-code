package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Vec2 = util.Vec2

//go:embed input
var input string

type PuzzleInput struct {
	rows     int
	cols     int
	antennas map[rune][]Vec2
}

func Parse(input string) PuzzleInput {
	var problem PuzzleInput
	problem.antennas = make(map[rune][]Vec2)

	scanner := bufio.NewScanner(strings.NewReader(input))
	var i int
	var line string
	for i = 0; scanner.Scan(); i++ {
		line = scanner.Text()
		for j, ch := range line {
			if ch != '.' {
				problem.antennas[ch] = append(problem.antennas[ch], Vec2{i, j})
			}
		}
	}
	problem.rows = i
	problem.cols = len(line)

	return problem
}

func (p PuzzleInput) IsValidLocation(loc Vec2) bool {
	row, col := loc[0], loc[1]
	return row >= 0 && row < p.rows && col >= 0 && col < p.cols
}

func (p PuzzleInput) Antinodes(p2 bool) util.Set[Vec2] {
	antinodes := make(util.Set[Vec2])

	for _, antennas := range p.antennas {
		for i, a1 := range antennas {
			for j, a2 := range antennas {
				if i <= j {
					continue
				}
				antinodes.AddAll(p.pairwiseAntinodes(a1, a2, p2))
			}
		}
	}

	return antinodes
}

func (p PuzzleInput) pairwiseAntinodes(a1 Vec2, a2 Vec2, p2 bool) []Vec2 {
	delta := a2.Sub(a1)
	var antinodes []Vec2

	// try to add a1 + delta*i, a1 + delta*(i + increment), a1 + delta*(i + 2*increment), ...
	var addNodes = func(i int, increment int) {
		for anti := a1.Add(delta.Mul(i)); p.IsValidLocation(anti); anti = a1.Add(delta.Mul(i)) {
			antinodes = append(antinodes, anti)
			i += increment
			if increment == 0 {
				break
			}
		}
	}

	if p2 {
		// add all points a1 + c*(a2-a1); c is an integer && the resulting point is in the grid

		// input is constructed s.t. dx, dy := (delta[0], delta[1]) are coprime for all pairs of antennas
		// otherwise, we would need to handle cases where fractional c would result in a valid
		// coordinate, e.g. (dx,dy)=(2,2) -> a1 + 0.5(a2-a1) = a1 + {1, 1} is a collinear point with int coordinates
		addNodes(0, 1)
		addNodes(-1, -1)
	} else {
		// same as above, but only consider c = -1, 2 because we need the point to be twice as far from
		// one antenna as the other
		addNodes(-1, 0)
		addNodes(2, 0)
	}
	return antinodes
}

func part1() {
	problem := Parse(input)
	ans := len(problem.Antinodes(false))
	fmt.Println("answer:", ans)
}

func part2() {
	problem := Parse(input)
	ans := len(problem.Antinodes(true))
	fmt.Println("answer:", ans)
}

func main() {
	util.RunChosenPart(part1, part2)
}
