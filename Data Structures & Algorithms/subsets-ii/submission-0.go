func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var path []int
	var dfs func(start int)
	dfs = func(start int) {
		subset := make([]int, len(path))
		copy(subset, path)
		result = append(result, subset)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// choose
			path = append(path, nums[i])

			// recurse
			dfs(i+1)

			// backtrack
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result

}
