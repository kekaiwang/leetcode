/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 */

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1002)
	sum := 0

	for _, v := range trips {
		diff[v[1]] += v[0]
		diff[v[2]] -= v[0]
	}

	for i := 0; i < len(diff); i++ {
		sum += diff[i]
		if sum > capacity {
			return false
		}
	}

	return true
}

// @lc code=end

