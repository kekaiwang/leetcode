/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return nil
	}

	return tree(1, n)
}

func tree(low, high int) []*TreeNode {
	if low > high {
		return []*TreeNode{nil}
	}

	var res []*TreeNode

	for i := low; i <= high; i++ {
		left := tree(low, i-1)
		right := tree(i+1, high)

		for _, l := range left {
			for _, r := range right 
				root := &TreeNode{Val: i}
				root.Left = l
				root.Right = r
				res = append(res, root)
			}
		}
	}

	return res·
}

// @lc code=end

