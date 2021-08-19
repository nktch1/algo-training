package main

func removeDuplicates(s string) string {
	var stack []rune

	for _, sym := range s {
		if len(stack) == 0 {
			stack = append(stack, sym)
			continue
		}

		if stack[len(stack)-1] == sym {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, sym)
	}

	return string(stack)
}
