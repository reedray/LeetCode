package easy

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head
	step := head.Next
	for step != nil {
		if step.Val == head.Val {
			head.Next = step.Next
			step = step.Next
		} else {
			head = head.Next
			step = step.Next
		}
	}
	return res
}

func main() {

}
