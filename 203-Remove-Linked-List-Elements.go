package main

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}

	iter := head

	for iter != nil && iter.Next != nil {
		if iter.Next.Val == val {
			iter.Next = iter.Next.Next
			continue
		}

		iter = iter.Next
	}

	return head
}