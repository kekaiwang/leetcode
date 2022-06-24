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
	// -----------------------
	// 后序遍历方法
	sb := &strings.Builder{}

	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("null,")
			return
		}

		dfs(root.Left)
		dfs(root.Right)

		val := strconv.Itoa(root.Val)
		sb.WriteString(val)
		sb.WriteString(",")
	}

	dfs(root)

	str := strings.TrimRight(sb.String(), ",")

	return str

	// -----------------------
	// 前序遍历方法
	// sb := &strings.Builder{}

	// var dfs func(*TreeNode)

	// dfs = func(root *TreeNode) {
	// 	if root == nil {
	// 		sb.WriteString("null,")
	// 		return
	// 	}

	// 	val := strconv.Itoa(root.Val)

	// 	sb.WriteString(val)
	// 	sb.WriteString(",")

	// 	dfs(root.Left)
	// 	dfs(root.Right)
	// }

	// dfs(root)

	// return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// -----------------------
	// 后序遍历方法

	sp := strings.Split(data, ",")
	var build func() *TreeNode

	build = func() *TreeNode {
		if len(sp) == 0 {
			return nil
		}

		last := sp[len(sp)-1]
		sp = sp[:len(sp)-1]
		if last == "null" {
			return nil
		}

		val, _ := strconv.Atoi(last)

		node := &TreeNode{Val: val}

		node.Right = build()
		node.Left = build()

		return node
	}

	return build()

	// -----------------------
	// 前序遍历方法
	// sp := strings.Split(data, ",")
	// var build func() *TreeNode

	// build = func() *TreeNode {
	// 	if sp[0] == "null" {
	// 		sp = sp[1:]
	// 		return nil
	// 	}

	// 	val, _ := strconv.Atoi(sp[0])

	// 	sp = sp[1:]

	// 	return &TreeNode{Val: val, Left: build(), Right: build()}
	// }

	// return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

