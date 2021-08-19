package main

// nums = [1 2 3 4 5], k = 3
// res = [3 4 5 1 2]
// инвертировать весь массив
// потом инвертировать подмассивы до и после k

func rotate(nums []int, k int)  {
	if nums == nil || len(nums) == 0 {
		return
	}

	k = k % len(nums)

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	n := len(nums)

	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}
