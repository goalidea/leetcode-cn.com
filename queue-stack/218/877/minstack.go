package minstack

type MinStack struct {
	s []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		s: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	mini := this.s[0]
	for _, v := range this.s {
		if v < mini {
			mini = v
		}
	}
	return mini
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
