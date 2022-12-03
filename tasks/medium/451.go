package main

import (
	"fmt"
	"sort"
	"strings"
)

func frequencySort(s string) string {
	hm := make(map[int32]int, 0)
	for _, val := range s {
		hm[val]++
	}
	keys := make([]int32, 0, len(hm))
	for key := range hm {
		keys = append(keys, key)
	}
	fmt.Println(hm, keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return hm[keys[i]] > hm[keys[j]]
	})
	builder := strings.Builder{}
	for _, val := range keys {
		for i := 0; i < hm[val]; i++ {
			builder.WriteString(string(val))
		}
	}

	return builder.String()
}

func main() {
	frequencySort("aaabbc")

}
