/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	//solution 2
	curr := headA
	cache := make(map[*ListNode]struct{})

	//store first linked list in a cache-map
	for curr != nil {
		cache[curr] = struct{}{}
		curr = curr.Next
	}

	curr = headB

	//check if curr node is present
	for curr != nil {
		_, ok := cache[curr]
		if ok {
			return curr
		}
		curr = curr.Next
	}

	return nil

	//solution 1: O(n^2) time complexity & O(1) space complexity
	/*
    first := headA

	for first != nil {
		second := headB
		for second != nil {
			if first == second {
				return first
			}
			second = second.Next
		}
		first = first.Next
	}
	return nil*/
}
// @lc code=end

