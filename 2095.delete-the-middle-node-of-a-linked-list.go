/*
 * @lc app=leetcode id=2095 lang=golang
 *
 * [2095] Delete the Middle Node of a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	
    slow, fast := head, head

	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	//if prev == nil => no loop iteration
	if prev == nil {
		return nil
	}
	
	prev.Next = prev.Next.Next

	return head
}
// @lc code=end

