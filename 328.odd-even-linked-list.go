/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	//handle when only one element
    if head == nil || head.Next == nil {
		return head
	}

	//core logic
	odd, even, evenStart := head, head.Next, head.Next

	for odd != nil && odd.Next != nil && even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}

	odd.Next = evenStart
	return head
}
// @lc code=end

