package main

// stack
// not thread safe
type intStack struct {
	stack []int
}

func (is *intStack) push(i int) {
	is.stack = append(is.stack, i)
}

func (is *intStack) pop() int {
	if len(is.stack) == 0 {
		panic("empty")
	}
	last := is.stack[len(is.stack)-1]
	is.stack = is.stack[:len(is.stack)-1] // remove last
	return last
}

// queue
// not thread safe
type intQueue struct {
	queue []int
}

func (iq *intQueue) push(i int) {
	iq.queue = append(iq.queue, i)
}

func (iq *intQueue) pop() int {
	if len(iq.queue) == 0 {
		panic("empty")
	}
	first := iq.queue[0]
	iq.queue = iq.queue[1:] // remove first
	return first
}

// max/min
func max(i1, i2 int) int {
	if i1 < i2 {
		return i2
	}
	return i1
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}
