package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Piece int

type GameError string

func (e GameError) Error() string {
	return string(e)
}

const (
	Empty Piece = iota
	Red
	Blue
	Draw
)

type Board [][]Piece

type Move int

func main() {
	fmt.Println("Welcome to Connect Four!")
	winner := runGame(newBoard())
	if winner == Red {
		fmt.Println("Red Wins!")
	} else if winner == Blue {
		fmt.Println("Blue Wins!")
	} else {
		fmt.Println("Draw!")
	}
}

func printBoard(board Board) {
	printableBoard := make([][]string, 6)
	for i := range printableBoard {
		printableBoard[i] = make([]string, 7)
	}

	for i := range board {
		for j := range board[i] {
			textver := " "
			if board[i][j] == Red {
				textver = "R"
			}
			if board[i][j] == Blue {
				textver = "B"
			}
			if board[i][j] == Empty {
				textver = "_"
			}
			printableBoard[i][j] = textver
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "1", "2", "3", "4", "5", "6", "7"})
	t.AppendRows([]table.Row{
		{"A", printableBoard[0][0], printableBoard[0][1], printableBoard[0][2], printableBoard[0][3], printableBoard[0][4], printableBoard[0][5], printableBoard[0][6]},
		{"B", printableBoard[1][0], printableBoard[1][1], printableBoard[1][2], printableBoard[1][3], printableBoard[1][4], printableBoard[1][5], printableBoard[1][6]},
		{"C", printableBoard[2][0], printableBoard[2][1], printableBoard[2][2], printableBoard[2][3], printableBoard[2][4], printableBoard[2][5], printableBoard[2][6]},
		{"D", printableBoard[3][0], printableBoard[3][1], printableBoard[3][2], printableBoard[3][3], printableBoard[3][4], printableBoard[3][5], printableBoard[3][6]},
		{"E", printableBoard[4][0], printableBoard[4][1], printableBoard[4][2], printableBoard[4][3], printableBoard[4][4], printableBoard[4][5], printableBoard[4][6]},
		{"F", printableBoard[5][0], printableBoard[5][1], printableBoard[5][2], printableBoard[5][3], printableBoard[5][4], printableBoard[5][5], printableBoard[5][6]},
	})
	t.SetStyle(table.StyleColoredBright)
	t.Render()

}

func newBoard() Board {
	board := make(Board, 6)
	for i := range board {
		board[i] = make([]Piece, 7)
	}
	return board
}

func testBoard() Board {
	board := newBoard()
	board[5][5] = Blue
	board[5][6] = Red
	return board
}

func runGame(board Board) Piece {
	turn := 0
	winner := Empty

	for winner == Empty {
		if turn == 0 {
			// Red's turn
			fmt.Println("Red's turn: ")
			fmt.Println("Here's the board!")
			printBoard(board)

			humanTurn(board, Red)
		} else {
			// blue's turn
			fmt.Println("Blue's turn: ")
			fmt.Println("Here's the board!")
			printBoard(board)

			humanTurn(board, Blue)
		}
		turn = (turn + 1) % 2
		winner = checkWinner(board)
	}

	return winner
}

func humanTurn(board Board, piece Piece) Piece {
	pc, err := makeMove(board, piece)
	for err != nil {
		fmt.Println(err)
		pc, err = makeMove(board, piece)
	}
	// return piece placed
	return pc
}

func checkWinner(board Board) Piece {
	rows := 6
	cols := 7
	emptyFound := false

	// vertical check
	for col := 0; col < cols; col++ {
		for row := 0; row < rows-3; row++ {
			if board[row][col] == Empty {
				emptyFound = true
				continue
			}
			if board[row][col] == board[row+1][col] &&
				board[row][col] == board[row+2][col] &&
				board[row][col] == board[row+3][col] {
				return board[row][col]
			}
		}
	}

	// horizontal check
	for row := 0; row < rows; row++ {
		for col := 0; col < cols-3; col++ {
			if board[row][col] == Empty {
				emptyFound = true
				continue
			}
			if board[row][col] == board[row][col+1] &&
				board[row][col] == board[row][col+2] &&
				board[row][col] == board[row][col+3] {
				return board[row][col]
			}
		}
	}

	// bottom-up diagonal check
	for row := 0; row < rows-3; row++ {
		for col := 0; col < cols-3; col++ {
			if board[row][col] == Empty {
				emptyFound = true
				continue
			}
			if board[row][col] == board[row+1][col+1] &&
				board[row][col] == board[row+2][col+2] &&
				board[row][col] == board[row+3][col+3] {
				return board[row][col]
			}
		}
	}

	// top-down diagonal check
	for row := 3; row < rows; row++ {
		for col := 0; col < cols-3; col++ {
			if board[row][col] == Empty {
				emptyFound = true
				continue
			}
			if board[row][col] == board[row-1][col+1] &&
				board[row][col] == board[row-2][col+2] &&
				board[row][col] == board[row-3][col+3] {
				return board[row][col]
			}
		}
	}

	if emptyFound == false {
		return Draw
	} else {
		return Empty
	}
}

func makeMove(board Board, piece Piece) (Piece, error) {
	column, err := tryMove()
	if err != nil {
		return Empty, err
	}

	row, err := checkFull(board, column)
	if err != nil {
		return Empty, err
	}

	board[row][column] = piece
	return piece, nil
}

func tryMove() (int, error) {
	var column int
	fmt.Println("Enter a column to place your piece: ")
	fmt.Scan(&column)
	if column < 0 || column > 7 {
		return 0, GameError("Invalid column")
	}
	return column - 1, nil
}

func checkFull(board Board, column int) (int, error) {
	for i := 6; i > 0; i-- {
		if board[i-1][column] == Empty {
			return i - 1, nil
		}
	}
	return -1, GameError("Column is full")
}
