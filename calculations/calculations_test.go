// package calculations

// import (
// 	"math"
// 	"testing"
// )
// func TestLinearRegression(t *testing.T) {
// 	tests := []struct {
// 		data     []int
// 		expectedM float64
// 		expectedC float64
// 	}{
// 		{[]int{1, 2, 3, 4, 5}, 1, 1},
// 		{[]int{2, 4, 6, 8, 10}, 2, 2}, 
// 		{[]int{1, 3, 5, 7, 9}, 2, 1},
// 		{[]int{3, 3, 3, 3, 3}, 0, 3},
// 	}

// 	for _, test := range tests {
// 		m, c := LinearRegression(test.data)
// 		if math.Abs(m-test.expectedM) > 1e-6 || math.Abs(c-test.expectedC) > 1e-6 {
// 			t.Errorf("LinearRegression(%v) = (%v, %v); expected (%v, %v)", test.data, m, c, test.expectedM, test.expectedC)
// 		}
// 	}
// }

// func TestPearsonCorrelation(t *testing.T) {
// 	tests := []struct {
// 		data     []int
// 		expected float64
// 	}{
// 		{[]int{1, 2, 3, 4, 5}, 1},
// 		{[]int{2, 4, 6, 8, 10}, 1},
// 		{[]int{1, 3, 5, 7, 9}, 1},
// 		{[]int{5, 4, 3, 2, 1}, -1},
// 		{[]int{3, 3, 3, 3, 3}, 0},
// 	}

// 	for _, test := range tests {
// 		r := PearsonCorrelation(test.data)
// 		if math.Abs(r-test.expected) > 1e-6 {
// 			t.Errorf("PearsonCorrelation(%v) = %v; expected %v", test.data, r, test.expected)
// 		}
// 	}
// }

package calculations

import (
	"math"
	"testing"
)

// TestLinearRegression tests the LinearRegression function.
// It compares the calculated slope (m) and intercept (c) with the expected values
// for different sets of input data.
func TestLinearRegression(t *testing.T) {
	// Define test cases with input data and expected slope (m) and intercept (c)
	tests := []struct {
		data     []int
		expectedM float64
		expectedC float64
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 1},  // Data with slope 1 and intercept 1
		{[]int{2, 4, 6, 8, 10}, 2, 2}, // Data with slope 2 and intercept 2
		{[]int{1, 3, 5, 7, 9}, 2, 1},  // Data with slope 2 and intercept 1
		{[]int{3, 3, 3, 3, 3}, 0, 3},  // Data with slope 0 and intercept 3 (constant values)
	}

	// Iterate over each test case
	for _, test := range tests {
		// Calculate the slope and intercept using LinearRegression function
		m, c := LinearRegression(test.data)
		// Check if the calculated values are within a small tolerance of the expected values
		if math.Abs(m-test.expectedM) > 1e-6 || math.Abs(c-test.expectedC) > 1e-6 {
			// Report an error if the calculated values do not match the expected values
			t.Errorf("LinearRegression(%v) = (%v, %v); expected (%v, %v)", test.data, m, c, test.expectedM, test.expectedC)
		}
	}
}

// TestPearsonCorrelation tests the PearsonCorrelation function.
// It compares the calculated Pearson correlation coefficient with the expected value
// for different sets of input data.
func TestPearsonCorrelation(t *testing.T) {
	// Define test cases with input data and expected Pearson correlation coefficient
	tests := []struct {
		data     []int
		expected float64
	}{
		{[]int{1, 2, 3, 4, 5}, 1},  // Perfect positive linear correlation
		{[]int{2, 4, 6, 8, 10}, 1}, // Perfect positive linear correlation
		{[]int{1, 3, 5, 7, 9}, 1},  // Perfect positive linear correlation
		{[]int{5, 4, 3, 2, 1}, -1}, // Perfect negative linear correlation
		{[]int{3, 3, 3, 3, 3}, 0},  // No correlation (constant values)
	}

	// Iterate over each test case
	for _, test := range tests {
		// Calculate the Pearson correlation coefficient using PearsonCorrelation function
		r := PearsonCorrelation(test.data)
		// Check if the calculated coefficient is within a small tolerance of the expected value
		if math.Abs(r-test.expected) > 1e-6 {
			// Report an error if the calculated coefficient does not match the expected value
			t.Errorf("PearsonCorrelation(%v) = %v; expected %v", test.data, r, test.expected)
		}
	}
}
