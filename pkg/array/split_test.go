package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitIntoBatches(t *testing.T) {
	t.Parallel()

	type testCase[T any] struct {
		name           string
		inputObjects   []T
		batchSize      int
		expectedOutput [][]T
	}

	testCases := []testCase[int]{
		{
			name:           "success",
			inputObjects:   []int{1, 2, 3, 4, 5},
			batchSize:      2,
			expectedOutput: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:           "success with empty array",
			inputObjects:   []int{},
			batchSize:      2,
			expectedOutput: [][]int{{}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := SplitIntoBatches(tc.inputObjects, tc.batchSize)

			assert.Equal(t, tc.expectedOutput, result)
		})
	}
}
