package main

import (
	"fmt"
	"math/rand"
	"time"
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
	//init game
	turn := "player"
	board := getNewBoard()
	board[3][3] = "X"
	board[4][4] = "X"
	board[3][4] = "O"
	board[4][3] = "O"

	// Game cicle
	for {
		playerValidMoves := getValidMoves(&board, playerTile)
		fmt.Println(playerTile, playerValidMoves)
		drawBoard(&board)
		computerValidMoves := getValidMoves(&board, computerTile)
		fmt.Println(computerTile, computerValidMoves)
		drawBoard(&board)

		if playerValidMoves == nil && computerValidMoves == nil {
			return board
		} else if turn == "player" {
			if playerValidMoves != nil {
				playerMove := getRandomMove(&board, playerTile)
				fmt.Println("player", playerTile, playerMove)
				makeMove(&board, playerTile, playerMove)
				turn = "computer"
			}
		} else if turn == "computer" {
			if computerValidMoves != nil {
				comupterMove := getRandomMove(&board, computerTile)
				fmt.Println("computer", computerTile, comupterMove)
				makeMove(&board, computerTile, comupterMove)
				turn = "player"
			}
		}

	}
}

// Return the move that flips the least number of tiles.
func getRandomMove(board *[fieldRow][fieldCol]string, tile string) [2]int {
	possibleMoves := getValidMoves(board, tile)
	rand.Seed(time.Now().UnixMicro())
	return possibleMoves[rand.Intn(len(possibleMoves))]
}

// Place the tile on the boardat xstart, ystart and flip
// any if opponent's pieces.
// Return False if this is an invalid move; True if it is valid.
func makeMove(board *[fieldRow][fieldCol]string, tile string, move [2]int) bool {
	tilesToFlip := calcValidMoves(board, tile, move[0], move[1])
	if tilesToFlip == nil {
		return false
	}
	board[move[0]][move[1]] = tile
	for cell := range tilesToFlip {
		row := tilesToFlip[cell][0]
		col := tilesToFlip[cell][1]
		board[row][col] = tile
	}
	return true
}

func getPlayerMove(board *[fieldRow][fieldCol]string, tile string) [2]int {
	panic("unimplemented")
}

// Return a list of [x, y] lists of validmoves
// for the given player on the board.
func getValidMoves(board *[fieldRow][fieldCol]string, tile string) [][2]int {
	var validMoves [][2]int
	// var cellValue []string
	for row, rowVal := range board {
		for col := range rowVal {
			tilesToFlip := calcValidMoves(board, tile, row, col)
			if tilesToFlip != nil {
				validMoves = append(validMoves, [2]int{row, col})
			}
			// validMoves = append(validMoves, [2]int{row, col})
			// cellValue = append(cellValue, val)
		}
	}
	return validMoves
}

func calcValidMoves(board *[fieldRow][fieldCol]string, tile string, startRow, startCol int) [][2]int {
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
	return "X", "O"
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
