package slice

func Contains(slice []int, number int) bool {
	for _, value := range slice {
		if value == number {
			return true
		}
	}

	return false
}
