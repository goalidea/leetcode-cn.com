package decodestring

import "container/list"

func decodeString(s string) string {
	l := len(s)
	var result string

	countstack, stringstack := list.New(), list.New()

	var count int
	for i := 0; i < l; i++ {
		if s[i] > 47 && s[i] < 58 {

			count = count*10 + int(s[i]-'0')

		} else if s[i] == '[' {

			countstack.PushBack(count)
			stringstack.PushBack(result)

			count = 0
			result = ""
		} else if s[i] == ']' {
			tmpcount := countstack.Remove(countstack.Back()).(int)
			tmpstring := stringstack.Remove(stringstack.Back()).(string)
			for j := 0; j < tmpcount; j++ {
				tmpstring += result
			}
			result = tmpstring
		} else {
			result += string(s[i])
		}
	}

	return result
}
