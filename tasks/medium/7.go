package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {

	num, result := 0, 0
	for x != 0 {
		num = x % 10
		x /= 10
		result = result*10 + num
	}
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
}

func main() {

	if 9646324351 < math.Pow(2, 31) {
		fmt.Println("less")
	}
	fmt.Println(math.Pow(2, 31))
}
