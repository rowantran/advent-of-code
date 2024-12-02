package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

//go:embed input
var input string

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func parse(input string) ([]int, []int) {
	list1, list2 := []int{}, []int{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		list1 = append(list1, mustAtoi(line[0]))
		list2 = append(list2, mustAtoi(line[1]))
	}
	return list1, list2
}

func part1() {
	list1, list2 := parse(input)

	sort.Ints(list1)
	sort.Ints(list2)

	distance := 0
	for i, v1 := range list1 {
		v2 := list2[i]
		distance += abs(v1 - v2)
	}

	fmt.Printf("total distance: %d\n", distance)
}

func part2() {
	list1, list2 := parse(input)

	mults := make(map[int]int)
	for _, val := range list2 {
		mults[val] += 1
	}

	score := 0
	for _, val := range list1 {
		score += (val * mults[val])
	}

	fmt.Println("total similarity score:", score)
}

func main() {
	util.RunChosenPart(part1, part2)
}
