package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {

	if root == nil {
		return []int{}
	}
	out := []int{root.Val}
	if root.Children != nil {
		for _, child := range root.Children {
			out = append(out, preorder(child)...)
		}
	}
	return out
}

func main() {

	x := []int{}
	y := []int{1, 2, 3}
	x = append(x, y...)
	fmt.Println(x)
}
