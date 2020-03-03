package sorting

import "testing"

type benchmarkCaseBubbleSort struct {
	testSlice []int
}

func BenchmarkBinary(b *testing.B) {
	b.ReportAllocs()

	tCase := benchmarkCaseBubbleSort{
		testSlice: []int{5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5},
	}

	for i := 0; i < b.N; i++ {
		BubbleSort(tCase.testSlice)
	}
}
