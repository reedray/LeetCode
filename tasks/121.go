package main

import "fmt"

func maxProfit(arr []int) int {

	maxProf := 0
	for i, j := 0, 1; j < len(arr); j++ {
		if arr[j] < arr[i] {
			i = j
			continue
		} else if arr[j]-arr[i] > maxProf {
			maxProf = arr[j] - arr[i]
		}
	}
	return maxProf
}

func main() {
	x := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(x))

}
