# Advent of Code 2023

My attempts at solutions for [Advent of Code 2023](https://adventofcode.com/2023) using Go.

Each day's puzzle is in its own folder:
- `day_xx.go` contains the code that generates the solution.
- `day_xx_test.go` contains a simple test that checks the results using the example data and the puzzle data. This helps me to know if I break something while refactoring. As the puzzles become more complicated, I'll often write more tests and use a TDD approach, which helps me to break down difficult problems into more manageable pieces.
- `day_xx_data.go` contains the puzzle data. The puzzle data is unique for each Advent user, to prevent simply copying another user's answers. But ideally a working solution will give the correct answer with any user's data.
- `day_xx_test_data.go` contains the example data from the puzzle description.

Typically I will develop the code's solution using the example data until I get the correct result. Then I run it again using the puzzle data and, if there are no errors, I submit. That will often result in the correct answer, but the puzzle data frequently contains cases that are not covered when testing with the example data.

The command below (assuming you're using Linux or Mac OS) create a new folder for the day indicated and populate it with (nearly) empty files with the above structure.
```
. newday.sh <day_xx>
```
