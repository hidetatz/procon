package main

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}

	if rowIndex == 0 {
		return []int{1}
	}

	ret := []int{1}
	for i := 0; i < rowIndex; i++ {
		temp := make([]int, len(ret)+1)
		for j := 0; j < len(temp); j++ {
			// first
			if j == 0 {
				temp[j] = ret[0]

				// last
			} else if len(ret) == j {
				temp[j] = ret[j-1]

				// mid
			} else {
				temp[j] = ret[j-1] + ret[j]
			}
		}
		ret = temp
	}

	return ret
}
