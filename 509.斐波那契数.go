/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	if n <= 1 {
		return n
	}

	p, q, r := 0, 0, 1

	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}

	return r
}

// @lc code=end

