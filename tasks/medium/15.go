package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
					continue
				}
			}
		}
	}
	return result
}

func main() {

	x := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(x))

}
