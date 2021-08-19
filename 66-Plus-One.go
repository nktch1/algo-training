package main

func plusOne(digits []int) []int {
	carry := 0

	for i := len(digits)-1; i >= 0; i-- {
		if i == len(digits)-1 || carry > 0 {
			v := (digits[i] + 1) % 10
			carry = (digits[i] + 1) / 10
			digits[i] = v
		}
	}

	if carry > 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
