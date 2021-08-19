package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var (
		i, j  int
		r []int
	)


	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
			continue
		}

		if nums1[i] > nums2[j] {
			j++
			continue
		}

		r = append(r, nums1[i])
		i++
		j++
	}

	return r
}
