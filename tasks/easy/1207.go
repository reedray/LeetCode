package easy

func uniqueOccurrences(arr []int) bool {
	hm := make(map[int]int)
	hm2 := make(map[int]int)
	for _, val := range arr {
		hm[val]++
	}
	for _, val := range hm {
		if _, ok := hm2[val]; !ok {
			hm2[val]++
		} else {
			return false
		}
	}
	return true
}
func main() {

}
