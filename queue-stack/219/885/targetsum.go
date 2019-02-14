package targetsum

func findTargetSumWays(nums []int, S int) int {
	l := len(nums)

	count := new(int)
	*count = 0

	dfs(nums, S, 0, l, count)
	return *count
}

func dfs(nums []int, S int, start int, lnums int, count *int) {
	if start >= lnums {
		if S == 0 {
			*count++
		}
		return
	}

	dfs(nums, S-nums[start], start+1, lnums, count)
	dfs(nums, S+nums[start], start+1, lnums, count)
}
