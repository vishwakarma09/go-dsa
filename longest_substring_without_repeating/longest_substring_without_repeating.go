package main

import (
	"fmt"
)

func main() {
	maxLen := lengthOfLongestSubstringSlidingWindow("abcabbd")
	fmt.Printf("max len of non repeating substring is %v", maxLen)
}

func lengthOfLongestSubstringSlidingWindow(s string) int {
	seen := make(map[byte]bool)
	maxLength, left := 0, 0

	for right := 0; right < len(s); right++ {
		for seen[s[right]] {
			delete(seen, s[left])
			left++
		}
		seen[s[right]] = true
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}
