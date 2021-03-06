package searching

import "testing"

func BenchmarkBinarySearch(b *testing.B) {
	b.ReportAllocs()

	var testSlice = make([]int, 0, 100)
	for i := 100; i > 0; i-- {
		testSlice = append(testSlice, i)
	}
	for i := 0; i < b.N; i++ {
		BinarySearch(testSlice, 1)
	}
}
