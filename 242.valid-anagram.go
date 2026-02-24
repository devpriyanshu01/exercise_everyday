/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
    smap := make(map[rune]int) 
	
	for _, r := range s {
		smap[r]++
	}
	// fmt.Println(smap)
	
	for _, r := range t {
		_, ok := smap[r]
		if ok {
			smap[r] = smap[r]-1
		}else {
			return false
		}
	}

	for _, v := range smap {
		if v != 0 {
			return false
		}
	}
	return true

}
// @lc code=end

