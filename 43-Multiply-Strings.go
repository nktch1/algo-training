package main

import (
	"fmt"
	"strings"
)

func prepareNums(str string) []int {
	var ints []int

	for _, el := range str {
		ints = append(ints, int(el)-'0')
	}

	return ints
}

func mul(num1 []int, num2 []int) []int {

	var (
		res    = make([]int, len(num1)+len(num2))
		n1, n2 = len(num1), len(num2)
	)

	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			sum := res[i+j+1] + num1[i]*num2[j] // к результату умножения добавили остаток с прошлого шага

			res[i+j] += sum / 10
			res[i+j+1] = sum % 10
		}
	}

	started := false
	startIdx := 0

	for i := 0; res[i] == 0 && !started && i < len(res)-1; i++ {
		if res[i] != 0 {
			started = true
		}

		startIdx++
	}

	return res[startIdx:]
}

func multiply(num1 string, num2 string) string {

	var (
		ints1 = prepareNums(num1)
		ints2 = prepareNums(num2)
		res   = mul(ints1, ints2)
		b     = strings.Builder{}
	)

	for _, el := range res {
		fmt.Fprintf(&b, "%d", el)
	}

	return b.String()
}
