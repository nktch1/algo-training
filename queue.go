package main

type Queue []int

func (q *Queue) Pop() int {
	if *q == nil {
		return -1
	}

	sl := []int(*q)

	val := sl[0]
	copy(sl[:], sl[1:])

	*q = sl[:len(sl)-1]

	return val
}

func (q *Queue) Push(val int) {
	*q = append(*q, val)
}

func (q *Queue) Len() int {
	return len(*q)
}

