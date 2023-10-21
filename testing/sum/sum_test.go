package sum

import "testing"

// This is a test function that uses the testing package in go.
// The function name starts with the Test keyword which is necessary for the compiler to recognize this as a test function.
func TestSumInt(t *testing.T) {

	// Defining a slice of integers x to test on.
	x := []int{1, 2, 3, 4, 5}

	// The expected result of the sum in the slice x.
	want := 15

	// Run the function SumInt with the slice x and store the result in the got variable.
	got := SumInt(x)

	// Checking if the function's output matches the expected output.
	// If they are not equal, the test will fail and print the error message with the expected and got values.
	if got != want {
		t.Errorf("sum of 1 to 5 should be %v; got %v", want, got) // The error message follows the format: "sum of 1 to 5 should be [expected]; got [actual]".
		// Uncomment next line to stop the test if it fails at this point.
		//t.Fatalf("sum of 1 to 5 should be %v; got %v", want, got)
	}

	/* Uncomment the following to add a new test case
	// Testing with a nil slice
		//x = nil
		// Run the function with the nil slice and store the result
		//got = SumInt(x)

		// Expected output with the nil slice
		//want = 0

		// Check if the function's output matches expected output
		//if got != want {
		//	t.Errorf("sum of nil should be %v; got %v", want, got)
		//}
	*/

	// New slice with different numbers for testing.
	x = []int{1, -1}

	// Run the function with the new slice and store the result.
	got = SumInt(x)

	// Expected output with the new slice.
	want = 0

	// Checking if the function's output matches the expected output with the new slice.
	if got != want {
		t.Errorf("sum of 1,-1 should be %v; got %v", want, got)
	}
}
