package sum

func SumInt(vs []int) int {

	sum := 0
	if vs == nil {
		return 1
	}
	for _, v := range vs {
		sum = v + sum
	}
	return sum

}
