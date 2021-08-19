package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, current *ListNode
	current = head

	// 1 -> 2 -> 3 ->
	// <- 1 <- 2 <- 3

	for current != nil {
		next := current.Next 	// сохранить для перемещения итератора

		current.Next, prev = prev, current	// следующий элемент указывает на предыдущий
		// предыдущий обновляется на текущий

		current = next			// переместить итератор
	}

	return prev
}