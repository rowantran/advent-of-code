package main

import (
	"fmt"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

//go:embed input
var input string

type File struct {
	id int
	originalSize int
	size int
}

type PuzzleInput struct {
	files []File
	gaps []int
}

func Parse(input string) PuzzleInput {
	var problem PuzzleInput
	for i, sizeRune := range input[:len(input)-1] {
		size := runeToInt(sizeRune)
		// ignore newline at end
		if size < 0 {
			continue
		}

		if i % 2 == 0 {
			problem.files = append(problem.files, File{i/2, size, size})
		} else {
			problem.gaps = append(problem.gaps, size)
		}
	}
	return problem
}

func (p PuzzleInput) CompactAndChecksum(allowFragmentation bool) int64 {
	var ans int64

	var processFile = func(i *int, file File, size int) {
		end := *i + size
		for ; *i < end; *i++ {
			fmt.Printf("processing index %d with file id %d\n", *i, file.id)
			ans += int64(*i * file.id)
		}
	}

	i := 0
	for f, file := range p.files {
		// pop and process next file from the beginning
		processFile(&i, file, file.size)
		if (file.size < file.originalSize) {
			i += (file.originalSize - file.size)
		}

		// process next gap:
		// choose files from the end until the gap is filled, removing files if the remaining gap
		// is larger than their full size
		if len(p.gaps) == 0 {
			continue
		}

		gapToFill := p.gaps[0]
		p.gaps = p.gaps[1:]
		for gapToFill > 0 {
			// if there are no remaining files to use, the rest of the gap will stay empty
			var isValidFileForGap = func(i int) bool {
				if p.files[i].size == 0 {
					return false
				}
				return allowFragmentation || p.files[i].size <= gapToFill
			}

			var j int
			for j = len(p.files)-1; j > f && !isValidFileForGap(j); j-- {}
			// no file was found
			if j <= f {
				i += gapToFill
				break
			}

			file = p.files[j]
			fileBlocksToUse := min(gapToFill, file.size)

			processFile(&i, file, fileBlocksToUse)
			gapToFill -= fileBlocksToUse
			p.files[j].size -= fileBlocksToUse
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
