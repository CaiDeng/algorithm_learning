package helper

import (
	"errors"
	"fmt"
	"math"
)

// Heap 最大堆或最小堆
type Heap struct {
	elem []int
	more func(a, b int) bool
}

// heapify 维护堆的性质
func (h *Heap) heapify(i int) error {
	if len(h.elem) == 0 {
		return nil
	}
	if i >= len(h.elem) {
		return errors.New("params is too long")
	}

	extremum := i
	if left, err := h.left(i); err == nil && h.more(h.elem[left], h.elem[extremum]) {
		extremum = left
	}

	if right, err := h.right(i); err == nil && h.more(h.elem[right], h.elem[extremum]) {
		extremum = right
	}

	if extremum == i {
		return nil
	}

	h.elem[extremum], h.elem[i] = h.elem[i], h.elem[extremum]

	return h.heapify(extremum)
}

// NewHeap 建堆
func NewHeap(a []int, more func(a, b int) bool) (Heap, error) {
	if a == nil {
		return Heap{}, fmt.Errorf("slice is nil")
	}
	heap := Heap{
		elem: a,
		more: more,
	}

	for i := (len(a) - 1) >> 1; i >= 0; i-- {
		if err := heap.heapify(i); err != nil {
			return Heap{}, err
		}
	}
	return heap, nil
}

// Insert 插入一个结点
func (h *Heap) Insert(val int) error {
	extremum := math.MinInt
	if h.more(extremum, 0) {
		extremum = math.MaxInt
	}
	h.elem = append(h.elem, extremum)
	if err := h.Set(len(h.elem)-1, val); err != nil {
		return err
	}
	return nil
}

// DeleteExtremum 删除并获取堆顶元素
func (h *Heap) DeleteExtremum() (int, error) {
	if len(h.elem) == 0 {
		return 0, fmt.Errorf("heap size is 0")
	}

	extremum := h.elem[0]
	h.elem[0] = h.elem[len(h.elem)-1]
	h.elem = h.elem[:len(h.elem)-1]
	if err := h.heapify(0); err != nil {
		return 0, err
	}

	return extremum, nil
}

// Set 修改结点的值
func (h *Heap) Set(index, val int) error {
	if index >= len(h.elem) {
		return errors.New("param index is more than heap size")
	}

	// val 比当前结点值小，只需维护堆的性质
	if h.more(h.elem[index], val) {
		h.elem[index] = val
		if err := h.heapify(index); err != nil {
			return err
		}
	}

	// val 比当前结点值大，只需向上检查堆的性质
	h.elem[index] = val
	parent, err := h.parent(index)
	for err == nil && h.more(h.elem[index], h.elem[parent]) {
		h.elem[index], h.elem[parent] = h.elem[parent], h.elem[index]
		index = parent
		parent, err = h.parent(index)
	}

	return nil
}

// Extremum 获取堆顶元素
func (h *Heap) Extremum() (int, error) {
	if len(h.elem) == 0 {
		return 0, errors.New("heap size is 0, no extremum")
	}
	return h.elem[0], nil
}

func (h *Heap) parent(index int) (int, error) {
	if index <= 0 || index >= len(h.elem) {
		return 0, fmt.Errorf("index is invalid, can not be less than 0 or more than heap size")
	}
	return (index - 1) >> 1, nil
}

func (h *Heap) left(index int) (int, error) {
	left := index<<1 + 1
	if index < 0 || left >= len(h.elem) {
		return 0, fmt.Errorf("index is invalid, can not be less than 0 or index's left child index can not be than heap size")
	}
	return left, nil
}

func (h *Heap) right(index int) (int, error) {
	right := index<<1 + 2
	if index < 0 || right >= len(h.elem) {
		return 0, fmt.Errorf("index is invalid, can not be less than 0 or index's right child index can not be than heap size")
	}
	return right, nil
}

// Sort 堆排序，最大堆是增序,最小堆是降序
func (h *Heap) Sort() ([]int, error) {
	backup := h.elem
	if len(h.elem) == 0 {
		return []int{}, nil
	}

	for len(h.elem) > 1 {
		length := len(h.elem)
		h.elem[length-1], h.elem[0] = h.elem[0], h.elem[length-1]
		h.elem = h.elem[:length-1]
		if err := h.heapify(0); err != nil {
			return []int{}, err
		}
	}
	return backup, nil
}

// Size 堆的大小
func (h *Heap) Size() int {
	return len(h.elem)
}
