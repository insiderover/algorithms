package searching

func linearSearch(arr []int, needle int) int {
	for i, v := range arr {
		if v == needle {
			return i
		}
	}

	return -1
}
