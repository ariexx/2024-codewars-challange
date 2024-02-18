package kata

func ArrayPArray(arr1 []int, arr2 []int) int {
	sum1 := 0
	sum2 := 0
	for _, v := range arr1 {
		sum1 += v
	}
	for _, v := range arr2 {
		sum2 += v
	}
	return sum1 + sum2
}
