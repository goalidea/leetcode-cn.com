package valid

import "container/list"

func isValid(s string) bool {
	l := len(s)

	if l == 0 {
		return true
	}

	if l%2 != 0 {
		return false
	}

	stack := list.New()
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack.PushBack(s[i])
		}

		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if stack.Len() == 0 {
				return false
			}
			v, ok := stack.Back().Value.(byte)
			if !ok {
				panic("get top of stack error")
			}
			if v+1 == s[i] || v+2 == s[i] {
				stack.Remove(stack.Back())
			}
		}
	}

	if stack.Len() == 0 {
		return true
	}

	return false
}
