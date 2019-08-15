package stack

type Item struct {
	value int
	next  *Item
}

type Stack struct {
	top  *Item
	size int
}

func New() *Stack {
	return &Stack{}

}
func (stack *Stack) Push(value int) {
	stack.size++
	stack.top = &Item{value, stack.top}
}

func (stack *Stack) Pop() int {
	switch {
	case stack.size > 0:
		value := stack.top.value
		stack.top = stack.top.next
		stack.size--
		return value
	default:
		return 0
	}
}
