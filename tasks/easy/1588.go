package easy

func sumOddLengthSubarrays(arr []int) int {

	var sum int
	dl := len(arr)
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i] + arr[i-1]
	}
	sum = arr[dl-1]
	for i := 2; i < dl; i += 2 {
		for j, z := i, 0; j < dl; j++ {
			if j == i {
				sum += arr[j]
				continue
			}
			sum += arr[j] - arr[z]
			z++
		}

	}

	return sum
}
