/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	var res []*TreeNode
	var dfs func(*TreeNode) string
	subTreeMap := map[string]int{}

	dfs = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}

		subTree := fmt.Sprintf("%v:%v:%v", root.Val, dfs(root.Left), dfs(root.Right))

		subTreeMap[subTree]++

		if subTreeMap[subTree] == 2 {
			res = append(res, root)
		}

		return subTree
	}

	dfs(root)

	return res
}

// @lc code=end

