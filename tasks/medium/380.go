package main

import "math/rand"

type RandomizedSet struct {
	hm  map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hm:  make(map[int]int, 0),
		arr: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hm[val]; !ok {
		this.hm[val] = len(this.arr)
		this.arr = append(this.arr, val)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {

	index, ok := this.hm[val]
	if !ok {
		return false
	}
	last := this.arr[len(this.arr)-1]
	if last == val {
		this.arr = this.arr[:len(this.arr)-1]
	} else {
		this.arr[index] = last
		this.arr = this.arr[:len(this.arr)-1]
		this.hm[last] = index
	}
	delete(this.hm, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.arr))
	return this.arr[index]

}
func main() {

}
