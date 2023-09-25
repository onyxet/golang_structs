package main

import (
	"fmt"
	"structs/pkg/game"
)

func main() {
	var takenSymbol string
	board := game.NewBoard()
	game.Greetings()
	player1 := game.NewPlayer(&takenSymbol)
	game.GreetingsPlayer(player1)
	player2 := game.NewPlayer(&takenSymbol)
	game.GreetingsPlayer(player2)
	currentPlayer := player1
	for {
		board.Display()
		currentPlayer.Move(board)
		if board.CheckWinner(currentPlayer) {
			fmt.Printf("%v is the Winner!\n", currentPlayer.Name)
			break
		}
		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}
	}
}
