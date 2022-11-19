package easy

func firstBadVersion(n int) int {
	left, mid := 0, 0
	for left < n {
		mid = left + (n-left)/2
		if isBadVersion(mid) {
			n = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {

}
