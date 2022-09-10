package main

import (
	"fmt"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	list := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return list

}

func reverseListIter(head *ListNode) *ListNode {

	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
func main() {
	c := ListNode{
		Val:  3,
		Next: nil,
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

	twoLists := reverseListIter(&a)
	printList(twoLists)

}

func printList(x *ListNode) {
	for x != nil {
		fmt.Println(x.Val)
		x = x.Next
	}
}
