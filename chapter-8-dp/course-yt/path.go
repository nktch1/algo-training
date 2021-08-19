package course_yt

func maxPriceOnTheWay(i, j int) int {
	// F(i, j) = a(i, j) + max(F(i+1, j); F(i, j+1))

	dp := make([][]int, shape[1])

	for idx := range dp {
		dp[idx] = make([]int, shape[0])
	}

	dp[0][shape[1]-1] = 1

	for ii := 0; ii < shape[1]; ii++ {
		for jj := shape[0] - 1; jj >= 0; jj-- {

			if ii > 0 && jj == shape[0]-1 {
				dp[ii][jj] += dp[ii-1][jj] + matrix[ii][jj]
				continue
			}

			if ii == 0 && jj < shape[0]-1 {
				dp[ii][jj] += dp[ii][jj+1] + matrix[ii][jj]
				continue
			}

			if ii == 0 {
				continue
			}

			dp[ii][jj] += max(dp[ii-1][jj], dp[ii][jj+1]) + matrix[ii][jj]
		}
	}

	return dp[shape[0]-1-j][i]
}

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}