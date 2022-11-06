package main

import "container/heap"

type element struct {
	value int
	count int
}

type priorityQueue []element

func (q priorityQueue) Len() int           { return len(q) }
func (q priorityQueue) Less(i, j int) bool { return q[i].count < q[j].count }
func (q priorityQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *priorityQueue) Push(x interface{}) {
	*q = append(*q, x.(element))
}

func (q *priorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func topKFrequent2(nums []int, k int) []int {
	count := make(map[int]int)

	for _, n := range nums {
		count[n]++
	}

	q := &priorityQueue{}
	heap.Init(q)
	for val, cnt := range count {
		heap.Push(q, element{value: val, count: cnt})
		if q.Len() > k {
			heap.Pop(q)
		}
	}

	res := make([]int, k, k)
	for i := 0; i < k; i++ {
		v := heap.Pop(q)
		if s, ok := v.(element); ok {
			res[i] = s.value
		}
	}

	return res
}
