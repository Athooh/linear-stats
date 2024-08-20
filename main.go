// package main

// import (
// 	"fmt"
// 	"os"

// 	"linear-stats/calculations"
// 	"linear-stats/utils"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: go run main.go <data-file>")
// 		return
// 	}

// 	filePath := os.Args[1]
// 	data, err := utils.ReadDataFromFile(filePath)
// 	if err != nil {
// 		fmt.Println("Error reading data:", err)
// 		return
// 	}

// 	slope, intercept := calculations.LinearRegression(data)
// 	correlation := calculations.PearsonCorrelation(data)

// 	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
// 	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", correlation)
// }

package main

import (
	"fmt"
	"os"

	"linear-stats/calculations"
	"linear-stats/utils"
)

func main() {
	// Check if the program was run with the required argument (data file path)
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <data-file>")
		return
	}

	// Get the file path from the command-line arguments
	filePath := os.Args[1]

	// Read data from the specified file
	data, err := utils.ReadDataFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	// Perform linear regression on the data
	slope, intercept := calculations.LinearRegression(data)

	// Calculate Pearson correlation coefficient
	correlation := calculations.PearsonCorrelation(data)

	// Display the results
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", correlation)
}
