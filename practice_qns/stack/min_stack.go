package stack

type MinStack struct {
	top *StackNode
	min int
}

type StackNode struct {
	val     int
	next    *StackNode
	lastmin int
}

func Constructor() MinStack {
	return MinStack{
		top: nil,
	}
}

func (this *MinStack) Push(val int) {
	var newtop *StackNode
	if this.top == nil {
		newtop = &StackNode{val: val, next: this.top}
		this.min = val
	} else {
		newtop = &StackNode{val: val, next: this.top, lastmin: this.min}
	}
	this.top = newtop
	if this.top.val < this.min {
		this.min = this.top.val
	}
}

func (this *MinStack) Pop() {
	if this.top.next == nil {
		// if there's only 0 or 1 item to pop
		this.top = nil
		return
	}
	this.min = this.top.lastmin
	*this.top = *this.top.next
}

func (this *MinStack) Top() int {
	return this.top.val
}

func (this *MinStack) GetMin() int {
	return this.min
}
