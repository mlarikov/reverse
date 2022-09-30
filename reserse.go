package main

import (
	"fmt"
)

const fieldRow int = 8
const fieldCol int = 8

func getNewBoard() [fieldRow][fieldCol]string {
	var board [fieldRow][fieldCol]string

	for row, rowVal := range board {
		for col := range rowVal {
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

	// turn := "player"
	board := getNewBoard()
	board[3][3] = "X"
	board[4][4] = "X"
	board[3][4] = "O"
	board[4][3] = "O"
	for {
		playerValidMoves := getValidMoves(&board, playerTile)
		fmt.Println(playerValidMoves)
		return board
	}
}

func getValidMoves(board *[fieldRow][fieldCol]string, tile string) [][2]int {
	var validMoves [][2]int
	// var cellValue []string
	for row, rowVal := range board {
		for col := range rowVal {
			tilesToFlip := calcValidMoves(*board, tile, row, col)
			if tilesToFlip != nil {
				validMoves = append(validMoves, [2]int{row, col})
			}
			// validMoves = append(validMoves, [2]int{row, col})
			// cellValue = append(cellValue, val)
		}
	}

	return validMoves
}

func calcValidMoves(board [fieldRow][fieldCol]string, tile string, startRow, startCol int) [][2]int {
	var tilesToFlip [][2]int

	if board[startRow][startCol] != " " || !isOnBoard(startRow, startCol) {
		return tilesToFlip
	}

	var otherTile string
	if tile == "X" {
		otherTile = "O"
	} else {
		otherTile = "X"
	}

	directions := [][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1},
		{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	for direction := range directions {
		rowDir := directions[direction][0]
		colDir := directions[direction][1]
		row, col := startRow, startCol
		row += rowDir //First step in the y direction
		col += colDir //First step in the x direction
		for isOnBoard(row, col) && board[row][col] == otherTile {
			// Keep moving in this direction.
			row += rowDir
			col += colDir
			if isOnBoard(row, col) && board[row][col] == tile {
				// There are pieces to flip over. Go in the reverse
				// direction until we reach the original space,
				// noting all the tiles along the way.
				for {
					row -= rowDir
					col -= colDir
					if row == startRow && col == startCol {
						break
					}
					tilesToFlip = append(tilesToFlip, [2]int{row, col})
				}
			}
		}
	}
	return tilesToFlip
}

func isOnBoard(row, col int) bool {
	return row <= fieldRow-1 && row >= 0 && col >= 0 && col <= fieldCol-1
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
