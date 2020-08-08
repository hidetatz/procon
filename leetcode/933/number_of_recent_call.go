package main

type RecentCounter struct {
	t []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.t = append(this.t, t)
	ret := 0
	for i := len(this.t) - 1; i >= 0; i-- {
		if t-this.t[i] <= 3000 {
			ret++
		} else {
			break
		}
	}

	return ret
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
