/*
 * @lc app=leetcode id=1781 lang=golang
 *
 * [1781] Sum of Beauty of All Substrings
 */

// @lc code=start
func beautySum(s string) int {
   beautySum := 0
   
   n := len(s)
   
   dict := make(map[rune]int)
   max := 0
   min := 0
   for i := 0; i < n; i++ {
	   for j := i; j < n; j++ {
		   sub := s[i:j+1]
		   dict[s[j]]++
		   
		   if dict[s[j]] > max {
			    max = dict[s[j]]
		   }else {
			    min = dict[s[j]]
		   }

		   beautySum += max-min
	   }
   } 
}
// @lc code=end

