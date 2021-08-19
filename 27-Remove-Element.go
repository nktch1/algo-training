package main

func removeElement(nums []int, val int) int {
	j := len(nums)-1

	for i := len(nums)-1; i >= 0; i-- {
		if nums[i] == val {
			nums[j], nums[i] = nums[i], nums[j]
			j--
		}
	}

	return j+1
}
