package main

import (
	"container/heap"
	"fmt"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	var ans int64 = 0
	var h IntHeap = gifts
	heap.Init(&h)
	for i := 0; i < k; i++ {
		maxGift := heap.Pop(&h).(int)
		// fmt.Println(maxGift)
		sqrtGift := int(math.Sqrt(float64(maxGift)))
		// ans += int64(sqrtGift)
		heap.Push(&h, sqrtGift)
	}

	for h.Len() > 0 {
		ans += int64(heap.Pop(&h).(int))
	}
	return ans
}

func main() {
	var gifts []int = []int{25, 64, 9, 4, 100}
	var k int = 4
	fmt.Println(pickGifts(gifts, k))
}
