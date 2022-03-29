/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)

	for i := 0; i < n; i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for _, v := range dp {
		if v > res {
			res = v
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

