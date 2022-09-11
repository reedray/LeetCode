package main

import (
	"LeetCode/models"
	"fmt"
	"time"
)

func middleNode(head *models.ListNode) *models.ListNode {
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

	x := models.ListNode{
		Val:  5,
		Next: nil,
	}
	h := models.ListNode{
		Val:  4,
		Next: &x,
	}
	c := models.ListNode{
		Val:  3,
		Next: &h,
	}
	b := models.ListNode{
		Val:  2,
		Next: &c,
	}
	//head1
	a := models.ListNode{
		Val:  1,
		Next: &b,
	}
	now := time.Now()
	middleNode(&a)
	fmt.Println(time.Since(now))

}
