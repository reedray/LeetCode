package main

import "fmt"

func twoSum(nums []int, target int) []int {

	hm := make(map[int]int)
	for i, num := range nums {
		if _, ok := hm[target-num]; !ok {
			hm[num] = i
		} else {
			return []int{hm[target-num], i}
		}
	}
	return nil
}

func main() {
	x := []int{3, 2, 4}
	fmt.Println(twoSum(x, 6))

}
