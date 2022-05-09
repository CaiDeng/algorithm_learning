package helper

// Stack 栈
type Stack struct {
	elems []int
}

func (s *Stack) push(value int) {
	s.elems = append(s.elems, value)
}

func (s *Stack) pop() (topElem int) {
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
