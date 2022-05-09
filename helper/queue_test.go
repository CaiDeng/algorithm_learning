package helper

import (
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	d, _ := q.Peek()
	if 1 != q.Size() || d.(int) != 1 {
		t.Error("queue size and enqueue failed")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue()

	_, success := q.Dequeue()
	if success || q.Size() > 0 {
		t.Error("empty queue dequeue failed")
	}

	q.Enqueue(1)
	d, success := q.Dequeue()
	if d.(int) != 1 || !success || q.Size() > 0 {
		t.Error("queue dequeue failed")
	}
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue()

	_, success := q.Peek()
	if success || q.Size() > 0 {
		t.Error("empty queue dequeue failed")
	}

	q.Enqueue(1)
	d, success := q.Peek()
	if 1 != d.(int) || !success || q.Size() != 1 {
		t.Error("queue peek failed")
	}
}
