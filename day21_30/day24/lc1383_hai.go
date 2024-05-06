package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap []int

// Implement the heap.Interface for MinHeap
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	const mod int = 1e9 + 7
	type pair struct{ s, e int }
	ps := make([]pair, n)
	for i := range ps {
		ps[i] = pair{speed[i], efficiency[i]}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].e > ps[j].e })

	h := &MinHeap{}
	heap.Init(h)
	var sum, res int
	for _, p := range ps {
		if h.Len() == k {
			sum = (sum - int(heap.Pop(h).(int)) + mod) % mod
		}
		sum = (sum + int(p.s)) % mod
		heap.Push(h, p.s)
		res = max(res, sum*p.e)
	}
	return res % mod
}

func main() {
	n := 6
	speed := []int{2, 10, 3, 1, 5, 8}
	efficiency := []int{5, 4, 3, 9, 7, 2}
	k := 2
	fmt.Println(maxPerformance(n, speed, efficiency, k))
}
