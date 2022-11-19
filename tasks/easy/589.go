package easy

import (
	"LeetCode/models"
)

func preorder(root *models.Node) []int {

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

}
