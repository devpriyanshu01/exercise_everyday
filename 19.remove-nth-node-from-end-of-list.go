/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//handle when head.Next == nil
	if head == nil || head.Next == nil {
		return nil
	}
    size := 0
	curr := head

	for curr != nil {
		curr = curr.Next
		size++
	}

	delNode := size-n
	if delNode == 0 {
		return head.Next
	}

	count := 0
	curr = head
	for curr != nil {
		count++
		if count == delNode {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}
	return head
}
// @lc code=end

