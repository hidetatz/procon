package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}

	counts := []Count{}
	for i, cnt := range m {
		counts = append(counts, Count{i, cnt})
	}

	sort.Slice(counts, func(i, j int) bool { return counts[i].Cnt > counts[j].Cnt })

	ret := []int{}
	for i := 0; i < k; i++ {
		ret = append(ret, counts[i].Val)
	}

	return ret

}

type Count struct {
	Val int
	Cnt int
}
