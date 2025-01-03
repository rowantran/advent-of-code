package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	_ "embed"

	"github.com/rowantran/advent-of-code/2024/util"
)

type Computer struct {
	regs    [3]int
	program []int
	ip      int
}

func Parse(input string) Computer {
	var c Computer

	parts := strings.Split(input, "\n\n")

	reg := 0
	scanner := bufio.NewScanner(strings.NewReader(parts[0]))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		c.regs[reg] = util.MustAtoi(fields[len(fields)-1])
		reg++
	}

	programNums := strings.Split(strings.Fields(parts[1])[1], ",")
	c.program = make([]int, len(programNums))
	for i := range programNums {
		c.program[i] = util.MustAtoi(programNums[i])
	}

	c.ip = 0

	return c
}

func (c *Computer) Run() []int {
	output := []int{}

	for c.ip < len(c.program) {
		operand := c.program[c.ip+1]

		switch c.program[c.ip] {
		case 0:
			// adv
			c.regs[0] = c.div(operand)
			c.ip += 2
		case 1:
			c.regs[1] = c.bxl(operand)
			c.ip += 2
		case 2:
			c.regs[1] = c.bst(operand)
			c.ip += 2
		case 3:
			c.ip = c.jnz(operand)
		case 4:
			c.regs[1] = c.bxc()
			c.ip += 2
		case 5:
			// out
			out := c.bst(operand)
			output = append(output, out)
			c.ip += 2
		case 6:
			// bdv
			c.regs[1] = c.div(operand)
			c.ip += 2
		case 7:
			// cdv
			c.regs[2] = c.div(operand)
			c.ip += 2
		}
	}

	return output
}

func (c *Computer) RunAndPrint() {
	output := c.Run()
	outputStrings := make([]string, len(output))
	for i := range output {
		outputStrings[i] = strconv.Itoa(output[i])
	}
	fmt.Println(strings.Join(outputStrings, ","))
}

func (c *Computer) parseCombo(arg int) int {
	switch {
	case arg <= 3:
		return arg
	case arg <= 6:
		return c.regs[arg-4]
	default:
		panic("invalid combo operand")
	}
}

func (c *Computer) div(arg int) int {
	arg = c.parseCombo(arg)
	return c.regs[0] / util.ExpInt(2, arg)
}

func (c *Computer) bxl(arg int) int {
	return c.regs[1] ^ arg
}

func (c *Computer) bst(arg int) int {
	arg = c.parseCombo(arg)
	return arg % 8
}

func (c *Computer) jnz(arg int) int {
	if c.regs[0] == 0 {
		return c.ip + 2
	} else {
		return arg
	}
}

func (c *Computer) bxc() int {
	return c.regs[1] ^ c.regs[2]
}

func solveQuine(c Computer) int {
	a := solveQuineRec(c, 0, 0)
	return a
}

// invariant: A generates the last i opcodes of the program successfully
func solveQuineRec(c Computer, a int, i int) int {
	if i == len(c.program) {
		return a
	}

	// try to find a value of n for the next 3 bits of A which will successfully generate
	// the opcode at index i (counting from the back)
	for n := range 8 {
		an := a*8 + n
		testComputer := c
		testComputer.regs[0] = an
		output := testComputer.Run()
		if output[len(output)-1-i] == c.program[len(c.program)-1-i] {
			// A_n is a valid extension of A to generate the last i+1 opcodes successfully
			an = solveQuineRec(c, an, i+1)
			if an != 0 {
				return an
			}
		}
	}

	return 0
}

//go:embed input
var input string

func main() {
	util.SolveChosenPart(func(isPart2 bool) int64 {
		computer := Parse(input)
		//fmt.Println(problem)
		if isPart2 {
			a := solveQuine(computer)
			computer.regs[0] = a
			computer.RunAndPrint()
			return int64(a)
		} else {
			computer.RunAndPrint()
			return int64(0)
		}
	})
}
