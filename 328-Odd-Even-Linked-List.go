package main

import "fmt"

func printList(head *ListNode) {
	iter := head

	for iter != nil {
		fmt.Println(iter.Val)
		iter = iter.Next
	}

	fmt.Println()
}


func oddEvenList(head *ListNode) *ListNode {

	// odd  - нечетный
	// even - четный

	var (
		evenIter, oddIter = &ListNode{}, &ListNode{}
		evenHead, oddHead = evenIter, oddIter
	)

	for idx := 0; head != nil; idx++ {
		if isEven(idx) {
			evenIter.Next = head 		// в лист с четными добавить текущий элемент
			evenIter = evenIter.Next	// сдвинуть итератор листа с четными
			head = head.Next			// сдвинуть итератор основного листа

			continue
		}

		oddIter.Next = head				// в лист с нечетными добавить текущий элемент
		oddIter = oddIter.Next			// сдвинуть итератор листа с нечетными
		head = head.Next				// сдвинуть итератор основного листа
	}

	oddIter.Next = nil
	evenIter.Next = oddHead.Next

	return evenHead.Next
}

func isEven(i int) bool {
	if i % 2 == 0 {
		return true
	}

	return false
}