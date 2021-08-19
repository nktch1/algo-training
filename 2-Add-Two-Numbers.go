package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		out   = &ListNode{}
		i     = out
		carry int
	)

	for {
		if l1 == nil && l2 == nil {
			break
		}

		tmp := carry

		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}

		i.Next = &ListNode{
			Val: tmp % 10,
		}

		i = i.Next

		carry = tmp / 10
	}

	if carry == 0 {
		return out.Next
	}

	i.Next = &ListNode{
		Val: carry,
	}

	return out.Next
}
