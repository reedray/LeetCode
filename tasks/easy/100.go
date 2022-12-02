package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p != nil && q != nil {
		if p.Val == q.Val {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		} else {
			return false
		}
	} else {
		return p == q
	}

}

func main() {

	b := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	c := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	a := &TreeNode{
		Val:   1,
		Left:  b,
		Right: c,
	}
	x := a

	isSameTree(a, x)
}
