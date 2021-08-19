package main

type MinStack struct {
	src  []int
	mins []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(val int)  {
	if len(this.mins) == 0 || this.mins[len(this.mins)-1] >= val {
		this.mins = append(this.mins, val)
	}

	this.src = append(this.src, val)
}


func (this *MinStack) Pop()  {
	if len(this.src) == 0 {
		return
	}

	v := this.src[len(this.src)-1]

	if this.mins[len(this.mins)-1] == v {
		this.mins = this.mins[:len(this.mins)-1]
	}

	this.src = this.src[:len(this.src)-1]
}


func (this *MinStack) Top() int {
	if len(this.src) == 0 {
		return 0
	}

	return this.src[len(this.src)-1]
}


func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		return 0
	}

	return this.mins[len(this.mins)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
