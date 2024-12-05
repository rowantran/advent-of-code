package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

//go:embed input
var input string

// shamelessly stolen from: https://pkg.go.dev/container/heap#Interface
// A TokenHeap is a min-heap of tokens.
type TokenHeap []Token

func (h TokenHeap) Len() int           { return len(h) }
func (h TokenHeap) Less(i, j int) bool { return h[i].index < h[j].index }
func (h TokenHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TokenHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Token))
}

func (h *TokenHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type TokenType int
const (
	Mul TokenType = iota
	Do
	Dont
)

type Token struct {
	index int
	ttype TokenType
}

func findAllLocations(line string, query string) []int {
	var locations []int
	for i := range line {
		found := true

		for j := range query {
			if (i+j) >= len(line) || line[i+j] != query[j] {
				found = false
				break
			}
		}
			
		if found {
			locations = append(locations, i)
		}
	}
	return locations
}

// try to process a single mul instruction starting at index i, return result or 0 if no valid instruction
func tryProcess(line string, i int) int {
	// mul(123,)
	// ^i      ^end
	end := min(len(line), i+8)
	searchWindow := line[i+4:end]
	commaLoc := strings.Index(searchWindow, ",")
	if commaLoc == -1 {
		return 0
	} 

	first, err := strconv.Atoi(searchWindow[:commaLoc])
	if err != nil {
		//fmt.Printf("error: %v\n", err)
		return 0
	}

	// index AFTER comma within line
	start := i+5+commaLoc
	// mul(123,939) 
	//         ^s  ^end
	end = min(len(line), start+4)
	searchWindow = line[start:end]
	closeLoc := strings.Index(searchWindow, ")")
	if closeLoc == -1 {
		return 0
	}

	second, err := strconv.Atoi(searchWindow[:closeLoc])
	if err != nil {
		//fmt.Printf("error: %v\n", err)
		return 0
	}

	return first * second
}

// return the result of processing valid mul instructions within the line
func process(line string) int {
	total := 0
	locs := findAllLocations(line, "mul(")

	for _, i := range locs {
		total += tryProcess(line, i)
	}

	return total
}

// input: line, whether mults are enabled at start of line
// output: result as defined above, whether mults are disabled at end of line
func processWithToggles(line string, enabled bool) (int, bool) {
	total := 0

	multLocs := findAllLocations(line, "mul(")
	dos := findAllLocations(line, "do()")
	donts := findAllLocations(line, "don't()")

	pq := &TokenHeap{}
	for _, loc := range multLocs {
		heap.Push(pq, Token{loc, Mul})
	}
	for _, loc := range dos {
		heap.Push(pq, Token{loc, Do})
	}
	for _, loc := range donts {
		heap.Push(pq, Token{loc, Dont})
	}

	for pq.Len() > 0 {
		token := heap.Pop(pq)
		ttype := token.(Token).ttype

		switch ttype {
		case Mul:
			if enabled {
				total += tryProcess(line, token.(Token).index)
			}
		case Do:
			enabled = true
		case Dont:
			enabled = false
		}
	}

	return total, enabled
}

func part1() {
	result := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		result += process(line)
	}
	fmt.Println("result:", result)
}

func part2() {
	result := 0
	enabled := true

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var lineResult int
		line := scanner.Text()
		lineResult, enabled = processWithToggles(line, enabled)
		result += lineResult
	}

	fmt.Println("result:", result)
}

func main() {
	util.RunChosenPart(part1, part2)
}
