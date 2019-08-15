package brackets

type Item struct {
	value uint8
	next  *Item
}

type Stack struct {
	top  *Item
	size int
}

func New() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(value uint8) {
	stack.size++
	stack.top = &Item{value, stack.top}
}

func (stack *Stack) Pop() uint8 {
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

func Bracket(str string) (bool, error) {
	stack := New()
	var character uint8
	var character2 uint8
	length := len(str)
	for i := 0; i < length; i++ {
		if str[i] == 123 || str[i] == 91 || str[i] == 40 {
			stack.Push(str[i])

		} else {
			character = stack.Pop()
			if character == 40 {
				character2 = str[i] - 1
			} else {
				character2 = str[i] - 2
			}
			if character != character2 {
				return false, nil
			}
		}
	}
	if stack.size != 0 {
		return false, nil
	} else {
		return true, nil
	}
}
