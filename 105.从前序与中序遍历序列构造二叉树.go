/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1)
}

func build(preorder []int, preStart, preEnd int,
	inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[preStart]}

	var index int
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}

	leftSize := index - inStart

	root.Left = build(preorder, preStart+1, leftSize+preStart,
		inorder, inStart, index-1)
	root.Right = build(preorder, preStart+leftSize+1, preEnd,
		inorder, index+1, inEnd)

	return root
}

// @lc code=end

