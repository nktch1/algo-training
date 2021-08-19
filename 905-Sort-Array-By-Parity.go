package main

func sortArrayByParity(nums []int) []int {
	k := 0

	for i := 0; i < len(nums); i++ {
		if isEven(nums[i]) {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}

	return nums
}
