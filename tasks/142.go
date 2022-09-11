package main

import "LeetCode/models"

func detectCycle(head *models.ListNode) *models.ListNode {

	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next
	var flag bool
	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			flag = true
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	if flag {
		slow = head
		fast = fast.Next
		for fast != slow {
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	}
	return nil
}

func main() {

	y := models.ListNode{
		Val:  7,
		Next: nil,
	}
	z := models.ListNode{
		Val:  6,
		Next: &y,
	}

	x := models.ListNode{
		Val:  5,
		Next: &z,
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
	y.Next = &c
	detectCycle(&a)

}
