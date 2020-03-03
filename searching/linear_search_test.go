package searching

import "testing"

type testCaseLinearSearch struct {
	testSlice []int
	needle    int
	expected  int
}

var testDataLinearSearch = []testCaseLinearSearch{
	{[]int{}, 10, -1},
	{[]int{5, 4, 3, 2, 1}, 10, -1},
	{[]int{5, 4, 3, 2, 1}, 5, 0},
	{[]int{5, 4, 3, 2, 1}, 1, 4},
}

func TestLinearSearch(t *testing.T) {
	for _, pair := range testDataLinearSearch {
		ix := linearSearch(pair.testSlice, pair.needle)
		if ix != pair.expected {
			t.Error(
				"Needle", pair.needle,
				"Expected index ", pair.expected,
				"got ", ix)
		}
	}
}
