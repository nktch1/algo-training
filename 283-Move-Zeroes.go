package main

func moveZeroes(nums []int) {
	j := 0

	for idx := range nums {
		if nums[idx] != 0 {
			nums[idx], nums[j] = nums[j], nums[idx]
			j++
		}
	}
}