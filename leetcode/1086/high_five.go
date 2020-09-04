package main

import "sort"

func highFive(items [][]int) [][]int {
	m := map[int][]int{}

	for _, result := range items {
		id := result[0]
		score := result[1]

		m[id] = append(m[id], score)
	}

	ret := [][]int{}
	for id, scores := range m {
		sort.Sort(sort.Reverse(sort.IntSlice(scores)))
		sum := 0
		for i := 0; i < 5; i++ {
			sum += scores[i]
		}

		ret = append(ret, []int{id, sum / 5})
	}

	sort.Slice(ret, func(i, j int) bool { return ret[i][0] < ret[j][0] })
	return ret
}
