package main

func removeElement(nums []int, val int) int {

	j := 0
	counter := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			counter++
			continue
		}
		nums[j] = nums[i]
		j++
	}
	return len(nums) - counter
}

func main() {

}
