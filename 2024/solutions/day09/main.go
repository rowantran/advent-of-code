package main

import (
	"fmt"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

//go:embed input
var input string

type File struct {
	id           int
	originalSize int
	size         int
}

type PuzzleInput struct {
	files []File
	gaps  []int
}

func Parse(input string) PuzzleInput {
	var problem PuzzleInput
	for i, sizeRune := range input[:len(input)-1] {
		size := runeToInt(sizeRune)
		// ignore newline at end
		if size < 0 {
			continue
		}

		if i%2 == 0 {
			problem.files = append(problem.files, File{i / 2, size, size})
		} else {
			problem.gaps = append(problem.gaps, size)
		}
	}
	return problem
}

func (p PuzzleInput) CompactAndChecksum(allowFragmentation bool) int64 {
	var ans int64

	// starting at index i, fill in n locations with as much file content as possible,
	// padding with empty space if needed
	var processChunk = func(i *int, n int, file File) {
		contentEnd := *i + file.size
		end := *i + n
		for ; *i < end; *i++ {
			//fmt.Printf("processing index %d with file id %d\n", *i, file.id)
			if *i < contentEnd {
				ans += int64(*i * file.id)
			}
		}
	}

	var isValidFileForGap = func(gap int, i int) bool {
		if p.files[i].size == 0 {
			return false
		}
		return allowFragmentation || p.files[i].size <= gap
	}

	i := 0
	for f := range p.files {
		// pop and process next file from the beginning
		file := p.files[f]
		processChunk(&i, file.originalSize, file)

		// process next gap, if it exists
		if f >= len(p.gaps) {
			continue
		}
		gap := p.gaps[f]
		for gap > 0 {
			var j int
			for j = len(p.files) - 1; j > f && !isValidFileForGap(gap, j); j-- {
			}
			// if no valid file was found, the rest of the gap will stay empty
			if j <= f {
				i += gap
				break
			}

			filled := min(gap, p.files[j].size)
			processChunk(&i, filled, p.files[j])
			gap -= filled
			p.files[j].size -= filled
		}
	}

	return ans
}

func runeToInt(r rune) int {
	return int(r - '0')
}

func part1() {
	problem := Parse(input)
	ans := problem.CompactAndChecksum(true)
	//fmt.Println(problem)
	fmt.Println("answer:", ans)
}

func part2() {
	problem := Parse(input)
	ans := problem.CompactAndChecksum(false)
	fmt.Println("answer:", ans)
}

func main() {
	util.RunChosenPart(part1, part2)
}
