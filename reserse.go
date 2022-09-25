package main

import (
	"fmt"
	"reflect"
)

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

func drawBoard(board [fieldRow][fieldCol]string) {
	fmt.Println("   1  2  3  4  5  6  7  8")
	fmt.Println(" +------------------------+")
	for i, row := range board {
		fmt.Printf("%v|", i+1)
		for _, val := range row {
			fmt.Print(" ", val, " ")
		}
		fmt.Printf("|%v\n", i+1)
	}
	fmt.Println(" +------------------------+")
	fmt.Println("   1  2  3  4  5  6  7  8")
}

func playGame(tab [][8]string) {

	tab[3][3] = "X"
	tab[4][4] = "X"
	tab[3][4] = "O"
	tab[4][3] = "O"
}

func main() {
	board := getNewBoard()
	tab := board[:][:]
	fmt.Println(reflect.TypeOf(tab))
	playGame(tab)
	drawBoard(board)
	fmt.Println(board[2][2])
}
