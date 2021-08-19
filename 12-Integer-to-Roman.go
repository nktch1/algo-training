package main

import "strings"

func intToRoman(num int) string {
	m := map[int]string{
		1    : "I",
		4    : "IV",
		5    : "V",
		9    : "IX",
		10   : "X",
		40   : "XL",
		50   : "L",
		90   : "XC",
		100  : "C",
		400  : "CD",
		500  : "D",
		900  : "CM",
		1000 : "M",
	}

	vars := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	var sb strings.Builder

	for num > 0 {
		current := getCurrent(num, vars)

		sb.WriteString(m[current])

		num -= current
	}

	return sb.String()
}

func getCurrent(num int, vars []int) int {
	r := 0

	for i := 0; i < len(vars) && vars[i] <= num; i++ {
		r = vars[i]
	}

	return r
}
