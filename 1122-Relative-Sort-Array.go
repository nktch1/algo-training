package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	dict := map[int]int{}

	for i, el := range arr2 {
		dict[el] = i
	}

	sort.Slice(arr1, func (i, j int) bool {
		ii, ok := dict[arr1[i]]
		if !ok {
			ii = arr1[i] + len(arr2)
		}

		jj, ok := dict[arr1[j]]
		if !ok {
			jj = arr1[j] + len(arr2)
		}

		return ii < jj
	})

	return arr1
}
