/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {

	s = strings.TrimSpace(s)

	words := strings.Split(s, " ")


	//remove empty strings
	cleaned := []string{}
	for _, word := range words {
		if word != "" {
			cleaned = append(cleaned, word)
		}
	}

	ans := ""
	for i := len(cleaned)-1; i >= 0; i-- {
		ans += cleaned[i] + " "
	}
	return strings.TrimSpace(ans)
}
// @lc code=end

