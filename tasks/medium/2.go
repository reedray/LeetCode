package main

import (
	"LeetCode/models"
	"fmt"
)

func addTwoNumbers(l1 *models.ListNode, l2 *models.ListNode) *models.ListNode {

	sum := 0
	start := new(models.ListNode)
	step := start
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		step.Next = &models.ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		step = step.Next
		sum /= 10
	}
	if sum > 0 {
		step.Next = &models.ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return start.Next
}

func main() {
	c := models.ListNode{
		Val:  3,
		Next: nil,
	}
	b := models.ListNode{
		Val:  4,
		Next: &c,
	}
	a := models.ListNode{
		Val:  2,
		Next: &b,
	}
	g := models.ListNode{
		Val:  4,
		Next: nil,
	}
	f := models.ListNode{
		Val:  4,
		Next: &g,
	}
	e := models.ListNode{
		Val:  6,
		Next: &f,
	}
	d := models.ListNode{
		Val:  5,
		Next: &e,
	}
	numbers := addTwoNumbers(&a, &d)
	printList1(numbers)

}
func printList1(l *models.ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
