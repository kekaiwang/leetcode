/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	slow, fast := 0, 0

	for fast < n {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}

	return
}

// @lc code=end

