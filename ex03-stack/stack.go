package stack

type (
	item struct {
		value int
		next  *item
	}

	Stack struct {
		top  *item
		size int
	}
)

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int) {
	s.top = &item{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value int) {
	if s.size > 0 {
		value = s.top.value
		s.top = s.top.next
		s.size--
		return value
	}
	return -1
}
