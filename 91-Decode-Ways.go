package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	ways := make([]int, len(s)+1)

	ways[0] = 1 // для строки длиной 0 - 1 вариант кодирования
	ways[1] = 0 // для строки длиной 1 - 2 варианта, если начинается на 0,
	// то тут ничего не поделаешь,

	if s[0] != '0' {  // иначе - 1 вариант
		ways[1] = 1
	}

	// "12" -> [1 1 2]

	for i := 2; i <= len(s); i++ { // начиная с третьего элемента массива
		oneInt, _ := strconv.Atoi(s[i-1:i]) // берем вариант с одним символом
		if oneInt >= 1 {
			ways[i] += ways[i-1] // и если он удовлетворяет условию,
		} // то обновляем количество вариантов для него

		twoInt, _ := strconv.Atoi(s[i-2:i]) // берем вариант с двумя символами
		if twoInt >= 10 && twoInt <= 26 {   // если по условию - ОК
			ways[i] += ways[i-2] // то обновляем количество вариантов для него
		} // ways[i-2] - количество вариантов для кодирования
	}

	fmt.Println(ways)

	return ways[len(s)]
}

