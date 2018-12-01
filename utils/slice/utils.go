package slice

func Sum(slice []int) (result int) {
	for _, value := range slice {
		result += value
	}

	return result
}
