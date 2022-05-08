package array

import "math"

type MinStack struct {
	elems    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	initMinStack := MinStack{
		make([]int, 0),
		[]int{math.MaxInt},
	}
	return initMinStack
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (this *MinStack) Push(x int) {
	this.elems = append(this.elems, x)
	currentMin := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(currentMin, x))
}

func (this *MinStack) Pop() {
	if len(this.elems) == 0 {
		return
	}
	this.elems = this.elems[:len(this.elems)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.elems[len(this.elems)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}
