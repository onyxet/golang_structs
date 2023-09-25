package game

import "fmt"

func Greetings() {
	fmt.Println("Welcome to Tic Tac Toe!\nThis game is for two persons so please choose either X or O and start the game!\n")
}

func GreetingsPlayer(p *Player) {
	fmt.Printf("Player %s gaming with %v\n", p.Name, p.Symbol)
}
