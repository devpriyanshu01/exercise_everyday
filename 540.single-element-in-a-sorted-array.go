/*
 * @lc app=leetcode id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 */

// @lc code=start
func singleNonDuplicate(nums []int) int {
	//below is implementation using xor - O(n)
	/*
		n := len(nums)
		if n == 1 {
			return nums[0]
		}

	    xor := nums[0]

		for i := 1; i < n; i++ {
			xor = xor ^ nums[i]
		}
		return xor
	*/

	//better approach using binary search
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	//check if first & last index
	if nums[0] != nums[1] {
		return nums[0]
	} else if nums[n-1] != nums[n-2] {
		return nums[n-1]
	} else {
		//for i = 1 to i = n-2
		s, e := 1, n-2

		m := (s + e) / 2

		for s <= e {
			if nums[m] != nums[m-1] && nums[m] != nums[m+1] {
				return nums[m]
			} else if m&1 == 1 {
				//m is at odd index
				if nums[m] == nums[m-1] {
					//we're on left part. Ans lies on right part
					s = m + 1
				} else {
					//we're on right part, Ans lies on left part
					e = m - 1
				}
			} else {
				//m is at even index
				if nums[m] == nums[m-1] {
					//we're on right part => Ans lies on left, So eliminate right part
					e = m - 1
				} else {
					//we're on left part => Ans lies on right part
					s = m + 1
				}
			}
			m = (s + e) / 2
		}
	}
	return -1

}

// @lc code=end

