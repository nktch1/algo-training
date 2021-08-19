package main

import (
	"fmt"
	"sort"
)

var dp = map[int]int{}

func getKth(lo int, hi int, k int) int {
	type valToPower struct {
		val int
		power int
	}

	var powers []valToPower

	for i := lo; i <= hi; i++ {
		cnt := 0
		recursion(i, &cnt)
		powers = append(powers, valToPower{i, cnt})
	}

	// fmt.Println(powers)

	sort.Slice(powers, func(i, j int) bool {
		return powers[i].power < powers[j].power
	})

	fmt.Println(dp)

	return powers[k-1].val
}

func recursion(i int, counter *int) {
	if i == 1 {
		return
	}

	nextValue := i / 2

	if !isEven(i) {
		nextValue = 3 * i + 1
	}


	recursion(nextValue, counter)

	*counter++


	//
	//if isEven(nextValue) {
	//	recursion(nextValue, counter)
	//}
	//
	//r, ok := dp[nextValue]
	//if ok {
	//	*counter += r
	//	return
	//}
	//
	//if isEven(i) {
	//	recursion(nextValue, counter)
	//} else {
	//
	//	recursion(nextValue, counter)
	//}

	dp[nextValue] = *counter
}