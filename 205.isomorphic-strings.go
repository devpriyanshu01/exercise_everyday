/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
   dict := make(map[string]string) 

   for i, r := range s {
	   dict[string(r)] = string(t[i])
   }

   for k, v := range dict {
	   temp := []string{k, v}
	   for x, y := range dict {
			if y == temp[1] {
				if x != temp[0] {
					return false
				}
			}
	   }
   }

   ans := ""
   for _, r := range s {
	   ans += dict[string(r)]
   }

   if ans == t {
	   return true
   }

   return false
   
}
// @lc code=end

