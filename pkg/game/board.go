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

func (b *Board) CheckWinner(player *Player) bool {
	for _, row := range b.Rows {
		if row[0] == "["+player.Symbol+"]" && row[1] == "["+player.Symbol+"]" && row[2] == "["+player.Symbol+"]" {
			return true
		}
	}
	for i := 0; i < 3; i++ {
		if b.Rows[0][i] == "["+player.Symbol+"]" && b.Rows[1][i] == "["+player.Symbol+"]" && b.Rows[2][i] == "["+player.Symbol+"]" {
			return true
		}
	}
	if b.Rows[0][0] == "["+player.Symbol+"]" && b.Rows[1][1] == "["+player.Symbol+"]" && b.Rows[2][2] == "["+player.Symbol+"]" {
		return true
	}

	if b.Rows[0][2] == "["+player.Symbol+"]" && b.Rows[1][1] == "["+player.Symbol+"]" && b.Rows[2][0] == "["+player.Symbol+"]" {
		return true
	}
	return false
}
