/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    curr := node
	prev := node

	for curr.Next != nil {
		curr.Val = curr.Next.Val
		prev = curr
		curr = curr.Next
	}

	prev.Next = nil
}
// @lc code=end

