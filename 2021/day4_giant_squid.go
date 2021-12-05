package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day4Part1() {
	randNums, matrices := randNumsAndMatrices()

	/*
		Questions
		- Can numbers be repeated in a matrix? looks unique so far

		Approach 1
		- Go over each random number
		- Go over each matrix and mark random number as seen
		- Go over each row and column that number is on to see if it's "complete"
		- Complexity: All rand numbers * All matrices * 5 (rows) * 5 (cols) * (5 * 2)

		Approach 2
		- Store each matrix in a map. Use incremental numbers to identify them 1....M.
		- Go over each number in each matrix and store the number and all the positions it exists
			in. Use the form -> map[int][][]int -> number:List[[]int{matrix id, row, col}].
			Use the map from step 1.
		- Since we know every row and col should be 5, we can keep track of what is seen for 5 rows and 5 columns for each matrix
			map[matrix_id]: [2][5]int{}. If any gets to 5, the matrix_id represents the matrix we want
		- Go over the matrix and calculate the sum of all numbers, not including the row or column that won
		- Complexity: M + (M * R * C) + (N*K) + (R*C) => M + (25M) + (N*K) + (R*C)

		M=Total number of matrices
		R=Rows=5
		C=Columns=5
		K=Total number of positions a number occurs in. If numbers are unique in each matrix, the worst case is M. If not, it's M*R*C (as in, the number can be in every position in every matrix=).
		N=Random Numbers
	*/

	// randNums = []int{
	// 	7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3,
	// 	26, 1,
	// }
	// matrices = [][][]int{
	// 	{
	// 		{22, 13, 17, 11, 0},
	// 		{8, 2, 23, 4, 24},
	// 		{21, 9, 14, 16, 7},
	// 		{6, 10, 3, 18, 5},
	// 		{1, 12, 20, 15, 19},
	// 	},
	// 	{
	// 		{3, 15, 0, 2, 22},
	// 		{9, 18, 13, 17, 5},
	// 		{19, 8, 7, 25, 23},
	// 		{20, 11, 10, 24, 4},
	// 		{14, 21, 16, 12, 6},
	// 	},
	// 	{
	// 		{14, 21, 17, 24, 4},
	// 		{10, 16, 15, 9, 19},
	// 		{18, 8, 23, 26, 20},
	// 		{22, 11, 13, 6, 5},
	// 		{2, 0, 12, 3, 7},
	// 	},
	// }

	numPos := make(map[int][][]int)

	for id, matrix := range matrices {
		for i := range matrix {
			for j := range matrix[i] {
				num := matrix[i][j]
				positions, exists := numPos[num]
				newPosition := []int{id, i, j}
				if !exists {
					numPos[num] = make([][]int, 0)
					numPos[num] = append(numPos[num], newPosition)
				} else {
					positions = append(positions, newPosition)
					numPos[num] = positions
				}
			}
		}
	}

	// Matrices Row and Col Count
	// [i][0] = rows for matrix with id i
	// [i][0][2] = count of values found for matrix with id i at row 2
	// [i][1] = cols for matrix with id j
	// [i][1][2] = count of values found for matrix with id i at col 2
	mRowColCnt := make([][][]int, len(matrices))
	for id := range mRowColCnt {
		mRowColCnt[id] = [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
	}

	result := 0
	stop := false
	for _, num := range randNums {
		// Check if number exists in at least one matrix
		positions, exists := numPos[num]
		if !exists {
			continue
		}

		for _, position := range positions {
			matrixID := position[0]
			row := position[1]
			col := position[2]

			// Since those numbers must not be included in the sum later, make them 0
			// The numPos and mRowColCnt variables store the info we need.
			matrices[matrixID][row][col] = 0

			mRowColCnt[matrixID][0][row]++
			if mRowColCnt[matrixID][0][row] == 5 {
				sum := sumOfUnmarkedNumbers(matrices[matrixID])
				result = sum * num
				fmt.Println(sum, num, result)
				stop = true
				break
			}

			mRowColCnt[matrixID][1][col]++
			if mRowColCnt[matrixID][1][col] == 5 {
				sum := sumOfUnmarkedNumbers(matrices[matrixID])
				result = sum * num
				fmt.Println(sum, num, result)
				stop = true
				break
			}
		}

		if stop {
			break
		}
	}

}

func sumOfUnmarkedNumbers(matrix [][]int) int {
	sum := 0
	for rowIdx := range matrix {
		// if rowOrCol == "row" && rowIdx == idx {
		// 	continue
		// }

		for colIdx := range matrix[rowIdx] {
			// if rowOrCol == "col" && colIdx == idx {
			// 	continue
			// }

			sum = sum + matrix[rowIdx][colIdx]
		}
	}

	return sum
}

func day4Part2() {

	randNums, matrices := randNumsAndMatrices()
	// randNums := []int{
	// 	7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3,
	// 	26, 1,
	// }
	// matrices := [][][]int{
	// 	{
	// 		{22, 13, 17, 11, 0},
	// 		{8, 2, 23, 4, 24},
	// 		{21, 9, 14, 16, 7},
	// 		{6, 10, 3, 18, 5},
	// 		{1, 12, 20, 15, 19},
	// 	},
	// 	{
	// 		{3, 15, 0, 2, 22},
	// 		{9, 18, 13, 17, 5},
	// 		{19, 8, 7, 25, 23},
	// 		{20, 11, 10, 24, 4},
	// 		{14, 21, 16, 12, 6},
	// 	},
	// 	{
	// 		{14, 21, 17, 24, 4},
	// 		{10, 16, 15, 9, 19},
	// 		{18, 8, 23, 26, 20},
	// 		{22, 11, 13, 6, 5},
	// 		{2, 0, 12, 3, 7},
	// 	},
	// }

	numPos := make(map[int][][]int)

	for id, matrix := range matrices {
		for i := range matrix {
			for j := range matrix[i] {
				num := matrix[i][j]
				positions, exists := numPos[num]
				newPosition := []int{id, i, j}
				if !exists {
					numPos[num] = make([][]int, 0)
					numPos[num] = append(numPos[num], newPosition)
				} else {
					positions = append(positions, newPosition)
					numPos[num] = positions
				}
			}
		}
	}

	// Matrices Row and Col Count
	// [i][0] = rows for matrix with id i
	// [i][0][2] = count of values found for matrix with id i at row 2
	// [i][1] = cols for matrix with id j
	// [i][1][2] = count of values found for matrix with id i at col 2
	mRowColCnt := make([][][]int, len(matrices))
	for id := range mRowColCnt {
		mRowColCnt[id] = [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
	}

	result := 0

	// A matrix is a board
	matrixWinPosition := 1
	// Which number did the matrix win with
	matrixWinNumber := 0
	// Record the position of each winning board
	matrixWinSlice := make([]int, len(matrices))

	for _, num := range randNums {
		// Check if number exists in at least one matrix
		positions, exists := numPos[num]
		if !exists {
			continue
		}

		for _, position := range positions {
			matrixID := position[0]
			row := position[1]
			col := position[2]

			// Don't consider matrices that have won before
			if matrixWinSlice[matrixID] > 0 {
				continue
			}

			// Since those numbers must not be included in the sum later, make them 0
			// The numPos and mRowColCnt variables store the info we need.
			matrices[matrixID][row][col] = 0
			mRowColCnt[matrixID][0][row]++
			mRowColCnt[matrixID][1][col]++

			if mRowColCnt[matrixID][0][row] == 5 || mRowColCnt[matrixID][1][col] == 5 {
				matrixWinSlice[matrixID] = matrixWinPosition
				matrixWinPosition++
				matrixWinNumber = num
			}
		}
	}

	lastMatrixID := 0
	lastMatrixPos := matrixWinSlice[lastMatrixID]
	for i := 1; i < len(matrixWinSlice); i++ {
		if matrixWinSlice[i] > lastMatrixPos {
			lastMatrixID = i
			lastMatrixPos = matrixWinSlice[i]
		}
	}

	sum := sumOfUnmarkedNumbers(matrices[lastMatrixID])
	result = sum * matrixWinNumber
	fmt.Println(sum, matrixWinNumber, result)
}

func randNumsAndMatrices() ([]int, [][][]int) {
	file, _ := os.Open("2021/day4_giant_squid.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		if scanner.Text() == "\n" {
			fmt.Println("New Line")
		}
		lines = append(lines, scanner.Text())
	}

	randNumsStr := lines[0]
	randNumsStrSlice := strings.Split(randNumsStr, ",")
	randNums := make([]int, len(randNumsStrSlice))

	for i := range randNumsStrSlice {
		randNum, _ := strconv.Atoi(randNumsStrSlice[i])
		randNums[i] = randNum
	}

	matrices := make([][][]int, 0)

	for line := 2; line < len(lines); line++ {
		if lines[line] == "" {
			continue
		}

		// Each matrix is 5 x 5
		matrixString := make([][]string, 5)
		for i := range matrixString {
			matrixString[i] = strings.Split(lines[line], " ")
			line++
		}

		matrixInt := make([][]int, 5)
		for i := range matrixString {
			matrixInt[i] = make([]int, 5)
			// This is needed because there are some empty strings in matrixString
			intIdx := 0
			for j := range matrixString[i] {
				if matrixString[i][j] == "" {
					continue
				}

				matrixInt[i][intIdx], _ = strconv.Atoi(matrixString[i][j])
				intIdx++
			}
		}
		matrices = append(matrices, matrixInt)
	}

	return randNums, matrices
}

func main() {
	day4Part1()
	day4Part2()
}
