package searching

import "testing"

type benchmarkCaseBinarySearch struct {
	testSlice []int
	needle    int
}

func BenchmarkBinary(b *testing.B) {
	b.ReportAllocs()

	tCase := benchmarkCaseBinarySearch{
		testSlice: []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 5},
		needle:    3,
	}

	for i := 0; i < b.N; i++ {
		BinarySearch(tCase.testSlice, tCase.needle)
	}
}
