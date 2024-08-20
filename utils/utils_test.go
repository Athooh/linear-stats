
package utils

import (
    "os"
    "testing"
)

func TestReadDataFromFile(t *testing.T) {
    // Create a temporary file with test data
    tempFile, err := os.CreateTemp("", "testdata*.txt")
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }
    defer os.Remove(tempFile.Name()) // Ensure the temporary file is removed

    sampleData := `189
113
121
114
145
110`

    if _, err := tempFile.Write([]byte(sampleData)); err != nil {
        t.Fatalf("Failed to write to temp file: %v", err)
    }
    tempFile.Close()

    // Test ReadDataFromFile function
    data, err := ReadDataFromFile(tempFile.Name())
    if err != nil {
        t.Fatalf("ReadDataFromFile returned an error: %v", err)
    }

    expectedData := []int{189, 113, 121, 114, 145, 110}
    if len(data) != len(expectedData) {
        t.Fatalf("Expected data length %d, got %d", len(expectedData), len(data))
    }

    for i, v := range data {
        if v != expectedData[i] {
            t.Errorf("Expected %d at index %d, got %d", expectedData[i], i, v)
        }
    }
}

func TestReadDataFromFileWithInvalidData(t *testing.T) {
    // Create a temporary file with invalid data
    tempFile, err := os.CreateTemp("", "testdata_invalid*.txt")
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }
    defer os.Remove(tempFile.Name()) // Ensure the temporary file is removed

    invalidData := `189
113
invalid
145`

    if _, err := tempFile.Write([]byte(invalidData)); err != nil {
        t.Fatalf("Failed to write to temp file: %v", err)
    }
    tempFile.Close()

    // Test ReadDataFromFile function with invalid data
    _, err = ReadDataFromFile(tempFile.Name())
    if err == nil {
        t.Fatalf("Expected an error for invalid data, but got none")
    }
}
