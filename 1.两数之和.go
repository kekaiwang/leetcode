/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for key, val := range nums {
		if v, ok := m[target-val]; ok {
			return []int{v, key}
		}

		m[val] = key
	}

	return nil
}

// @lc code=end

