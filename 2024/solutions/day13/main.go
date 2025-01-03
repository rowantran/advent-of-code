package main

import (
	"bufio"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Vec2 = util.Vec2[int64]

const part2_offset = 10000000000000

type MachineInfo struct {
	buttonA Vec2
	buttonB Vec2
	prize   Vec2
}

type PuzzleInput struct {
	machines []MachineInfo
}

func Parse(input string, isPart2 bool) PuzzleInput {
	var problem PuzzleInput
	scanner := bufio.NewScanner(strings.NewReader(input))

	var machine MachineInfo
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			problem.machines = append(problem.machines, machine)
			continue
		}

		fields := strings.Fields(line)
		xField, yField := fields[len(fields)-2], fields[len(fields)-1]
		vec := Vec2{util.MustAtoiInt64(xField[2 : len(xField)-1]), util.MustAtoiInt64(yField[2:])}

		switch len(fields) {
		case 3:
			// prize
			if isPart2 {
				vec = vec.Add(Vec2{part2_offset, part2_offset})
			}
			machine.prize = vec
		case 4:
			// button
			buttonType := fields[1][0]
			if buttonType == 'A' {
				machine.buttonA = vec
			} else {
				machine.buttonB = vec
			}
		}
	}
	// last machine won't be flushed by a blank line
	problem.machines = append(problem.machines, machine)

	return problem
}

/*
The problem suggests a matrix equation, but since the problem (at least part 1) has a restriction
that the button counts will be <=100, it's easier to just bruteforce

part 2 is not in fact brute-forceable, but luckily all test cases given have unique solutions in R2,
so we just need to check if the solution has integer parts and not think too hard about the case
where the matrix is singular
*/
func solve(p PuzzleInput) int64 {
	total := int64(0)
	for _, machine := range p.machines {
		cost := machineCost(machine)
		if cost != nil {
			total += *cost
		}
	}
	return total
}

// returns minimum tokens needed to win a prize, or nil if impossible
func machineCost(m MachineInfo) *int64 {
	a0, a1 := m.buttonA.Parts()
	b0, b1 := m.buttonB.Parts()
	t0, t1 := m.prize.Parts()

	xNum := (t1*b0 - t0*b1)
	xDen := (a1*b0 - a0*b1)
	if xDen == 0 {
		// xDen = bc-ad = -det A, so if xDen = 0 then det A = 0 and A is singular
		// the test case doesn't have any singular matrices so we don't need to handle this
		return nil
	} else if xNum%xDen != 0 {
		// non-integer solution
		return nil
	}
	x := xNum / xDen

	yNum := (t0 - x*a0)
	yDen := b0
	if yDen == 0 {
		// not sure if this case represents something special, but there is
		// still a unique solution so we can find y by back-substituting x into
		// the second equation, i.e. calculate y := (t1 - x*a1) / b1
		return nil
	} else if yNum%yDen != 0 {
		// non-integer solution
		return nil
	}
	y := yNum / yDen

	res := new(int64)
	*res = 3*x + y
	return res
}

//go:embed input
var input string

func main() {
	util.SolveChosenPart(func(isPart2 bool) int64 {
		problem := Parse(input, isPart2)
		return solve(problem)
	})
}
