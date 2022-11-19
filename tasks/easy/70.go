package easy

func climbStairs(n int) int {
	stepL, stepR := 1, 2
	counter := 1
	for counter < n {
		stepL, stepR = stepR, stepL+stepR
		counter++
	}
	return stepL

}

func main() {
	climbStairs(2)
}
