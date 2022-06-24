/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
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

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}

	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("null,")
			return
		}

		val := strconv.Itoa(root.Val)

		sb.WriteString(val)
		sb.WriteString(",")

		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)

	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode

	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}

		val, _ := strconv.Atoi(sp[0])

		sp = sp[1:]

		return &TreeNode{Val: val, Left: build(), Right: build()}
	}

	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

