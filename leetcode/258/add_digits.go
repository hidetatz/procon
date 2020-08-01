package main

func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	// e.g. num is 734:
	// 734 = 7 * 100      + 3 * 10      + 4
	//     = 7 * (99 + 1) + 3 * (9 + 1) + 4
	//     = 9(7 * 11 + 3) + (7 + 3 + 4)
	// because 9(7 * 11 + 3) is a multiple of 9, the mod is always 0.
	// so, the mod of 734 / 9 will be the same as
	// the mod of 7 + 3 + 4
	if mod := num % 9; mod != 0 {
		return mod
	}

	return 9
}
