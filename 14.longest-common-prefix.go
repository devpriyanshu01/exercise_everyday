/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
   if len(strs) < 2 {
	return strs[0]
   }
   //find the smallest string
   s := 0
   ls := len(strs[0]) 
   for i := 0; i < len(strs); i++ {
	   if len(strs[i]) < ls {
		   s = i
		   ls = len(strs[i])
	   }
   }

   //core logic
   ans := ""
   t := -1
   for i := 0; i < len(strs); i++ {
	   flag := true
	   j := 0
	   for j < len(strs[s]) {
		   if strs[i][j] != strs[s][j] {
			   t = j
			   flag = false
			   break
		   }
	   } 
	   if !flag {
		  ans = strs[s][:t]  
		  break
	   }
   }
   return ans
}
// @lc code=end

