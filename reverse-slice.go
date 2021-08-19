package main

func rev(ints []int) {
	i := 0
	j := len(ints) - 1
	for i < j {
		ints[i], ints[j] = ints[j], ints[i]
		i++
		j--
	}
}
