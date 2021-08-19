package main

func isPalindrome(head *ListNode) bool {

	var fast, slow = head, head

	// позволяет найти mid, быстрый указатель успевает
	// дойти до конца списка, когда медленный оказывается в середине
	// slow и указывает на середину
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 1 -> 2 -> 2 -> 1
	// slow -> 2 -> 1

	var (
		current = slow
		prev *ListNode
	)

	// чтобы решить за константу по памяти
	// нужно развернуть вторую половину списка (после точки mid)
	for current != nil {
		tmp := current.Next
		current.Next, prev = prev, current
		current = tmp
	}

	var left, right = head, prev

	// одновременные итерации по двум спискам
	// если элементы не совпадут, вернуть false
	for right != nil {
		if left.Val != right.Val {
			return false
		}

		left = left.Next
		right = right.Next
	}

	// если не нашлось двух разных элементов в списках
	// то вернуть true
	return true
}
