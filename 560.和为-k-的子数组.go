/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}

	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}

	return count
	// count := 0
	// for start := 0; start < len(nums); start++ {
	// 	sum := 0
	// 	for end := start; end >= 0; end-- {
	// 		sum += nums[end]
	// 		if sum == k {
	// 			count++
	// 		}
	// 	}
	// }
	// return count
}

// @lc code=end

