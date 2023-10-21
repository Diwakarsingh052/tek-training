package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// This function is a test for the function named "DoubleHandler". It is used to validate the behaviour of DoubleHandler.
// The function "TestDoubleHandler" takes a pointer to a testing.T object as an argument, which provides methods for
// automated testing of Go source code. This object contains all the information about a single test run.
func TestDoubleHandler(t *testing.T) {
	// A series of test scenarios (test cases) are defined. Each scenario is represented by a struct consisting of name,
	// value, msg, double and statusCode. The struct array 'tt' represents the test cases.
	tt := [...]struct {
		name       string
		value      string
		msg        string
		double     int
		statusCode int
	}{
		// Here are the test cases.
	}

	// Loop through each of the test cases.
	for _, tc := range tt {
		// The t.Run method allows for running "subtests", one for each test case. The first argument to Run is the name of
		// the subtest, the second argument is the actual test function.
		t.Run(tc.name, func(t *testing.T) {
			// A new HTTP request is created for each test.
			// This request simulates calling the DoubleHandler function and passing the 'value' as a query parameter.
			r, err := http.NewRequest(http.MethodGet, "localhost:8080/double?v="+tc.value, nil)
			// Any error occurred while creating the NewRequest is considered as an error in test setup. So, if this occurs
			// it fails the test immediately.
			require.NoError(t, err, "problem in constructing request")

			// httptest.NewRecorder() creates a new ResponseRecorder. ResponseRecorder implements http.ResponseWriter, but
			// additionally captures the response. This enables checking the response later on in test.
			rec := httptest.NewRecorder()
			// Call the function 'doubleHandler' being tested.
			doubleHandler(rec, r)

			// get the generated response by the handler.
			res := rec.Result()
			// read the response body.
			b, err := io.ReadAll(res.Body)
			// Expect no error in reading the body
			require.NoError(t, err, "could not fetch the body")
			// Compare the actual response with the expected one.
			msg := string(bytes.TrimSpace(b))
			// Checking whether actual response and expected response are same
			require.Equal(t, tc.msg, msg)
			// Checking whether the status code in response and expected status code are same.
			require.Equal(t, tc.statusCode, res.StatusCode)

			// Converting the original value to integer before doubling it to match format of expected response
			v, err := strconv.Atoi(tc.value)
			// If error occurs while converting string to integer, the original number must have been non-integer, so its
			// double is expected as 0.
			if err != nil {
				require.Equal(t, 0, tc.double)
			}
			// Checking whether the calculated double value and the expected double value are same.
			require.Equal(t, tc.double, v*2)

		})
	}
}
