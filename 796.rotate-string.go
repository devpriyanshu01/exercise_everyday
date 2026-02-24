/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 */

// @lc code=start
func rotateString(s string, goal string) bool {
   checker := s

   for i := 0; i < len(s); i++ {
	   checker = checker[1:] + string(checker[0])
	   fmt.Println("checker string: ", checker)
	   if checker == goal {
		   return true
	   }
   }
   return false
}
// @lc code=end

