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
	printBoard(board)
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
		prograss := false

		if solveSingles(board, options) {
			prograss = true
		}
		if solveHiddenSingles(board, options) {
			prograss = true
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

		if !prograss {
			fmt.Println("could not solve any further")
			return
		}
	}
}

func solveSingles(board [][]int, options [][]map[int]bool) bool {
	prograss := false
	for i := range board {
		for j := range board[i] {
			if board[i][j] != 0 {
				continue
			}

			refreshOption(board, options, Coordinates{ x: i, y: j})
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
				return prograss
			}

			if avilableOptions == 1 {
				setBoardAndRefresh(board, options,
				Coordinates{x: i, y: j}, avilableOption)
				prograss = true
			}
		}
	}

	return prograss
}

func solveHiddenSingles(board [][]int, options [][]map[int]bool) bool {
	prograss := false
	for rows := range options {
		occurence := make(map[int][]Coordinates, 9)

		for columns := range options {
			markHiddenSingles(board, options, occurence, rows, columns)
		}

		if findHiddenSingles(board, options, occurence) {
			prograss = true
		}
	}

	for columns := range options {
		occurence := make(map[int][]Coordinates, 9)

		for rows := range options {
			markHiddenSingles(board, options, occurence, rows, columns)
		}

		if findHiddenSingles(board, options, occurence) {
			prograss = true
		}
	}

	for x := 0; x < 8; x += 3 {
		for y := 0; y < 8; y += 3 {
			occurence := make(map[int][]Coordinates, 9)

			blockRow := (x / 3) * 3
			blockCol := (y / 3) * 3

			for r:= blockRow; r < blockRow+3; r++ {
				for c:= blockCol; c < blockCol+3; c++ {
					markHiddenSingles(board, options, occurence, r, c)
				}
			}

			if findHiddenSingles(board, options, occurence) {
				prograss = true
			}
		}
	}


	return prograss
}

func markHiddenSingles(board [][]int, options[][]map[int]bool, occurence map[int][]Coordinates, rows, columns int) {
	for key, value := range options[rows][columns] {
		if value {
			occurence[key] = append(occurence[key], Coordinates {
				x: rows,
				y: columns,
			})
		}
	}
}

func findHiddenSingles(board [][]int, options[][]map[int]bool, occurence map[int][]Coordinates) bool {
	prograss := false
	for key, value := range occurence {
		if len(value) == 1 {
			setBoardAndRefresh(board, options,
			Coordinates{x: value[0].x, y: value[0].y}, key)
			prograss = true
		}
	}

	return prograss
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

func setBoardAndRefresh(board [][]int, options [][]map[int]bool, coord Coordinates, val int) {
	x, y := coord.x, coord.y

	board[x][y] = val

	for columns := range board {
		options[x][columns][val] = false
	}

	for rows := range board {
		options[rows][y][val] = false
	}

	blockRow := (x / 3) * 3
	blockCol := (y / 3) * 3

	for r := blockRow; r < blockRow+3; r++ {
		for c := blockCol; c < blockCol+3; c++ {
			options[r][c][val] = false
		}
	}
}

func printBoard(board [][]int) {
	for i := range board {
		strNums := make([]string, len(board[i]))
		for j, n := range board[i] {
			strNums[j] = strconv.Itoa(n)
		}

		fmt.Printf("%s\n", strings.Join(strNums, " "))
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
