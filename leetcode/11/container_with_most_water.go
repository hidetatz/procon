package main

func maxArea(height []int) int {
	ret := 0
	for i, j := 0, len(height)-1; i < j; {
		ret = max(ret, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ret
}

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
