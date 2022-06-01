/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)

	return newHead
}

func reverse(head *ListNode, last *ListNode) *ListNode {
	var pre *ListNode
	cur, nxt := head, head

	for cur != last {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	return pre
}

// @lc code=end

