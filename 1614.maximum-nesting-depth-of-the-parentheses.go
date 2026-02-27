/*
 * @lc app=leetcode id=1614 lang=golang
 *
 * [1614] Maximum Nesting Depth of the Parentheses
 */

// @lc code=start
func maxDepth(s string) int {
    count := 0
	max := 0
	
	for _, v := range s {
		if v == '(' {
			count++
			if count > max {
				max = count
			}
		}else if v == ')' {
			count--
		}
	}
	return max
}
// @lc code=end

