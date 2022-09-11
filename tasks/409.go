package main

func longestPalindrome(s string) int {
	hm := make(map[uint8]int)
	maxLen := 0
	for i := 0; i < len(s); i++ {
		hm[s[i]]++
	}
	for _, i := range hm {
		maxLen += i / 2 * 2
		if maxLen%2 == 0 && i%2 == 1 {
			maxLen += 1
		}
	}
	return maxLen
}

func main() {
	s := "abccccdd"
	longestPalindrome(s)

}
