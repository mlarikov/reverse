package main

import "fmt"

const fieldRow int = 8
const fieldCol int = 8

func getNewBoard() [fieldRow][fieldCol]string {
	var board [fieldRow][fieldCol]string

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			board[row][col] = " "
		}
	}

	return board
}

func main() {
	board := getNewBoard()
	fmt.Println(board)
}
