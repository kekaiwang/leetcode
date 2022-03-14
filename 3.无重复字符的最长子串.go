/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)

	right, ans := -1, 0

	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}

		ans = max(ans, right-i+1)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

