package easy

import (
	"fmt"
	"strings"
)

func halvesAreAlike(s string) bool {

	vowels := map[string]int{
		"a": 0,
		"e": 0,
		"i": 0,
		"o": 0,
		"u": 0,
	}
	lower := strings.ToLower(s)
	a := lower[:len(lower)/2]
	b := lower[len(lower)/2:]
	fmt.Println(a, b)
	aCounter, bCounter := 0, 0
	for i := 0; i < len(a); i++ {
		if _, ok := vowels[string(a[i])]; ok {
			aCounter++
		}
		if _, ok := vowels[string(b[i])]; ok {
			bCounter++
		}

	}
	if aCounter == bCounter {
		return true
	}
	return false
}
func main() {
	str := "book"
	halvesAreAlike(str)
}
