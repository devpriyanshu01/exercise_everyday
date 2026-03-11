/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	out := head

	for out.Next != nil {
		in := out.Next 
			for in != nil {
				if out.Val > in.Val {
					swap(out, in)
				}
				in = in.Next
			}
		out = out.Next
	}
	return head
}



func swap(out *ListNode, in *ListNode) {
	temp := out.Val 
	out.Val = in.Val
	in.Val = temp
}
// @lc code=end

