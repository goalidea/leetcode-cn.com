package pivotindex

func pivotIndex(nums []int) int {
	var sum, lsum int

	if len(nums) == 0 {
		return -1
	}

	for _, v := range nums {
		sum += v
	}

	for i := range nums {
		if lsum == sum-lsum-nums[i] {
			return i
		}

		lsum += nums[i]
	}

	return -1
}
