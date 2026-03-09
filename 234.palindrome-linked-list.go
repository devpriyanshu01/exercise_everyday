/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    //step 1. find middle element
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow
	//step 2. reverse the linked list from middle element
	var prev *ListNode

	curr := mid
	for curr != nil {
		temp := curr
		curr = curr.Next
		temp.Next = prev
		prev = temp
	}

	//check pointers from head & mid. If they are all equal 
	left, right := head, prev

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}

	return true

}
// @lc code=end

