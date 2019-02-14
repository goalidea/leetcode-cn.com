package plusone

func plusOne(digits []int) []int {
	l, tmp := len(digits), []int{}
	for l > 0 {
		if r := digits[l-1]; r+1 < 10 {
			digits[l-1]++
			tmp = append(tmp, digits...)
			break
		} else {
			digits[l-1] = 0
			if l == 1 {
				tmp = append(tmp, 1)
				tmp = append(tmp, digits...)
				break
			} else {
				l--
			}
		}
	}
	return tmp
}
