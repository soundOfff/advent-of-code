package rules

var (
	Tolerance = 1
)

func MaxDiff(input []int, max int) bool {
	for i := 1; i < len(input); i++ {
		if absValue(input[i]-input[i-1]) > max {
			return false
		}
	}
	return true
}

func MinDiff(input []int, min int) bool {
	for i := 1; i < len(input); i++ {
		if absValue(input[i]-input[i-1]) < min {
			return false
		}
	}
	return true
}

func absValue(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
