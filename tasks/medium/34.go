package main

import "fmt"

func searchRange(nums []int, target int) []int {

	left := helper(nums, target, true)
	right := helper(nums, target, false)
	return []int{left, right}
}
func helper(nums []int, target int, isLeft bool) int {
	res := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			res = mid
			if isLeft {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return res
}

func main() {

	x := []int{8, 8, 8, 8, 8, 8, 9}
	fmt.Println(x[3:5])
	searchRange(x, 8)

}
