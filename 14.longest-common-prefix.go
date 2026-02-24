/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
    s := 0
	ls := len(strs[0]) 
	for i, str := range strs {
		if len(strs[i]) < ls {
			ls = len(str)
			s = i 
		}
	}
	fmt.Println("value of smallest index s: ", s)

	//core logic
	for i := 0; i < len(strs[s]); i++ {
		flag := true
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != strs[s][i] {
				flag = false
				break
			}
		}
		if !flag {
			return strs[s][:i]
		}
		if i == len(strs[s])-1 {
			return strs[s]
		}
	}
	return ""
}
// @lc code=end

