package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadDataFromFile(filePath string) ([]int, error) {
	// Open the specified file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}
		// Convert the line to an integer
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		// Add the integer to the data slice
		data = append(data, num)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
