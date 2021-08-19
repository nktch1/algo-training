package main

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for idx := range nums {
		v := target - nums[idx]
		val, e := m[v]
		if e {
			return []int{val, idx}
		}
		m[nums[idx]] = idx
	}

	return []int{}
}

// =========================================

type CompWithIdx struct {
	val int
	idx int
}

func twoSumLogN(nums []int, target int) []int {
	var (
		arr  []CompWithIdx
		i, j = 0, len(nums) - 1
	)

	for idx := range nums {
		arr = append(arr, CompWithIdx{
			val: nums[idx],
			idx: idx,
		})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].val < arr[j].val
	})

	for i < j {
		comp := arr[i].val + arr[j].val

		if comp == target {
			return []int{arr[i].idx, arr[j].idx}
		}

		if comp < target {
			i++
		}

		if comp > target {
			j--
		}
	}

	return []int{}
}
