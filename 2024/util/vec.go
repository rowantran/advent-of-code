package util

import "strings"

type Vec2[T ~int | ~int64] [2]T

// parse a Vec2[int] from string in format "x,y"
func NewVec2Int(str string) Vec2[int] {
	parts := strings.Split(str, ",")
	return Vec2[int]{MustAtoi(parts[0]), MustAtoi(parts[1])}
}

// parse a Vec2[int64] from string in format "x,y"
func NewVec2Int64(str string) Vec2[int64] {
	parts := strings.Split(str, ",")
	return Vec2[int64]{MustAtoiInt64(parts[0]), MustAtoiInt64(parts[1])}
}

func (a Vec2[T]) Add(b Vec2[T]) Vec2[T] {
	return Vec2[T]{a[0] + b[0], a[1] + b[1]}
}

func (v Vec2[T]) Mul(c T) Vec2[T] {
	return Vec2[T]{c * v[0], c * v[1]}
}

func (a Vec2[T]) Sub(b Vec2[T]) Vec2[T] {
	return a.Add(b.Mul(-1))
}

func (v Vec2[T]) Parts() (T, T) {
	return v[0], v[1]
}

func (a Vec2[T]) Dot(b Vec2[T]) T {
	return a[0]*b[0] + a[1]*b[1]
}

func (a Vec2[T]) IsOrthogonal(b Vec2[T]) bool {
	return a.Dot(b) == 0
}
