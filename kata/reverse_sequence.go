package kata

func ReverseSeq(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = n - i
	}

	return result
}
