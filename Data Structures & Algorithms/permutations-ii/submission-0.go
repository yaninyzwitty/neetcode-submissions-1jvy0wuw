func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var path []int
	used := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			perm := make([]int, len(path))
			copy(perm, path)
			result = append(result, perm)
			return

		}

		for i := range nums {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// choose
			used[i] = true
			path = append(path, nums[i])
			dfs()

			// unchoose
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return result


}
