package rules

func Decreasing(input []int) bool {
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			return false
		}
	}
	return true
}

func Increasing(input []int) bool {
	for i := 1; i < len(input); i++ {
		if input[i] < input[i-1] {
			return false
		}
	}
	return true
}
