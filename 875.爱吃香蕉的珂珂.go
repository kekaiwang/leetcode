/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	left, right := 0, 0

	for _, pile := range piles {
		right = max(right, pile)
	}

	for left < right {
		mid := left + (right-left)/2

		if mid == 0 {
			return 1
		}

		if f(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func f(nums []int, x int) int {
	hours := 0

	for i := 0; i < len(nums); i++ {
		hours += nums[i] / x
		if nums[i]%x > 0 {
			hours++
		}
	}

	return hours
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

