package searching

import "testing"

type testCaseBinarySearch struct {
	name      string
	testSlice []int
	needle    int
	expected  int
}

var testDataBinarySearch = []testCaseBinarySearch{
	{"test 1", []int{}, 10, -1},
	{"test 2", []int{1, 2, 3, 4, 5}, 10, -1},
	{"test 3", []int{1, 2, 3, 4, 5}, -10, -1},
	{"test 4", []int{1, 3, 5, 7, 9}, 4, -1},
	{"test 5", []int{1, 2, 3, 4, 5}, 1, 0},
	{"test 6", []int{1, 2, 3, 4, 5}, 5, 4},
	{"test 7", []int{1, 2, 3, 4, 5}, 3, 2},
	{"test 8", []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 5}, 3, 11},
}

func TestBinarySearch(t *testing.T) {
	for _, pair := range testDataBinarySearch {
		t.Run(pair.name, func(t *testing.T) {
			ix := BinarySearch(pair.testSlice, pair.needle)
			if ix != pair.expected {
				t.Error(
					"Needle", pair.needle,
					"Expected index ", pair.expected,
					"got ", ix)
			}
		})
	}
}
