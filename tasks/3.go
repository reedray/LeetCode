package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {

	left, right, l := 0, 0, 0
	hm := make(map[uint8]int)
	for ; right < len(s); right++ {
		hm[s[right]]++
		for hm[s[right]] > 1 {
			hm[s[left]]--
			left++
		}
		l = max(l, right-left+1)
	}
	return l
}

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2

}

func main() {
	s := "pwwkew"
	palindrome := lengthOfLongestSubstring(s)
	fmt.Println(palindrome)

}
