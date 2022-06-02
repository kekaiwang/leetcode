/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
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

