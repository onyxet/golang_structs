package main

import (
	"structs/pkg/character"
	"structs/pkg/world"
)

func main() {
	//var choice int
	world.WelcomeMsg()
	asylum := world.NewAsylum()
	pers := character.NewProtagonist()

	rooms := world.NewRooms(asylum.RoomNum)
	for {
		choice := world.RoomChoice(rooms)
		pers.OpenDoor(&rooms[choice], choice)
	}

}
