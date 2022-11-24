package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	first := head
	second := head.Next
	dummy = second.Next
	second.Next = first
	first.Next = swapPairs(dummy)
	return second

}

func main() {

	b := ListNode{
		Val:  4,
		Next: nil,
	}
	a := ListNode{
		Val:  3,
		Next: &b,
	}

	e := ListNode{
		Val:  2,
		Next: &a,
	}
	d := ListNode{
		Val:  1,
		Next: &e,
	}

	swapPairs(&d)

}
