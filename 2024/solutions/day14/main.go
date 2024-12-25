package main

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Vec2 = util.Vec2[int]

type Robot struct {
	startPos Vec2
	velocity Vec2
}

func ParseRobot(line string) Robot {
	fields := strings.Fields(line)
	// strip off the initial "p=" or "v=" from the two fields, then convert to vec2
	return Robot{util.NewVec2Int(fields[0][2:]), util.NewVec2Int(fields[1][2:])}
}

func (r Robot) FinalQuadrant() ([2]bool, error) {
	fp := r.finalPos()
	x, y := fp.Parts()
	if x == width/2 || y == height/2 {
		return [2]bool{}, fmt.Errorf("final position was in center row or column")
	}
	return [2]bool{x > width/2, y > height/2}, nil
}

func (r Robot) finalPos() Vec2 {
	unwrappedPos := r.startPos.Add(r.velocity.Mul(iterations))
	finalPos := Vec2{mod(unwrappedPos[0], width), mod(unwrappedPos[1], height)}
	//fmt.Printf("final pos for robot with startPos %v, vel %v = %v\n", r.startPos, r.velocity, finalPos)
	return finalPos
}

func mod(a int, b int) int {
	return ((a%b)+b)%b
}

type PuzzleInput struct {
	robots []Robot
}

func Parse(input string) PuzzleInput {
	var problem PuzzleInput
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		problem.robots = append(problem.robots, ParseRobot(line))
	}
	return problem
}

func solve(p PuzzleInput, isPart2 bool) int64 {
	quadrantCounts := make(map[[2]bool]int)
	for _, r := range p.robots {
		quad, err := r.FinalQuadrant()
		if err == nil {
			quadrantCounts[quad]++
		}
	}

	ans := 1
	for _, count := range quadrantCounts {
		ans *= count
	}
	return int64(ans)
}

//go:embed input
var input string

const iterations = 100
//const width, height = 11, 7
const width, height = 101, 103

func main() {
	util.SolveChosenPart(func(isPart2 bool) int64 {
		problem := Parse(input)
		//fmt.Println(problem)
		return solve(problem, isPart2)
	})
}
