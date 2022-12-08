package easy

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
