package helper

// Stack 栈
type Stack struct {
	elems []interface{}
}

// NewStack 栈初始化
func NewStack() Stack {
	return Stack{make([]interface{}, 0)}
}

// Push 入栈
func (s *Stack) Push(value interface{}) {
	s.elems = append(s.elems, value)
}

// Pop 出栈
func (s *Stack) Pop() (topElem interface{}) {
	if len(s.elems) == 0 {
		return nil
	}
	topElem = s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return topElem
}

// Empty 判断栈为空
func (s *Stack) Empty() bool {
	if len(s.elems) == 0 {
		return true
	}
	return false
}
