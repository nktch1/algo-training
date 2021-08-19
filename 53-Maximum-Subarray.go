package main

import "math"

// алгоритм Кадана
// если предыдущий элемент не отрицательный, добавить к текущему
// а потом найти максимальный элемент массива

func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
	}

	return maxInSlice(nums)
}

func maxInSlice(nums []int) int {
	max := math.MinInt32

	for _, el := range nums {
		if el > max {
			max = el
		}
	}

	return max
}
