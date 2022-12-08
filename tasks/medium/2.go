package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sum := 0
	start := new(ListNode)
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
		step.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		step = step.Next
		sum /= 10
	}
	if sum > 0 {
		step.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return start.Next
}
