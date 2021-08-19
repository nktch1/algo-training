package main

func sumBinaryStrings(s1, s2 string) string {

	var (
		r1    = []rune(s1)
		r2    = []rune(s2)
		n     = max(len(r1), len(r2))
		r     = make([]rune, n)
		carry = false
	)

	r1 = fill(r1, n)
	r2 = fill(r2, n)

	for i := n - 1; i >= 0; i-- {
		v := false

		if r1[i] == '1' {
			v = true
		}

		if r1[i] == r2[i] && v { // 1 + 1 + 1 = 1; 1 + 1 + 0 = 0
			r[i] = '0'

			if carry {
				r[i] = '1'
			}

			carry = true

			continue
		}

		if r1[i] == r2[i] && !v { // 0 + 0 + 0 = 0; 0 + 0 + 1 = 1
			r[i] = '0'

			if carry {
				r[i] = '1'
			}

			carry = false

			continue
		}

		if r1[i] != r2[i] && carry { // 0 + 1 + 1 = 0
			r[i] = '0'
			carry = true

			continue
		}

		if r1[i] != r2[i] && !carry { // 0 + 1 + 0 = 1
			r[i] = '1'
			carry = false
		}
	}

	if carry {
		r = append([]rune("1"), r...)
	}

	return string(r)
}

func fill(slice []rune, c int) []rune {
	if len(slice) == c {
		return slice
	}

	r := make([]rune, c-len(slice))
	for i := range r {
		r[i] = '0'
	}

	return append(r, slice...)
}

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}
