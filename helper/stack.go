package helper

// Stack 栈
type Stack struct {
	elems []interface{}
}

func (s *Stack) push(value interface{}) {
	s.elems = append(s.elems, value)
}

func (s *Stack) pop() (topElem interface{}) {
	if len(s.elems) == 0 {
		return -1
	}
	topElem = s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return topElem
}

func (s *Stack) isEmpty() bool {
	if len(s.elems) == 0 {
		return true
	}
	return false
}
