package world

import (
	"fmt"
	"math/rand"
	"time"
)

type Asylum struct {
	Name    string
	RoomNum int
}

type Room struct {
	HasBed           bool
	HasWindow        bool
	HasBarsOnWindows bool
	HasWardrobe      bool
	HasBloodyTraces  bool
	HasLight         bool
	HasLockedDoor    bool
	HasEnemyInside   bool
	Visited          bool
	IsExit           bool
}

func NewAsylum() *Asylum {
	rand.Seed(time.Now().UnixNano())
	roomNum := rand.Intn(15) + 5 // Ensure we have not less than 5
	var name string
	for name == "" {
		fmt.Println("You see the sign...\nIs that the name of the place?\nPlease enter what do you see:")
		fmt.Scanln(&name)
	}
	return &Asylum{
		Name:    name,
		RoomNum: roomNum,
	}
}

func NewRooms(roomNum int) []Room {
	rand.Seed(time.Now().UnixNano())
	rooms := make([]Room, roomNum)
	for i := 0; i < roomNum; i++ {
		rooms[i] = Room{
			HasBed:           randBool(),
			HasWindow:        randBool(),
			HasBarsOnWindows: randBool(),
			HasWardrobe:      randBool(),
			HasBloodyTraces:  randBool(),
			HasLight:         randBool(),
		}
		rooms[cap(rooms)-1] = Room{HasEnemyInside: true}
		rooms[cap(rooms)-2] = Room{HasLockedDoor: true}
		rooms[cap(rooms)-3] = Room{IsExit: true}
	}
	return rooms
}
