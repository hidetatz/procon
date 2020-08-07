package main

import "fmt"

type MovingAverage struct {
	size int
	vals []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	this.vals = append(this.vals, val)
	s := 0

	if len(this.vals) < this.size {
		for _, v := range this.vals {
			s += v
		}
		return float64(s) / float64(len(this.vals))
	}

	cnt := 0
	for i := len(this.vals) - 1; ; i-- {
		s += this.vals[i]
		cnt++
		if cnt == this.size {
			break
		}
	}

	fmt.Println(s)

	return float64(s) / float64(this.size)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
