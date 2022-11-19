package easy

import (
	"LeetCode/models"
	"fmt"
)

func reverseList(head *models.ListNode) *models.ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	list := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return list

}

func reverseListIter(head *models.ListNode) *models.ListNode {

	var prev *models.ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
func main() {
	c := models.ListNode{
		Val:  3,
		Next: nil,
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

	twoLists := reverseListIter(&a)
	printList(twoLists)

}

func printList(x *models.ListNode) {
	for x != nil {
		fmt.Println(x.Val)
		x = x.Next
	}
}
