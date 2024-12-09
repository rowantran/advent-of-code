package util

import "fmt"

type Set[T comparable] map[T]struct{}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) AddAll(vals []T) {
	for _, v := range vals {
		s.Add(v)
	}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) String() string {
	res := "{"

	i := 0
	for key := range s {
		res += fmt.Sprintf("%v", key)
		// if printing last element
		if i+1 == s.Size() {
			res += "}"
		} else {
			res += ", "
		}
		i += 1
	}

	return res
}

func CreateSet[T comparable]() Set[T] {
	return make(Set[T])
}
