/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
   dict := make(map[string]int) 
   dict["I"] = 1
   dict["X"] = 10
   dict["L"] = 50
   dict["C"] = 100
   dict["V"] = 5
   dict["D"] = 500
   dict["M"] = 1000
   dict["IV"] = 4
   dict["IX"] = 9
   dict["XL"] = 40
   dict["XC"] = 90
   dict["CD"] = 400
   dict["CM"] = 900

   i := len(s)-1
   sum := 0
   var curr string
   for i > 0 {
		curr = s[i-1:i+1]
	   
	   value, ok := dict[curr]
	   
	   if ok {
		   sum += value
		   i -= 2
	   }else {
		   sum += dict[string(s[i])]
		   i--
	   }
   }
   if i == 0 {
	   sum += dict[string(s[i])]
   }
   return sum
}
// @lc code=end

