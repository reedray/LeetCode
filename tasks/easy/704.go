package easy

func search(nums []int, target int) int {

	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	x := []int{-1, 0, 3, 5, 9, 12}

	search(x, 9)
}
