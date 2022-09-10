package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil {
		slow = slow.Next
		if fast.Next.Next == nil {
			break
		}
		fast = fast.Next.Next

	}

	return slow
}

func main() {

	x := ListNode{
		Val:  5,
		Next: nil,
	}
	h := ListNode{
		Val:  4,
		Next: &x,
	}
	c := ListNode{
		Val:  3,
		Next: &h,
	}
	b := ListNode{
		Val:  2,
		Next: &c,
	}
	//head1
	a := ListNode{
		Val:  1,
		Next: &b,
	}
	now := time.Now()
	middleNode(&a)
	fmt.Println(time.Since(now))

}
