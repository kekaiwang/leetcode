/*
 * @lc app=leetcode.cn id=1038 lang=golang
 *
 * [1038] 从二叉搜索树到更大和树
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
func bstToGst(root *TreeNode) *TreeNode {
	var traverse func(*TreeNode)
	var sum int

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Right)
		sum += root.Val
		root.Val = sum
		traverse(root.Left)
	}

	traverse(root)

	return root
}

// @lc code=end

