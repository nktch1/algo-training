package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoListsRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)

	return l2
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		last = &ListNode{}
		head = last
	)

	for l1 != nil && l2 != nil { // итерации пока не кончится более короткий список
		if l1.Val < l2.Val {
			last.Next = l1 	// запомнить значение
			l1 = l1.Next 	// сдвинуть указатель на следующий элемент первого списка
		} else {
			last.Next = l2
			l2 = l2.Next
		}

		last = last.Next	// движение по новому списку
	}

	if l1 != nil {			// на данном этапе остался один список, в котором есть элементы
		last.Next = l1		// нужно найти его и добавить весь остаток списка в результат
	} else {
		last.Next = l2
	}

	return head.Next
}
