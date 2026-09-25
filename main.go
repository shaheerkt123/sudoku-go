package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)


type Coordinates struct {
	x int
	y int
}

func main() {
	rows, column := 9, 9
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, column)
	}
	getProblem(board)
	solveBoard(board)
	fmt.Println(board)
}


func solveBoard(board [][]int) {
	options := make([][]map[int]bool, 9)
	for i := range options {
		options[i] = make([]map[int]bool, 9)
		for j := range options[i] {
			options[i][j] = make(map[int]bool, 9)
			for k := 1; k <= 9;k++ {
				options[i][j][k] = true
			}
		}
	}

	fmt.Println("lets solve")
	for {
		progress := false
		for i := range board {
			for j := range board[i] {
				if board[i][j] != 0 {
					continue
				}

				for columns := range board {
					options[i][j][board[i][columns]] = false
				}

				for rows := range board {
					options[i][j][board[rows][j]] = false
				}

				quadrentRow := (i / 3) * 3
				quadrentCol := (j / 3) * 3

				for r := quadrentRow; r < quadrentRow+3; r++ {
					for c := quadrentCol; c < quadrentCol+3; c++ {
						options[i][j][board[r][c]] = false
					}
				}

				avilableOptions := 0
				avilableOption := 0
				for key, value := range options[i][j] {
					if value {
						avilableOptions++
						avilableOption = key
					}
				}
				if avilableOptions < 1 {
					fmt.Println("invalid problem")
					return
				}

				if avilableOptions == 1 {
					board[i][j] = avilableOption
					progress = true
				}
			}
		}

		solved := true
		for i := range board {
			if slices.Contains(board[i], 0) {
				solved = false
			}
		}

		if solved {
			fmt.Println("Solved")
			return
		}

		if !progress {
			fmt.Println("could not solve any further")
			return
		}
	}
}

func refreshOption(board [][]int, options [][]map[int]bool, coord Coordinates) {
	x := coord.x
	y := coord.y

	if board[x][y] != 0 {
		return
	}

	for columns := range board {
		options[x][y][board[x][columns]] = false
	}

	for rows := range board {
		options[x][y][board[rows][y]] = false
	}

	blockRow := (x / 3) * 3
	blockCol := (y / 3) * 3

	for r := blockRow; r < blockRow+3; r++ {
		for c := blockCol; c < blockCol+3; c++ {
			options[x][y][board[r][c]] = false
		}
	}
}

func getProblem(board [][]int) {
	fmt.Println("Enter the problem:")
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 9 {
			fmt.Printf("Error: not 9 numbers\n")
			return
		}

		for j, val := range fields {
			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			board[i][j] = num
		}

		if i >= 8 {
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("stream error: %v\n", err)
	}
}
