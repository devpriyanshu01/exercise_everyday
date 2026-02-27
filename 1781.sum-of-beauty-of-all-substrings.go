/*
 * @lc app=leetcode id=1781 lang=golang
 *
 * [1781] Sum of Beauty of All Substrings
 */

// @lc code=start
func beautySum(s string) int {
   beautySum := 0
   
   n := len(s)
   
   dict := make(map[byte]int)
   for i := 0; i < n; i++ {
	   for j := i; j < n; j++ {
		   dict[s[j]]++
		   
		   //calcualte beauty
		   beautySum += beautyCount(dict)
		   
	   }
	   //empty the map
	   clear(dict)
   } 
   return beautySum
}

//calculate beauty count for each substring
func beautyCount(dict map[byte]int) int {
	max := math.MinInt
	min := math.MaxInt

	for _, v := range dict {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max-min
}
// @lc code=end

