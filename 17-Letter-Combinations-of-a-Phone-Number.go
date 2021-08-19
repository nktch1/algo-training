package main

import "fmt"

func letterCombinations(digits string) []string {
	//m := map[string]string{}{}

	s := []rune{}
	r := []string{}

	for _, el := range digits {
		switch el {
		case '2':
			s = append(s, []rune{'a', 'b', 'c'}...)
		case '3':
			s = append(s, []rune{'d', 'e', 'f'}...)
		case '4':
			s = append(s, []rune{'g', 'h', 'i'}...)
		case '5':
			s = append(s, []rune{'j', 'k', 'l'}...)
		case '6':
			s = append(s, []rune{'m', 'n', 'o'}...)
		case '7':
			s = append(s, []rune{'p', 'q', 'r', 's'}...)
		case '8':
			s = append(s, []rune{'t', 'u', 'v'}...)
		case '9':
			s = append(s, []rune{'w', 'x', 'y', 'z'}...)
		}
	}

	fmt.Println(s)

	return []string{}
}
