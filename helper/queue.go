package helper

// Queue 队列
type Queue struct {
	list []interface{}
}

// NewQueue 队列初始化
func NewQueue() Queue {
	return Queue{make([]interface{}, 0)}
}

// Enqueue 进入队列
func (q *Queue) Enqueue(val interface{}) {
	q.list = append(q.list, val)
}

// Dequeue 出列
func (q *Queue) Dequeue() (peek interface{}, success bool) {
	if len(q.list) == 0 {
		success = false
		return
	}

	peek = q.list[0]
	success = true
	q.list = q.list[1:]
	return
}

// Peek 查看队头信息
func (q *Queue) Peek() (peek interface{}, success bool) {
	if len(q.list) == 0 {
		success = false
		return
	}

	peek = q.list[0]
	success = true
	return
}

// Size 获取队列长度
func (q *Queue) Size() int {
	return len(q.list)
}

// Empty 判断队列是否为空
func (q *Queue) Empty() bool {
	if len(q.list) == 0 {
		return true
	}
	return false
}
