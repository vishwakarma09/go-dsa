package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	// Map to store the last seen index of each character
	lastSeen := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, char := range runes {
		// If character was seen after the start of the current window,
		// move the start to the position after the last occurrence.
		if idx, ok := lastSeen[char]; ok && idx >= start {
			start = idx + 1
		}

		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}

		lastSeen[char] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
