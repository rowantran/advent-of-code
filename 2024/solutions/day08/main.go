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

				for _, anti := range p.pairwiseAntinodes(a1, a2, p2) {
					if p.IsValidLocation(anti) {
						antinodes.Add(anti)
					}
				}
			}
		}
	}

	return antinodes
}

func (p PuzzleInput) pairwiseAntinodes(a1 Vec2, a2 Vec2, p2 bool) []Vec2 {
	delta := a2.Sub(a1)
	var antinodes []Vec2

	if p2 {
		anti := a1
		for i := 0; p.IsValidLocation(anti); {
			antinodes = append(antinodes, anti)
			i++
			anti = a1.Add(delta.Mul(i))
		}

		anti = a1.Add(delta.Mul(-1))
		for i := -1; p.IsValidLocation(anti); {
			antinodes = append(antinodes, anti)
			i--
			anti = a1.Add(delta.Mul(i))
		}
	} else {
		antinodes = appendIfValid(antinodes, a2.Add(delta), p)
		antinodes = appendIfValid(antinodes, a1.Sub(delta), p)
	}
	return antinodes
}

func appendIfValid(locs []Vec2, loc Vec2, p PuzzleInput) []Vec2 {
	if p.IsValidLocation(loc) {
		return append(locs, loc)
	} else {
		return locs
	}
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
