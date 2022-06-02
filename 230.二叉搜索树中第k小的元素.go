/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	var traverse func(*TreeNode)

	var res int
	var rank int

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)
		rank++
		if rank == k {
			res = root.Val
			return
		}
		traverse(root.Right)
	}

	traverse(root)

	return res
}

// @lc code=end

