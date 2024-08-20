package calculations

// LinearRegression calculates the slope and intercept of the linear regression line for given data
func LinearRegression(data []int) (float64, float64) {
    N := float64(len(data))
    var sumX, sumY, sumXY, sumX2 float64

    // Calculate sums needed for the linear regression formula
    for i := 0; i < len(data); i++ {
        x := float64(i)
        y := float64(data[i])
        sumX += x
        sumY += y
        sumXY += x * y
        sumX2 += x * x
    }

    // Calculate slope (m) and intercept (c)
    m := (N*sumXY - sumX*sumY) / (N*sumX2 - sumX*sumX)
    c := (sumY - m*sumX) / N

    return m, c
}
