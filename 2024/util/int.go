package util

import "strconv"

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Abs(i int) int {
	if (i < 0) {
		return -i
	} else {
		return i
	}
}

func RuneToInt(r rune) int {
	return int(r - '0')
}
