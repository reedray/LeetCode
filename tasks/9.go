package main

import "strconv"

func isPalindrome(x int) bool {
	itoa := strconv.Itoa(x)
	if itoa[0] == '-' {
		return false
	}
	for i := 0; i < len(itoa); i++ {
		if itoa[i] != itoa[len(itoa)-i-1] {
			return false

		}
	}
	return true
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverse := 0
	for x > reverse {

		reverse = reverse*10 + x%10
		x /= 10

	}
	return reverse == x || x == reverse/10

}
