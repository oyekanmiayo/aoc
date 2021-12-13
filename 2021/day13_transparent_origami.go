package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		Place items in a queue
		Place instructions in a an array?
		For each x
			-> if x of point is > x of fold, new x = largest x seen so far - (x of point + 1) + (x of fold - half of largest x seen so far)
		For each y
			-> if y of point > y of fold, new y = largest y seen so far - (y of point + 1) + (y of fold - half of largest y seen so far)
	*/

	// Go over each line. Have a conditional extract instructions etc
	file, _ := os.Open("2021/day13_transparent_origami.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	queue := make([][]int, 0)
	instructions := make([][]string, 0)

	instructionsSeen := false
	xLen := math.MinInt32
	yLen := math.MinInt32

	for scanner.Scan() {
		if scanner.Text() == "" {
			instructionsSeen = true
			continue
		}

		if !instructionsSeen {
			positionStr := scanner.Text()
			posSlice := strings.Split(positionStr, ",")
			x, _ := strconv.Atoi(posSlice[0])
			y, _ := strconv.Atoi(posSlice[1])

			if x > xLen {
				xLen = x
			}

			if y > yLen {
				yLen = y
			}

			queue = append(queue, []int{x, y})
		} else {
			instruction := scanner.Text()
			instruction = strings.ReplaceAll(instruction, "fold along ", "")
			instrSlice := strings.Split(instruction, "=")
			instructions = append(instructions, instrSlice)
		}

	}

	// Increment because current calculations are 0-indexed
	xLen++
	yLen++

	for _, instr := range instructions {

		seen := make(map[int]map[int]bool)
		currQueueLen := len(queue)

		for i := 0; i < currQueueLen; i++ {
			pos := queue[0]
			queue[0] = nil
			queue = queue[1:]

			if instr[0] == "x" {
				foldX, _ := strconv.Atoi(instr[1])

				// can never be equal
				if pos[0] < foldX {
					if seen[pos[0]][pos[1]] {
						continue
					}

					queue = append(queue, pos)

					_, exist := seen[pos[0]]
					if !exist {
						seen[pos[0]] = make(map[int]bool)
					}
					seen[pos[0]][pos[1]] = true

				} else {
					newPos := make([]int, 2)
					newPos[0] = (foldX*2 + 1) - (pos[0]) - 1
					newPos[1] = pos[1]

					if seen[newPos[0]][newPos[1]] {
						continue
					}

					queue = append(queue, newPos)

					_, exist := seen[newPos[0]]
					if !exist {
						seen[newPos[0]] = make(map[int]bool)
					}
					seen[newPos[0]][newPos[1]] = true

				}
			}

			if instr[0] == "y" {
				foldY, _ := strconv.Atoi(instr[1])

				if pos[1] < foldY {
					if seen[pos[0]][pos[1]] {
						continue
					}

					queue = append(queue, pos)

					_, exist := seen[pos[0]]
					if !exist {
						seen[pos[0]] = make(map[int]bool)
					}
					seen[pos[0]][pos[1]] = true

				} else {
					newPos := make([]int, 2)
					newPos[0] = pos[0]
					newPos[1] = (foldY*2 + 1) - (pos[1]) - 1

					if seen[newPos[0]][newPos[1]] {
						continue
					}

					queue = append(queue, newPos)

					_, exist := seen[newPos[0]]
					if !exist {
						seen[newPos[0]] = make(map[int]bool)
					}
					seen[newPos[0]][newPos[1]] = true

				}
			}

		}

		if instr[0] == "x" {
			foldX, _ := strconv.Atoi(instr[1])
			xLen = foldX
		} else {
			foldY, _ := strconv.Atoi(instr[1])
			yLen = foldY
		}

		fmt.Println(xLen, yLen)
		printMatrix(queue, xLen, yLen)
		// fmt.Println("")
	}

	fmt.Println(len(queue), queue)
	// printMatrix(queue, xLen, yLen)

}

func printMatrix(queue [][]int, xLen, yLen int) {
	matrix := make([][]string, yLen)
	for i := range matrix {
		matrix[i] = make([]string, 0)
		for j := 0; j < xLen; j++ {
			matrix[i] = append(matrix[i], ".")
		}
	}

	for _, pos := range queue {
		matrix[pos[1]][pos[0]] = "#"
	}

	for i := range matrix {
		fmt.Println(matrix[i])
	}
}
