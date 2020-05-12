package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var board = [][]string{
	{" ", " ", " "},
	{" ", " ", " "},
	{" ", " ", " "},
}

var turns = 0

func turnHuman() {
	success := false
	for !success {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Ваш ход (00-22):")
		text, _ := reader.ReadString('\n')
		s_t := strings.Split(text, "")
		i, _ := strconv.Atoi(s_t[0])
		j, _ := strconv.Atoi(s_t[1])

		if board[i][j] == " " {
			board[i][j] = "A"
			success = true
		} else {
			fmt.Println("Занято")
		}
	}
}

func minimax(board [][]string, depth int, isMax bool) int {
	result := checkWin()

	score := 0

	if result == "A" {
		return -10
	} else if result == "B" {
		return 10
	} else if result == "tie" {
		return 0
	}

	if isMax {
		bestScore := -999999
		for i, row := range board {
			for j, _ := range row {
				if board[i][j] == " " {
					board[i][j] = "B"
					score = minimax(board, depth+1, false)
					board[i][j] = " "
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := 999999
		for i, row := range board {
			for j, _ := range row {
				if board[i][j] == " " {
					board[i][j] = "A"
					score = minimax(board, depth+1, true)
					board[i][j] = " "
					if score < bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	}

}

func turnComputer() {
	bestScore := -999999
	bestMove := []int{0, 0}
	for i, row := range board {
		for j, _ := range row {
			if board[i][j] == " " {
				board[i][j] = "B"
				score := minimax(board, 0, false)
				board[i][j] = " "
				if score > bestScore {
					bestScore = score
					bestMove[0] = i
					bestMove[1] = j
				}
			}
		}
	}
	board[bestMove[0]][bestMove[1]] = "B"
}

func equals3(a string, b string, c string) bool {
	return (a == b && b == c && a != " ")
}

func checkWin() string {

	winner := ""

	for i, _ := range board {
		if equals3(board[i][0], board[i][1], board[i][2]) {
			winner = board[i][0]
		}
	}

	for i, _ := range board {
		if equals3(board[0][i], board[1][i], board[2][i]) {
			winner = board[0][i]
		}
	}

	if equals3(board[0][0], board[1][1], board[2][2]) {
		winner = board[0][0]
	}

	if equals3(board[2][0], board[1][1], board[0][2]) {
		winner = board[2][0]
	}

	openSpots := 0

	for i, row := range board {
		for j, _ := range row {
			if board[i][j] == " " {
				openSpots++
			}
		}
	}

	if winner == " " && openSpots == 0 {
		return "tie"
	} else {
		return winner
	}
}

func printBoard() {
	for _, row := range board {
		fmt.Println(row)
	}
}

func main() {
	winner := ""
	for true {
		fmt.Println("Компьютер походил:")
		turnComputer()
		printBoard()
		turns++
		winner = checkWin()

		if winner != "" {
			break
		}

		turnHuman()
		printBoard()
		turns++
		winner = checkWin()

		if winner != "" {
			break
		}
	}
	fmt.Println(winner)
}
