/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{}
	count := []int{}
	queue = append(queue, root)
	count = append(count, 1)

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		depth := count[i]

		if node.Left == nil && node.Right == nil {
			return depth
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			count = append(count, depth+1)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			count = append(count, depth+1)
		}
	}

	return 0
}

// @lc code=end

