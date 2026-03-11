/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans *ListNode
	var ansHead *ListNode

	curr1, curr2 := l1, l2

	carry := 0
	for curr1 != nil && curr2 != nil {
		sum := curr1.Val + curr2.Val
		sum += carry

		rem := sum % 10
		carry = sum / 10

		node := &ListNode{
			Val: rem,
		}

		if ans == nil {
			ans = node
			ansHead = ans
		} else {
			ans.Next = node
			ans = ans.Next
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	for curr1 != nil {
		sum := curr1.Val + carry
		val := sum % 10
		carry = sum / 10

		node := &ListNode{
			Val: val,
		}

		ans.Next = node
		ans = ans.Next
		curr1 = curr1.Next
	}

	for curr2 != nil {
		sum := curr2.Val + carry
		val := sum % 10
		carry = sum / 10

		node := &ListNode{
			Val: val,
		}

		ans.Next = node
		ans = ans.Next
		curr2 = curr2.Next
	}

	if carry != 0 {
		ans.Next = &ListNode {
			Val : carry,
		}
	}
	return ansHead
}

// @lc code=end

