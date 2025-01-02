package util

import "container/heap"

type Heap[T comparable] struct {
	data     []T
	indices  map[T]int
	lessFunc func(T, T) bool
}

func NewHeap[T comparable](lessFunc func(T, T) bool) Heap[T] {
	return Heap[T]{
		data:     make([]T, 0),
		indices:  make(map[T]int),
		lessFunc: lessFunc,
	}
}

func (h *Heap[T]) Len() int           { return len(h.data) }
func (h *Heap[T]) Less(i, j int) bool { return h.lessFunc(h.data[i], h.data[j]) }
func (h *Heap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.indices[h.data[i]] = i
	h.indices[h.data[j]] = j
}
func (h *Heap[T]) Push(x any) {
	h.indices[x.(T)] = len(h.data)
	h.data = append(h.data, x.(T))
}
func (h *Heap[T]) Pop() any {
	v := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.indices[v] = -1
	return v
}
func (h *Heap[T]) Update(val T, newVal T) {
	i := h.indices[val]
	if i == -1 {
		// val is not in heap
		return
	}

	h.data[i] = newVal
	h.indices[val] = -1
	h.indices[newVal] = i
	heap.Fix(h, i)
}
