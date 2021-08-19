package main

import "sort"

func groupAnagrams(strs []string) [][]string {

	var (
		dict = map[string][]string{}
		res  [][]string
	)

	for idx := range strs {
		runes := []rune(strs[idx])

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		dict[string(runes)] = append(
			dict[string(runes)],
			strs[idx],
		)
	}

	for _, v := range dict {
		res = append(res, v)
	}

	return res
}
