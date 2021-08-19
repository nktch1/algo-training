package main

func isAnagram(s string, t string) bool {
	m := map[rune]int{}

	for _, el := range s {
		m[el]++
	}

	for _, el := range t {
		m[el]--
		if m[el] == 0 {
			delete(m, el)
		}
	}

	return len(m) == 0
}
