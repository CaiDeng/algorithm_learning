package array

type goStack struct {
	elems []int
}

func (this *goStack) push(value int) {
	this.elems = append(this.elems, value)
}

func (this *goStack) pop() (topElem int) {
	if len(this.elems) == 0 {
		return -1
	}
	topElem = this.elems[len(this.elems)-1]
	this.elems = this.elems[:len(this.elems)-1]
	return topElem
}

func (this *goStack) isEmpty() bool {
	if len(this.elems) == 0 {
		return true
	}
	return false
}

func newgoStack() goStack {
	initGoStack := goStack{
		make([]int, 0),
	}
	return initGoStack
}

type CQueue struct {
	inputStack  goStack
	outputStack goStack
}

func NewCQueue() CQueue {
	initCQueue := CQueue{
		newgoStack(),
		newgoStack(),
	}
	return initCQueue
}

func (this *CQueue) AppendTail(value int) {
	this.inputStack.push(value)
}

func (this *CQueue) DeleteHead() int {
	if !this.outputStack.isEmpty() {
		return this.outputStack.pop()
	}

	if this.inputStack.isEmpty() {
		return -1
	}

	for !this.inputStack.isEmpty() {
		this.outputStack.push(this.inputStack.pop())
	}
	return this.outputStack.pop()
}
