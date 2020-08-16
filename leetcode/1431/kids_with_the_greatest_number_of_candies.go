package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for _, candy := range candies {
		if max < candy {
			max = candy
		}
	}

	enoughCandyAmount := max - extraCandies

	ret := []bool{}
	for _, candy := range candies {
		ret = append(ret, enoughCandyAmount <= candy)
	}

	return ret
}
