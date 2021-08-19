package main

func singleNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n := 0
	for idx := range nums {
		n ^= nums[idx]
	}

	return n
}