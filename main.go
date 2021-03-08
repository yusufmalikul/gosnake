package main

import (
	"container/list"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type SnakeCord struct {
	PosI int
	PosJ int
}

type Food struct {
	PosI int
	PosJ int
}

const HEAD = "o"
const FOOD = "*"
const EMPTY = " "

var grow = false
var food = Food{}

func main() {

	clearScreen()

	var boardSize int

	fmt.Println("Welcome to snake game!")
	fmt.Println("Please input the board size.")
	fmt.Print("Board size: ")
	_, err := fmt.Scan(&boardSize)
	if err != nil {
		log.Fatal("must input number")
	}

	board := make([][]string, boardSize)
	for i := range board {
		board[i] = make([]string, boardSize)
	}

	// for generating random number
	rand.Seed(time.Now().UnixNano())

	// randomize snake initial position
	snakeBody := list.New()
	snakeBody.PushBack(SnakeCord{
		PosI: rand.Intn((boardSize-1)-0) + 0,
		PosJ: rand.Intn((boardSize-1)-0) + 0,
	})

	// randomize first food
	food = Food{
		PosI: rand.Intn((boardSize-1)-0) + 0,
		PosJ: rand.Intn((boardSize-1)-0) + 0,
	}

	// fill empty space and place snake head
	printBoard(board, snakeBody)

	var command, round, score int
	for {
		clearScreen()
		fmt.Println("round:", round, "score:", score)
		printBoard(board, snakeBody)
		sc1 := snakeBody.Front().Value.(SnakeCord)
		sc := snakeBody.Back().Value.(SnakeCord)
		fmt.Println("Front:", sc1.PosI, sc1.PosJ)
		fmt.Println("Back:", sc.PosI, sc.PosJ)
		fmt.Println("Length:", snakeBody.Len())
		fmt.Println("Grow:", grow)
		fmt.Println(food)
		command = getInstruction()
		if command == 0 {
			fmt.Println("Good bye!")
			os.Exit(0)
		}

		snakeBody = moveHead(board, snakeBody, command, &score)
		round++
	}

}

func moveHead(board [][]string, snakeBody *list.List, command int, score *int) *list.List {

	head, _ := snakeBody.Front().Value.(SnakeCord)

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
		clearScreen()
		printBoard(board, snakeBody)
		fmt.Println("Your head collided with the wall!")
		os.Exit(0)
	}

	grow = false
	for i := range board {
		for j := range board[i] {
			if i == currI && j == currJ && board[i][j] == FOOD {
				*score++
				grow = true
				food = Food{
					PosI: rand.Intn((len(board)-1)-0) + 0,
					PosJ: rand.Intn((len(board)-1)-0) + 0,
				}
			} else if i == currI && j == currJ && board[i][j] == HEAD {
				// snake collision with it's own body
				clearScreen()
				printBoard(board, snakeBody)
				fmt.Println("You eat your own body!")
				os.Exit(0)
			}
		}
	}

	snakeBody.PushFront(SnakeCord{
		PosI: currI,
		PosJ: currJ,
	})

	if !grow {
		e := snakeBody.Back()
		fmt.Println(snakeBody.Len())
		snakeBody.Remove(e)
		fmt.Println(snakeBody.Len())
	}

	return snakeBody
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func getInstruction() int {
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

func printBoard(board [][]string, snakeBody *list.List) {

	// fill empty space
	for i := range board {
		for j := range board[i] {
			board[i][j] = EMPTY
		}
	}

	// place food
	for i := range board {
		for j := range board[i] {
			if i == food.PosI && j == food.PosJ {
				board[i][j] = FOOD
			}
		}
	}

	// place snake position
	e := snakeBody.Front()
	for e != nil {
		body, _ := e.Value.(SnakeCord)
		board[body.PosI][body.PosJ] = HEAD
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
