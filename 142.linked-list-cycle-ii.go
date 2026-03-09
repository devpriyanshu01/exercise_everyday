/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    /*nodeElement := make(map[*ListNode]struct{})
	
	curr := head

	for curr != nil {
		_, ok := nodeElement[curr]
		if ok {
			return curr
		}
		nodeElement[curr] = struct{}{}
		curr = curr.Next
	}
	return nil */

	slow, fast := head, head

	for slow.Next != nil && fast.Next != nil {
		if slow == fast {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	if slow.Next == nil || fast.Next == nil {
		return nil
	}else {
		//step 2.
		slow = head

		for slow != fast {
			slow = slow.Next
			fast = fast.Next.Next	
		}

		return slow
	}

	return nil
	
}
// @lc code=end

