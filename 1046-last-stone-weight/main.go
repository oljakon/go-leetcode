package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	intHeap := IntHeap(stones)
	heap.Init(&intHeap)

	for intHeap.Len() > 1 {
		x, y := heap.Pop(&intHeap).(int), heap.Pop(&intHeap).(int)
		if x != y {
			heap.Push(&intHeap, x-y)
		}
	}

	if intHeap.Len() == 0 {
		return 0
	}

	return heap.Pop(&intHeap).(int)
}
