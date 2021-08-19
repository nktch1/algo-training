package main

// сначала транспонировать, потом отразить

// 1 2 -> 1 3 -> 3 1
// 3 4	  2 4    4 2

func rotateMatrix(matrix [][]int)  {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}

	for i := 0; i < len(matrix); i++ {
		reverseMatrix(matrix[i])
	}
}

func reverseMatrix(nums []int) {
	i, j := 0, len(nums)-1

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
