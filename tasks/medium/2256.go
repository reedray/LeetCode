package main

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
func minimumAverageDifference(nums []int) int {

	dl := len(nums)
	var index int
	max := 100000
	var min, left, right int
	for i := 1; i < dl; i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	for i := 0; i < dl; i++ {
		left = nums[i] / (i + 1)
		if i == dl-1 {
			right = 0
		} else {
			right = (nums[dl-1] - nums[i]) / (dl - i - 1)
		}
		min = abs(left - right)
		if min < max {
			index = i
			max = min
		}
	}
	return index
}

func main() {
	x := []int{4, 2, 0}
	minimumAverageDifference(x)

}
