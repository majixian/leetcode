/*
Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "uqinntq"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	max, count := 1, 1
	var (
		index_begin   int
		index_repeat  int
		curmax_string string
		cur           byte
	)

	if len(s) == 0 {
		return 0
	}

	index_begin = 0

	for i := 1; i < len(s); i++ {
		cur = s[i]
		curmax_string = s[index_begin : count+index_begin]

		index_repeat = strings.IndexByte(curmax_string, cur)

		if index_repeat == -1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = count - index_repeat
			index_begin = index_begin + index_repeat + 1
		}
	}
	return max
}
