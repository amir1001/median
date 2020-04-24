package median

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	title string
	A     []float64
	B     []float64
}

type errorTestCase struct {
	title string
	A     []float64
	B     []float64
}

var testCases = []testCase{
	{
		title: "test case #1",
		A:     []float64{},
		B:     []float64{9},
	}, {
		title: "test case #2",
		A:     []float64{},
		B:     []float64{3, 5},
	}, {
		title: "test case #3",
		A:     []float64{5},
		B:     []float64{3},
	}, {
		title: "test case #4",
		A:     []float64{3, 5},
		B:     []float64{},
	}, {
		title: "test case #5",
		A:     []float64{3},
		B:     []float64{7, 9},
	}, {
		title: "test case #6",
		A:     []float64{3, 7},
		B:     []float64{9},
	}, {
		title: "test case #7",
		A:     []float64{3, 7, 9},
		B:     []float64{},
	}, {
		title: "test case #8",
		A:     []float64{},
		B:     []float64{3, 7, 9},
	}, {
		title: "test case #9",
		A:     []float64{8, 9},
		B:     []float64{1, 2},
	}, {
		title: "test case #10",
		A:     []float64{1, 8},
		B:     []float64{2, 9},
	}, {
		title: "test case #11",
		A:     []float64{1, 9},
		B:     []float64{2, 8},
	}, {
		title: "test case #12",
		A:     []float64{1, 2, 5, 7},
		B:     []float64{3, 4, 6, 8, 9},
	}, {
		title: "test case #12",
		A:     []float64{1, 2, 3, 4, 4, 5, 7},
		B:     []float64{3, 4, 6, 8, 9},
	},
}

var errorTestCases = []errorTestCase{
	{
		title: "empty input",
		A:     []float64{},
		B:     []float64{},
	}, {
		title: "unsorted input",
		A:     []float64{1, 2, 9, 8, 7, 5, 4},
		B:     []float64{39, 38, 37, 36, 33, 32, 21, 20, 18, 17},
	},
}

func TestGetMedian(t *testing.T) {
	for _, tc := range testCases {
		var mergedSlice []float64
		mergedSlice = append(mergedSlice, tc.A...)
		mergedSlice = append(mergedSlice, tc.B...)
		sort.Float64s(mergedSlice)
		length := len(mergedSlice)
		parity := length % 2
		var median float64
		if parity == 1 {
			median = mergedSlice[length/2]
		} else {
			median = (mergedSlice[length/2] + mergedSlice[length/2-1]) / 2
		}

		m, err := GetMedian(tc.A, tc.B)
		assert.Equal(t, median, m, tc.title)
		assert.Empty(t, err, tc.title)
	}
}

func TestGetMedianError(t *testing.T) {
	for _, tc := range errorTestCases {
		_, err := GetMedian(tc.A, tc.B)
		assert.Error(t, err, tc.title)
	}
}
