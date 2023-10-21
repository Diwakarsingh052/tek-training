package sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Declare the Go package as 'sum'
package sum



// Define a function to test the SumInt function.
// This function uses table-driven testing, where each struct in the array
// represents a different test case.
func TestSumInt(t *testing.T) {

	// Define an array of anonymous structs to act as different test cases.
	// Each struct consists of a test name, the input numbers as a slice, and the expected output.
	tt := [...]struct {
		//fields
		name    string // Name of the test
		numbers []int  // Input numbers for the test
		want    int    // Expected result for the test
	}{
		{
			name:    "one to five",  // Test case where numbers from one to five are summed
			numbers: []int{1, 2, 3, 4, 5},
			want:    15,  // Expect the sum to be 15 in this case
		},
		{
			name:    "nil slice",  // Test case where a nil slice is passed, expecting sum as 0
			numbers: nil,
			want:    0,
		},
		{
			name:    "one minus one",  // Test case where 1 and -1 are summed
			numbers: []int{1, -1},
			want:    0,  // Expect the sum to be 0 in this case
		},
	}

	// Iterate over each test case in 'tt'
	for _, tc := range tt {
		// Use 't.Run()' to execute a sub-test for each test case
		t.Run(tc.name, func(t *testing.T) {

			// Call SumInt with the input numbers for the current test case
			got := SumInt(tc.numbers)

			// Assert that the result is equal to the expected value for the current test case
			require.Equal(t, tc.want, got)
		})
	}
}