package course_yt

const endless = -99

func maxPriceOnTheWayWithRedRooms(i, j, k int) int {

	// A). комнаты сверху и справа красные: 		max(F(i+1, j, k-1), F(i, j+1, k-1))
	// B). комнаты сверху и справа синие:   		max(F(i+1, j, k), 	F(i, j+1, k))
	// C). комната сверху красная, справа синяя:   	max(F(i+1, j, k-1), F(i, j+1, k))
	// D). комната сверху синяя, справа красная:   	max(F(i+1, j, k), 	F(i, j+1, k-1))

	// F(i, j, k) = a(i, j) + {A, B, C, D}

	// порядок вычисления: F(1,1) <- F(i,n) <- F(n,j) <- F(n,n)

	var (
		dp        = make([][]int, shape[1])
		visitedRR int
	)

	for idx := range dp {
		dp[idx] = make([]int, shape[0])
	}

	// в начальной точке значение равно весу комнаты
	dp[0][shape[1]-1] = matrix[0][shape[1]-1]

	for ii := 0; ii < shape[1]; ii++ {
		for jj := shape[0]-1; jj >= 0; jj-- {

			// заполнить первую строку матрицы
			if ii == 0 && jj < shape[0]-1 {
				dp[ii][jj] += matrix[ii][jj] + dp[ii][jj+1]

			// заполнить последний столбец матрицы
			} else if ii > 0 && jj == shape[0]-1 {
				dp[ii][jj] += matrix[ii][jj] + dp[ii-1][jj]

			// заполнить все остальное
			} else if ii != 0 && jj != shape[0]-1 {
				dp[ii][jj] += matrix[ii][jj] + max(dp[ii-1][jj], dp[ii][jj+1])
			}

			// если текущая комната красная
			if isRedRoom(ii, jj) {
				visitedRR++

				// и превышен лимит посещения красных комнат
				if visitedRR > k {
					// присвоить (-∞), чтобы max() не выбрал этот путь в дальнейшем
					dp[ii][jj] = endless
				}
			}
		}
	}

	return dp[shape[0]-1-j][i]
}

func isRedRoom(i, j int) bool {
	for _, po := range redRoomsPos {
		if po[0] == i && po[1] == j {
			return true
		}
	}

	return false
}