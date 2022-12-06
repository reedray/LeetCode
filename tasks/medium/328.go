package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	var lastOdd *ListNode
	var prelastOdd *ListNode
	var lastEven *ListNode
	odd := head
	even := head.Next
	firstEven := head.Next
	for odd != nil && odd.Next != nil {
		odd.Next = odd.Next.Next
		prelastOdd = odd
		odd = odd.Next
		lastOdd = odd
		if even.Next != nil {
			even.Next = even.Next.Next
			lastEven = even
			even = even.Next
		} else {
			lastEven = even
		}
	}
	if lastOdd == nil {
		lastOdd = prelastOdd
	}
	lastOdd.Next = firstEven
	lastEven.Next = nil
	return head
}
