package sorting

import "testing"

func BenchmarkInsertionSort(b *testing.B) {
	b.ReportAllocs()

	var testSlice = []int{5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5}

	for i := 0; i < b.N; i++ {
		InsertionSort(testSlice)
	}
}
