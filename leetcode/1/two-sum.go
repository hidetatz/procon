package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum_optimized(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		m[n] = i
	}

	for i, n := range nums {
		ans := target - n
		a, ok := m[ans]
		if !ok {
			continue
		}

		if a == i {
			continue
		}

		return []int{i, a}
	}

	return []int{}
}

func main() {
	twoSum(nil, 0)
}
