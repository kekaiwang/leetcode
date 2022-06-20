/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	mem := make([][]int, n+1)

	for i := range mem {
		mem[i] = make([]int, n+1)
	}

	var count func(low, high int) int

	count = func(low, high int) int {
		if low > high {
			return 1
		}

		if mem[low][high] != 0 {
			return mem[low][high]
		}

		res := 0

		for i := low; i <= high; i++ {
			left := count(low, i-1)
			right := count(i+1, high)

			res += left * right
		}

		mem[low][high] = res

		return res
	}

	return count(1, n)
}

// @lc code=end

