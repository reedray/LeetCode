package easy

func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
func main() {

	x := []int{1, 1, 2, 2, 1, 3, 5}
	removeDuplicates(x)
}
