/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
func frequencySort(s string) string {
   dict := make(map[rune]int) 

   for _, r := range s {
       dict[r]++
   }

  var builder strings.Builder 

   //iterate to form answer
   for len(dict) != 0 {
	   max := getMaxFreq(dict) 

	   ans := ""
	   for k, v := range dict {
		   if v == max {
			   for i := 0; i < max; i++ {
				   ans += string(k)
			   }
			   delete(dict, k)
		   }
	   }
	   slices.Sort([]rune(ans))
	   builder.WriteString(ans) 
   }
   return builder.String()
}

func getMaxFreq(dict map[rune]int) int {
	max := 0
	for _, v := range dict {
		if v > max {
			max = v
		}
	}
	return max
}
// @lc code=end

