/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	// 当当前数组 <= 1 时返回即可
	n := len(nums)
	if n <= 1 {
		return n
	}

	// 使用快慢指针进行修改数据
	slow, fast := 0, 0
	// 快指针走的快同时要比数组长度小
	for fast < n {
		if nums[slow] != nums[fast] {
			// 将 slow 加一后复制 fast 到 slow 处
			slow++
			nums[slow] = nums[fast]
		}

		fast++
	}

	return slow + 1
}

// @lc code=end

