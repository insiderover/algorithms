package sorting

func SelectionSort(arr []int) {
	var min int
	for i := 0; i < len(arr)-1; i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
