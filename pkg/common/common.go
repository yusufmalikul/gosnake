package common

import (
	"container/list"
	"fmt"
	"gosnake/pkg/model"
	"math/rand"
	"os"
)

const BODY = "O"
const FOOD = "*"
const EMPTY = " "

// global variable
var grow = false
var Food = model.Food{}

func MoveHead(board [][]string, snakeBody *list.List, command int, score *int) *list.List {

	head, _ := snakeBody.Front().Value.(model.SnakeCord)

	currI := head.PosI
	currJ := head.PosJ

	switch command {
	case 1:
		currI--
	case 2:
		currI++
	case 3:
		currJ--
	case 4:
		currJ++
	}

	if currI < 0 || currI > len(board)-1 || currJ < 0 || currJ > len(board)-1 {
		fmt.Println("\nGame Over!")
		fmt.Println("Your head collided with the wall!")
		os.Exit(0)
	}

	snakeBody.PushFront(model.SnakeCord{
		PosI: currI,
		PosJ: currJ,
	})

	grow = false

	if currI == Food.PosI && currJ == Food.PosJ {
		*score++
		grow = true

		// sometimes the food will be placed at the exact location with the
		// snake body, so we need to regenerate the food location if it happens

		for board[Food.PosI][Food.PosJ] == BODY || (Food.PosI == currI && Food.PosJ == currJ) {
			Food = model.Food{
				PosI: rand.Intn((len(board)-1)-0) + 0,
				PosJ: rand.Intn((len(board)-1)-0) + 0,
			}
		}

	} else if board[currI][currJ] == BODY {
		// snake collision with it's own body
		fmt.Println("\nGame Over!")
		fmt.Println("You eat your own body!")
		os.Exit(0)
	}

	if !grow {
		e := snakeBody.Back()
		fmt.Println(snakeBody.Len())
		snakeBody.Remove(e)
		fmt.Println(snakeBody.Len())
	}

	return snakeBody
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func GetInstruction() int {
	fmt.Println()
	fmt.Println("Where to go?")
	fmt.Println("1. Up")
	fmt.Println("2. Down")
	fmt.Println("3. Left")
	fmt.Println("4. Right")
	fmt.Println("0. Exit")
	fmt.Print("choice: ")
	var command int
	_, err := fmt.Scan(&command)
	if err != nil {
		fmt.Println(err)
	}
	return command
}

func PrintBoard(board [][]string, snakeBody *list.List) {

	// fill empty space
	for i := range board {
		for j := range board[i] {
			board[i][j] = EMPTY
		}
	}

	// place food
	for i := range board {
		for j := range board[i] {
			if i == Food.PosI && j == Food.PosJ {
				board[i][j] = FOOD
			}
		}
	}

	// place snake position
	e := snakeBody.Front()
	for e != nil {
		body, _ := e.Value.(model.SnakeCord)
		board[body.PosI][body.PosJ] = BODY
		e = e.Next()
	}

	// ----------- print board ---------- //

	// print top border
	fmt.Print(" ")
	for range board {
		fmt.Print("-")
	}

	// print side border and board content
	fmt.Println()
	for i := range board {
		for j := range board[i] {
			if j == 0 { // left side
				fmt.Print("|")
				fmt.Print(board[i][j])
			} else if j == len(board[i])-1 { // right side
				fmt.Print(board[i][j])
				fmt.Print("|")

			} else {
				fmt.Print(board[i][j]) // board content
			}
		}
		fmt.Println()
	}

	// print bottom border
	fmt.Print(" ")
	for range board {
		fmt.Print("-")
	}
	println()
}
