package main

func strStr(haystack string, needle string) int {
	first, j := 0, 0

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[j] {
			if j == len(needle)-1 {
				first = i - j
				return first
			}
			j++
			continue
		} else {
			i = i - j
			j = 0
		}
	}
	return -1
}
