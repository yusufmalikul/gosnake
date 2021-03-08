package main

import (
	"container/list"
	"fmt"
	"gosnake/pkg/common"
	"gosnake/pkg/model"
	"log"
	"math/rand"
	"time"
)

func main() {

	common.ClearScreen()

	var boardSize int

	fmt.Println("========== SNAKE ==========")
	fmt.Println("Welcome to snake game!")
	fmt.Println("Please input the board size.")
	for boardSize <= 1 {
		fmt.Print("Board size: ")
		_, err := fmt.Scan(&boardSize)
		if err != nil {
			log.Fatal("must input number")
		}
		if boardSize <= 1 {
			fmt.Println("must greater than 1")
		}
	}

	board := make([][]string, boardSize)
	for i := range board {
		board[i] = make([]string, boardSize)
	}

	// for generating random number
	rand.Seed(time.Now().UnixNano())

	// randomize snake initial position
	snakeBody := list.New()
	snakeBody.PushBack(model.SnakeCord{
		PosI: rand.Intn((boardSize-1)-0) + 0,
		PosJ: rand.Intn((boardSize-1)-0) + 0,
	})

	// randomize first food
	common.Food = model.Food{
		PosI: rand.Intn((boardSize-1)-0) + 0,
		PosJ: rand.Intn((boardSize-1)-0) + 0,
	}

	// fill empty space and place snake head
	common.PrintBoard(board, snakeBody)

	run(board, snakeBody)

}

func run(board [][]string, snakeBody *list.List) {
	var command, round, score int
	for {
		common.ClearScreen()
		fmt.Println("========== SNAKE ==========")
		fmt.Println("board:", len(board), "x", len(board))
		fmt.Println("round:", round)
		fmt.Println("score:", score)
		fmt.Println("length:", snakeBody.Len())
		fmt.Printf("head coord: i=%v j=%v\n", snakeBody.Front().Value.(model.SnakeCord).PosI, snakeBody.Front().Value.(model.SnakeCord).PosJ)
		fmt.Println()
		common.PrintBoard(board, snakeBody)
		command = common.GetInstruction()
		if command == 0 {
			fmt.Println("\nGood bye! Thanks for playing.")
			return
		}

		snakeBody = common.MoveHead(board, snakeBody, command, &score)
		round++
	}
}
