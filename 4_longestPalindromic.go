package main

import "fmt"

func main() {
	s := "caab"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	longest := s[0:1]
	if n == 1 {
		return longest
	}

	var pal_s1, pal_s2 string

	for i := 0; i < len(s)-1; i++ {
		//奇数
		pal_s1 = findPalindrome(s, i, i, n)
		if len(pal_s1) > len(longest) {
			longest = pal_s1
		}

		//偶数
		pal_s2 = findPalindrome(s, i, i+1, n)
		if len(pal_s2) > len(longest) {
			longest = pal_s2
		}
	}
	return longest
}

func findPalindrome(s string, l int, r int, n int) string {
	pos1, pos2 := l, r

	if l == 0 && r == 1 && s[l] == s[r] {
		return s[l : r+1]
	}

	for pos1 > 0 && pos2 < n-1 && s[pos1] == s[pos2] {
		pos1--
		pos2++
	}

	if s[pos1] == s[pos2] {
		return s[pos1 : pos2+1]
	} else {
		return s[pos1+1 : pos2]
	}
}
