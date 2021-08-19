package main

import "math/rand"

type Solution struct {
	src []int
}


func Constructor(nums []int) Solution {
	return Solution{src: nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.src
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {

	var nw = make([]int, len(this.src))
	copy(nw, this.src)

	rand.Shuffle(len(nw), func (i, j int) {
		nw[i], nw[j] = nw[j], nw[i]
	})

	return nw
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
