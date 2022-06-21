/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}

	var index, max = -1, math.MinInt32
	for i := low; i <= high; i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}

	root := &TreeNode{Val: max}

	root.Left = build(nums, low, index-1)
	root.Right = build(nums, index+1, high)

	return root
}

// @lc code=end

