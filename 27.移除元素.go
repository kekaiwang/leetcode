/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}

	fast, slow := 0, 0
	for fast < n {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

// @lc code=end

