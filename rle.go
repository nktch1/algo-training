package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rle(raw string) string {
	if len(raw) == 0 {
		return raw
	}

	if len(raw) == 1 {
		return raw + "1"
	}

	var (
		res      = strings.Builder{}
		rawRunes = []rune(raw)
		c        = 1

		n = len(rawRunes)
		i int
	)

	for i = 0; i < n-1; i++ {
		if rawRunes[i] != rawRunes[i+1] {
			res.WriteString(string(rawRunes[i]))
			res.WriteString(strconv.Itoa(c))

			c = 1

			continue
		}

		c++
	}

	if rawRunes[n-1] != rawRunes[i-1] {
		res.WriteString(string(rawRunes[n-1]))
		res.WriteString(strconv.Itoa(1))
	} else {
		c++

		res.WriteString(string(rawRunes[n-1]))
		res.WriteString(strconv.Itoa(c))
	}

	fmt.Println(res.String())

	return res.String()
}