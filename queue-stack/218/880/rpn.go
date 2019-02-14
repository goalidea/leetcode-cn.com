package rpn

import (
	"container/list"
	"strconv"
)

func evalRPN(tokens []string) int {
	l := len(tokens)
	var result int

	stack := list.New()
	for i := 0; i < l; i++ {
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			if v, err := strconv.Atoi(tokens[i]); err != nil {
				panic(err)
			} else {
				stack.PushBack(v)
			}
		}
		if tokens[i] == "+" {
			t1 := stack.Remove(stack.Back()).(int)
			t2 := stack.Remove(stack.Back()).(int)
			t := t2 + t1
			stack.PushBack(t)
		}
		if tokens[i] == "-" {
			t1 := stack.Remove(stack.Back()).(int)
			t2 := stack.Remove(stack.Back()).(int)
			t := t2 - t1
			stack.PushBack(t)
		}
		if tokens[i] == "*" {
			t1 := stack.Remove(stack.Back()).(int)
			t2 := stack.Remove(stack.Back()).(int)
			t := t2 * t1
			stack.PushBack(t)
		}
		if tokens[i] == "/" {
			t1 := stack.Remove(stack.Back()).(int)
			t2 := stack.Remove(stack.Back()).(int)
			t := t2 / t1
			stack.PushBack(t)
		}
	}
	if stack.Len() == 1 {
		result = stack.Remove(stack.Back()).(int)
	}
	return result
}
