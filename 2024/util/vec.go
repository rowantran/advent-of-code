package util

type Vec2 [2]int

func (a Vec2) Add(b Vec2) Vec2 {
	return Vec2{a[0]+b[0], a[1]+b[1]}
}

func (v Vec2) Mul(c int) Vec2 {
	return Vec2{c*v[0], c*v[1]}
}

func (a Vec2) Sub(b Vec2) Vec2 {
	return a.Add(b.Mul(-1))
}
