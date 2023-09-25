package main

import "structs/pkg/game"

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

		// Check for a win condition or a tie (not shown in this example).
		// You can add your win/tie condition check here and break out of the loop if needed.

		// Alternate between players.
		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}
	}
}
