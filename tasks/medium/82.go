package main

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	cur := head
	prev := dummy
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			prev.Next = cur.Next
		} else {
			prev = prev.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

func main() {

}
