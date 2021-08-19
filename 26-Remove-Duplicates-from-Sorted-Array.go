package main

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return len(nums)
	}

	idx := 1 							// первое значение всегда уникально

	for j := 0; j < len(nums)-1; j++ { 	// итерации до предпоследнего значения
		if nums[j] != nums[j+1] { 		// если текущий элемент не равен следующему
			nums[idx] = nums[j+1] 		// то по второму указателю располагаем новое уникальное значение
			idx++                 		// смещаем второй указатель
		}
	}

	return idx
}
