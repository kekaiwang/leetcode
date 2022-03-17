/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	connectNode(root.Left, root.Right)

	return root
}

func connectNode(node1, node2 *Node) {
	if node1 == nil && node2 == nil {
		return
	}

	if node1 == nil {
		node1.Next = nil
		return
	}

	if node2 == nil {
		node2.Next = nil
		return
	}

	node1.Next = node2

	connectNode(node1.Left, node1.Right)
	connectNode(node2.Left, node2.Right)
	connectNode(node1.Right, node2.Left)
}

// @lc code=end

