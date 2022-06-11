package homework

func reverse(input []int64) (result []int64) {
	if len(input) == 0 {
		return
	}
	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, (input)[i])
	}
	return
}
