package main

func flipAndInvertImage(A [][]int) [][]int {
	ret := [][]int{}

	for _, image := range A {
		fixed := make([]int, len(image))
		for i := len(image) - 1; i >= 0; i-- {
			if image[i] == 0 {
				fixed[len(image)-1-i] = 1
			} else {
				fixed[len(image)-1-i] = 0
			}
		}
		ret = append(ret, fixed)
	}

	return ret
}
