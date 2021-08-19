package chapter_8_dp

import (
	"fmt"
	"math"
)

//	максимальное			максимальная
//	общее					общая
//	подмножество			подстрока

//    a b c d a f c			  a b c d a f c
//	a 1 1 1 1 1 1 1			a 1 0 0 0 1 0 0
//	b 1 2 2 2 2 2 2			b 0 2 0 0 0 0 0
//	c 1 2 3 3 3 3 3			c 0 0 3 0 0 0 1
//	d 1 2 3 4 4 4 4			d 0 0 0 4 0 0 0
//	r 1 2 3 4 4 4 4			r 0 0 0 0 0 0 0
//	f 1 2 3 4 4 5 5			f 0 0 0 0 0 1 0
//	k 1 2 3 4 4 5 5			k 0 0 0 0 0 0 0

func maxCommonSubstring(s1, s2 string) int {

	var (
		r1 = []rune(s1)
		r2 = []rune(s2)
		dp = make([][]int, len(r1))
		max = math.MinInt32
	)

	for idx := range dp {
		dp[idx] = make([]int, len(r2))
	}

	for i := 0; i < len(r1); i++ {
		for j := 0; j < len(r2); j++ {

			if r1[i] == r2[j] {
				if i == 0 || j == 0 {
					dp[i][j] += 1
					continue
				}

				v := dp[i-1][j-1] + 1

				dp[i][j] += v

				if max < v {
					max = v
				}
			}
		}
	}

	for i := range dp {
		fmt.Println(dp[i])
	}

	return max
}

func maxCommonSubsequence(s1, s2 string) int {

	var (
		r1 = []rune(s1)
		r2 = []rune(s2)
		dp = make([][]int, len(r1))
		max = math.MinInt32
	)

	for idx := range dp {
		dp[idx] = make([]int, len(r2))
	}

	for i := 0; i < len(r1); i++ {
		for j := 0; j < len(r2); j++ {

			if r1[i] == r2[j] {
				if i == 0 || j == 0 {
					dp[i][j] += 1
					continue
				}

				v := dp[i-1][j-1] + 1

				dp[i][j] += v

				if max < v {
					max = v
				}
			} else {
				if i == 0 || j == 0 {
					dp[i][j] = dp[0][0]
					continue
				}

				dp[i][j] = maxInt(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	for i := range dp {
		fmt.Println(dp[i])
	}

	return max
}

func maxInt(i, j int) int {
	if i < j {
		return j
	}

	return i
}
