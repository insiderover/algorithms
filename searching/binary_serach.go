package searching

func BinarySearch(arr []int, needle int) int {
	if len(arr) == 0 || needle < arr[0] || needle > arr[len(arr)-1] {
		return -1
	}

	leftBound := 0
	rightBound := len(arr) - 1
	var mid int

	for {
		mid = (leftBound + rightBound) / 2

		if arr[mid] == needle {
			return mid
		}

		if leftBound > rightBound {
			return -1
		}

		if arr[mid] > needle {
			rightBound = mid - 1
		} else {
			leftBound = mid + 1
		}
	}
}
