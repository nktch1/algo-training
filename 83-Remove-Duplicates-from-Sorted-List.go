package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		iter = head
		prev = iter
	)

	iter = iter.Next

	for iter != nil {
		if iter.Val == prev.Val {
			prev.Next = iter.Next
		} else {
			prev = iter
		}

		iter = iter.Next
	}

	return head
}
