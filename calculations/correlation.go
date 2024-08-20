package calculations

import "math"

// PearsonCorrelation calculates the Pearson correlation coefficient
// between the index positions of the elements and their values in the input slice.
// The coefficient measures the linear correlation between two variables.
func PearsonCorrelation(data []int) float64 {
    N := float64(len(data)) // Convert length of data to float64 for precision
    var sumX, sumY, sumXY, sumX2, sumY2 float64

    // Iterate through each element in the data slice
    for i := 0; i < len(data); i++ {
        x := float64(i)         // Index position (x-axis value)
        y := float64(data[i])   // Data value (y-axis value)
        sumX += x               // Sum of x values
        sumY += y               // Sum of y values
        sumXY += x * y          // Sum of the product of x and y values
        sumX2 += x * x          // Sum of squares of x values
        sumY2 += y * y          // Sum of squares of y values
    }

    // Calculate the Pearson correlation coefficient (r)
    numerator := (N * sumXY) - (sumX * sumY) // Numerator of the formula
    denominator := math.Sqrt((N*sumX2 - sumX*sumX) * (N*sumY2 - sumY*sumY)) // Denominator of the formula

    if denominator == 0 {
        return 0 // Avoid division by zero; return 0 if denominator is 0
    }

    r := numerator / denominator // Compute the correlation coefficient
    return r
}
