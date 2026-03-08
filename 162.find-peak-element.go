/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {
	//Below is Brute Force code in O(n)
	/*
	n := len(nums)

	if n == 1 {
		return 0
	}
   
   for i := 1; i < n-1; i++ {
	   if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
		   return i
	   }
   }

   if nums[0] > nums[1] && nums[0] >= nums[n-1] {
	   return 0
   }

   if nums[n-1] > nums[n-2] && nums[n-1] > nums[0] {
	   return n-1
   }

   return -1 */


   //OPTIMAL APPROACH: 
   n := len(nums)
   
   if n == 1 {
	    return 0
   }

   if nums[0] > nums[1] {
	    return 0
   }else if nums[n-1] > nums[n-2] {
	    return n-1
   }

   s, e := 1, n-2
   m := (s+e)/2

   for s <= e {
	   if nums[m] > nums[m-1] && nums[m] > nums[m+1] {
		   return m
	   }else if nums[m] > nums[m-1] && nums[m] < nums[m+1] {
		   s = m+1
	   }else {
		   e = m-1
	   }
	   m = (s+e)/2
   }
   return -1
}
// @lc code=end

