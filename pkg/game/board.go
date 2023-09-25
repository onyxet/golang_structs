package game

import "fmt"

type Board struct {
	Rows   [][]string
	Winner Player
	Count  int
}

func NewBoard() *Board {
	return &Board{
		Rows: [][]string{
			{"[]", "[]", "[]"},
			{"[]", "[]", "[]"},
			{"[]", "[]", "[]"},
		},
	}
}

func (b *Board) Display() {
	for _, row := range b.Rows {
		fmt.Println(row)
	}
}
