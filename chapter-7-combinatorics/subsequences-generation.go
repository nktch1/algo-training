package chapter_7_combinatorics

var finished = false

func backtrack(a []int, k int) {

	var (
		c           [2]int
		ncandidates int
	)

	if isASolution(a, k) {
		processSolution(a, k)
		return
	}

	k++
	constructCandidates(a, k, c, &ncandidates)

	for i := 0; i < ncandidates; i++ {
		a[k] = c[i]

		makeMove(a, k)
		backtrack(a, k)
		unmakeMove(a, k)

		if finished {
			return
		}
	}

}

func unmakeMove(a []int, k int) {

}

func makeMove(a []int, k int) {

}

func constructCandidates(a []int, k int, c []int, i *int) {

}

func processSolution(a []int, k int) {

}

func isASolution(a []int, k int) bool {
	return true
}
