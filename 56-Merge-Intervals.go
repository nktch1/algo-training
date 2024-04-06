package main

import "sort"

func merge56(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastAdded := result[len(result)-1]

		if intervals[i][0] <= lastAdded[1] {
			lastAdded[1] = max(lastAdded[1], intervals[i][1])
			continue
		}

		result = append(result, intervals[i])
	}

	return result
}
