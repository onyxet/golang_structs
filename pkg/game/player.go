package game

import "fmt"

type Player struct {
	Name   string
	Symbol string
}

func NewPlayer(existingSymbol *string) *Player {
	var name, symbol string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)
	if name == "" {
		panic("Name can't be null")
	}

	for {
		fmt.Println("Please enter your symbol (X or O):")
		fmt.Scanln(&symbol)
		if symbol != "X" && symbol != "O" {
			fmt.Println("Symbol must be either 'X' or 'O'")
		}
		if symbol == *existingSymbol {
			fmt.Printf("Symbol %v is already taken", *existingSymbol)
			continue
		}
		break
	}

	*existingSymbol = symbol
	return &Player{
		Name:   name,
		Symbol: symbol,
	}
}

func (p *Player) Move(b *Board) {
	var movement int
	fmt.Println("Please select your movement (1-9):")
	fmt.Scanln(&movement)

	row := (movement - 1) / 3
	col := (movement - 1) % 3

	if b.Rows[row][col] != "[]" {
		fmt.Println("Invalid movement. The cell is already occupied.")
		return
	}
	b.Rows[row][col] = "[" + p.Symbol + "]"
}
