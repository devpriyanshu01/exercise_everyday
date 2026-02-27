/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */


// @lc code=start
func longestPalindrome(s string) string {
    ans := ""
	ansLen := 0

	n := len(s)

	for i := 0; i < len(s); i++ {
		//for odd length palindromes
		l, r := i, i
		
		for l >= 0 && r < n && s[l] == s[r] {
			if len(s[l:r+1]) > ansLen {
				ansLen = len(s[l:r+1])
				ans = s[l:r+1]
			}
			l--
			r++
		}

		//for even length palindromes
		l, r = i-1, i
		for l >= 0 && r < n && s[l] == s[r] {
			if len(s[l:r+1]) > ansLen {
				ansLen = len(s[l:r+1])
				ans = s[l:r+1]
			}
			l--
			r++
		}

	}

	
	return ans
}

// func isPalindrome(str string) bool {
// 	i, j := 0, len(str)-1

// 	for j > i {
// 		if str[i] != str[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }
// @lc code=end

