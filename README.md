
---

# Linear Statistics Project

This project provides a set of tools for calculating various statistical measures, including the Pearson Correlation Coefficient, using data from input files.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [Contact](#contact)

## Introduction

The Linear Statistics Project is designed to help users calculate key statistical metrics such as the Pearson Correlation Coefficient. This tool can be useful for data analysis and understanding relationships between different data sets.

## Features

- **Pearson Correlation Calculation:** Computes the Pearson Correlation Coefficient for a given dataset.
- **Modular Design:** Easily extendable to include more statistical measures.
- **Comprehensive Testing:** Includes unit tests to ensure the accuracy of calculations.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://learn.zone01kisumu.ke/git/seodhiambo/linear-stats.git
   cd linear-statistics
   ```

2. **Install dependencies:**

   Make sure you have [Go](https://golang.org/doc/install) installed. The project does not have any external dependencies, but ensure you have the correct Go environment set up.

3. **Build the project:**

   ```bash
   go build -o main .
   ```

4. **Run the Project:**

   ```bash
   ./bin/linear-stats && ./main data.txt 
   ```


## Usage

1. **Prepare your data file:**

   Create a text file with your dataset, where each line contains a single integer value.

2. **Run the program:**

   ```bash
   go run main.go data.txt
   ```

   This will calculate the Pearson Correlation Coefficient and print the result to the console.

## Testing

This project includes unit tests to verify the accuracy of the statistical calculations.

1. **Run the tests:**

   ```bash
   go test ./...
   ```

   The tests will run and output the results to the console.

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push`).
6. Open a pull request.


## Contact

For questions or suggestions, please contact the project maintainer:

- **Name:** Seth Athooh
- **Email:** [oathooh@gmail.com]
- **Gitea:** [seodhiambo](https://learn.zone01kisumu.ke/git/seodhiambo/linear-stats.git)

---
