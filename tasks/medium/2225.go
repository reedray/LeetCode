package main

import (
	"sort"
)

func findWinners(matches [][]int) [][]int {

	res := make([][]int, 2)
	for i := range res {
		res[i] = make([]int, 0)
	}

	winnersLoosers := make(map[int]int, 0)
	for _, match := range matches {
		winnersLoosers[match[0]] += 0
		winnersLoosers[match[1]]++
	}

	for k, v := range winnersLoosers {
		if v == 0 {
			res[0] = append(res[0], k)

		}
		if v == 1 {
			res[1] = append(res[1], k)
		}

	}

	sort.Ints(res[0])
	sort.Ints(res[1])
	return res
}
