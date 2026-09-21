package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
