package temperature

import "container/list"

func dailyTemperatures(T []int) []int {
	l := len(T)
	result := make([]int, l)

	stack := list.New()
	for i := l - 1; i >= 0; i-- {
		for stack.Len() != 0 && T[i] >= T[stack.Back().Value.(int)] {
			stack.Remove(stack.Back())
		}

		if stack.Len() == 0 {
			result[i] = 0
		} else {
			result[i] = stack.Back().Value.(int) - i
		}

		stack.PushBack(i)
	}

	return result
}
