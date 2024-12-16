package util

import "strconv"

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func MustAtoiInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func RuneToInt(r rune) int {
	return int(r - '0')
}

func DigitCountInt64(a int64) int {
	var p int = 0
	var base int64 = 1

	if a == 0 {
		return 1
	}

	for base <= a {
		p++
		base *= 10
	}

	// p is the smallest int s.t. 10^p > a
	return p
}

func DigitCount(a int) int {
	return DigitCountInt64(int64(a))
}

// calculate b to the nth power
func ExpInt64(b int64, n int) int64 {
	res := int64(1)
	for range n {
		res *= b
	}
	return res
}

func ExpInt(b int, n int) int {
	return int(ExpInt64(int64(b), n))
}
