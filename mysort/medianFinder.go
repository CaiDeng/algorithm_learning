package mysort

import (
	"container/heap"
	"sort"
)

type minhp struct {
	sort.IntSlice
}

func (h *minhp) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *minhp) Pop() interface{} {
	a := h.IntSlice
	extremum := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return extremum
}

type maxhp struct {
	minhp
}

func (h maxhp) Less(i, j int) bool {
	return i > j
}

// MedianFinder 中位数数据结构
type MedianFinder struct {
	minHeap minhp
	maxHeap maxhp
}

// Constructor initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

// AddNum 插入一个数字
func (m *MedianFinder) AddNum(num int) {
	if m.maxHeap.Len() == m.minHeap.Len() {
		heap.Push(&m.maxHeap, num)
		max := heap.Pop(&m.maxHeap)
		heap.Push(&m.minHeap, max)

	} else {
		heap.Push(&m.minHeap, num)
		min := heap.Pop(&m.minHeap)
		heap.Push(&m.maxHeap, min)
	}
}

// FindMedian 获取中位数
func (m *MedianFinder) FindMedian() float64 {

	if m.minHeap.Len() == m.maxHeap.Len() {
		min := m.minHeap.IntSlice[0]
		max := m.maxHeap.minhp.IntSlice[0]
		return float64(max+min) / 2
	}
	min := m.minHeap.IntSlice[0]
	return float64(min)
}
