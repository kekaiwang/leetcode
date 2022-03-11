/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}

	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false) - 1

	if left <= right && right < len(nums) &&
		nums[left] == target && nums[right] == target {
		res = []int{left, right}
	}

	return res
}

func binarySearch(nums []int, target int, lower bool) int {
	left, right := 0, len(nums)-1
	ans := len(nums)

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target || (lower && nums[mid] >= target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}

	return ans
}

// @lc code=end

