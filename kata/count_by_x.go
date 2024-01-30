package kata

func CountBy(x, n int) []int {
	var result []int
	for i := 1; i <= n; i++ {
		result = append(result, x*i)
	}
	return result
}
