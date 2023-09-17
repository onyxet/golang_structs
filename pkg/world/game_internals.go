package world

import (
	"fmt"
	"github.com/bengadbois/flippytext"
	"math/rand"
	"os"
)

// WelcomeMsg  helper
func WelcomeMsg() {
	flippytext.New().Write("Welcome to the game very player!\n" +
		"Today, we're going to go through small quest happened in abandoned asylum\n" +
		"You find yourself inside dark hall of old building\n" +
		"Last thing you remember was your house and how you went sleep\n" +
		"Now you alone in the dark lying in your pigama\n" +
		"You need to escape but main door closed...\n" +
		"It's raining outside and there is no sun, seems it's night...\n" +
		"Damn, you need to find the exit and find out what happened as soon as possible...\n" +
		"Be aware of stamina level. You can't escape if you're drained\n")
}

func WelcomePlayerMsg() {

}

// Helper function to generate random boolean value
func randBool() bool {
	return rand.Intn(2) == 1
}

func OpenDoorMsg() {
	fmt.Println("You slowly opening the door...")
}

func EndOfTheGame() {
	fmt.Println("Maybe next time...")
	os.Exit(0)
}

func RoomChoice(rooms []Room) int {
	var choice int
	fmt.Printf("You see %v rooms in a front. Please choose which one do you need\n", len(rooms)-1)
	for i, v := range rooms {
		if i == 0 {
			fmt.Printf("%v. Door %v - this is programmers thriller, so door can be 0 :)\n", i, i)
			continue
		}
		if v.Visited == true {
			fmt.Printf("%v. Door %v and you have visited it\n", i, i)
		}
		if v.Visited == false {
			fmt.Printf("%v. Door %v and you haven't visited it yet\n", i, i)
		}
	}
	//fmt.Printf("%v. Return to hall\n", len(rooms))
	fmt.Printf("%v. You can always exit the game...\n", len(rooms))
	fmt.Scanln(&choice)
	if choice == len(rooms) {
		EndOfTheGame()
	}
	//if choice == len(rooms) {
	//	break
	//}
	return choice
}
