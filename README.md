# Advent of Code 2023

My attempts at solutions for [Advent of Code 2023](https://adventofcode.com/2023) using Go.

Each day's puzzle is in its own folder:
- `day_xx.go`: contains the code that generates the solution
- `day_xx_test.go`: contains a simple test that checks the results using the sample data and the puzzle data. This helps me to know if I break something while refactoring.
- `day_xx_data.go`: contains the puzzle data (unique to each Advent account)
- `day_xx_test_data.go`: contains the sample data within the puzzle description.

Typically I will develop the code's solution using the sample data until I get the same result. Then I run it again using the puzzle data and submit. *Usually* that will result in the correct answer, but sometimes the puzzle data has cases that do not work the first time around.
