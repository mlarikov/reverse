package main

import (
	"fmt"
)

const fieldRow int = 8
const fieldCol int = 8

func getNewBoard() [fieldRow][fieldCol]string {
	var board [fieldRow][fieldCol]string

	for row, rowVal := range board {
		for col, _ := range rowVal {
			board[row][col] = " "
		}
	}
	// for row := 0; row < len(board); row++ {
	// 	for col := 0; col < len(board[row]); col++ {
	// 		board[row][col] = " "
	// 	}
	// }

	return board
}

func drawBoard(board *[fieldRow][fieldCol]string) {
	fmt.Println("   1 2 3 4 5 6 7 8")
	fmt.Println(" +-----------------+")
	for i, row := range board {
		fmt.Printf("%v| ", i+1)
		for _, val := range row {
			fmt.Printf("%v ", val)
		}
		fmt.Printf("|%v\n", i+1)
	}
	fmt.Println(" +-----------------+")
	fmt.Println("   1 2 3 4 5 6 7 8")
}

func playGame(playerTile string, computerTile string) [fieldRow][fieldCol]string {
	board := getNewBoard()
	board[3][3] = "X"
	board[4][4] = "X"
	board[3][4] = "O"
	board[4][3] = "O"
	return board
}

// Determine the score by counting the tiles. Return a dictionary
// with keys 'X' and 'O'.
func getScoreOfBoard(board [fieldRow][fieldCol]string) map[string]int {
	var xscore, oscore int
	for _, rowVal := range board {
		for _, val := range rowVal {
			if val == "X" {
				xscore++
			}
			if val == "O" {
				oscore++
			}

		}
	}
	return map[string]int{"X": xscore, "O": oscore}
}

func enterPlayerTile() (string, string) {
	return "X", "Y"
}

func main() {
	for {
		playerTile, computerTile := enterPlayerTile()
		finalBoard := playGame(playerTile, computerTile)
		drawBoard(&finalBoard)
		scores := getScoreOfBoard(finalBoard)
		fmt.Println(scores)
		break
	}
}
