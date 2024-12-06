package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"regexp"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

//go:embed input
var input string

var mulRegex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
var doRegex = regexp.MustCompile(`do\(\)`)
var dontRegex = regexp.MustCompile(`don't\(\)`)

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
	// fields below should only be used for Mul instructions
	first int
	second int
}

// match: output of FindAllStringSubmatchIndex on a string
// submatchIndex: submatch index where 0 is the full match, 1 is first submatch etc.
func getSubmatchString(str string, match []int, submatchIndex int) string {
	submatchIndices := match[2*submatchIndex:2*submatchIndex+2]
	return str[submatchIndices[0]:submatchIndices[1]]
}

// return the result of processing valid mul instructions within the line
func process(line string) int {
	total := 0
	locs := mulRegex.FindAllStringSubmatch(line, -1)

	for _, match := range locs {
		first, second := util.MustAtoi(match[1]), util.MustAtoi(match[2])
		total += (first * second)
	}

	return total
}

// input: line, whether mults are enabled at start of line
// output: result as defined above, whether mults are disabled at end of line
func processWithToggles(line string, enabled bool) (int, bool) {
	total := 0

	multLocs := mulRegex.FindAllStringSubmatchIndex(line, -1)
	dos := doRegex.FindAllStringIndex(line, -1)
	donts := dontRegex.FindAllStringIndex(line, -1)

	pq := &TokenHeap{}
	for _, match := range multLocs {
		index := match[0]
		first := util.MustAtoi(getSubmatchString(line, match, 1))
		second := util.MustAtoi(getSubmatchString(line, match, 2))
		heap.Push(pq, Token{index: index, ttype: Mul, first: first, second: second})
	}
	for _, match := range dos {
		heap.Push(pq, Token{index: match[0], ttype: Do})
	}
	for _, match := range donts {
		heap.Push(pq, Token{index: match[0], ttype: Dont})
	}

	for pq.Len() > 0 {
		token := heap.Pop(pq).(Token)
		ttype := token.ttype

		switch ttype {
		case Mul:
			if enabled {
				total += (token.first * token.second)
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
