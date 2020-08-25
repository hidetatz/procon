package main

func numberOfSteps(num int) int {
	ret := 0
	for num != 0 {
		ret++
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
	}

	return ret
}
