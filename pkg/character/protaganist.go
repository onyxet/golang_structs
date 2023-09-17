package character

import (
	"fmt"
	"github.com/bengadbois/flippytext"
	"os"
	"structs/pkg/inventory"
	"structs/pkg/world"
)

type Protagonist struct {
	Name     string
	Surname  string
	Stamina  int
	Tool     inventory.Tool
	isHunted bool
}

func NewProtagonist() (protagonist *Protagonist) {
	var name, surname string
	for name == "" {
		fmt.Println("Oh, I remember my name... It's...")
		fmt.Scanln(&name)
	}
	for surname == "" {
		fmt.Println("And surname...")
		fmt.Scanln(&surname)
	}

	return &Protagonist{
		Name:    name,
		Surname: surname,
		Stamina: 100,
	}
}

func CheckStamina(p *Protagonist) {
	if p.Stamina <= 0 {
		fmt.Printf("You have %v stamina\nYou can save your life being drained.\n You died...\n", p.Stamina)
		world.EndOfTheGame()
	}
	fmt.Printf("You have %v stamina\n", p.Stamina)
}

func (p *Protagonist) WhatsInRoom(r *world.Room) {
	// There might be a loop here! Since we're going to check all things around
	var choice int
	if p.isHunted && r.HasWardrobe {
		fmt.Println("Room has wardrobe, you inside, waiting...")
		CheckStamina(p)
		fmt.Println("Get out from wardrobe?\n1. Yes\n2. No")
		fmt.Scanln(&choice)
		for choice >= 2 {
			fmt.Println("Waiting... but it's quite nervous")
			p.Stamina = -5
			CheckStamina(p)
		}
		if choice == 1 {
			fmt.Println("Let's see what else do we have here...")
		}
	}
	if r.HasBloodyTraces {
		fmt.Println("Seems like somebody was killed there or so.\nBlood traces over all room")
		p.Stamina -= 10
		CheckStamina(p)
	}
	if r.HasBed == true {
		fmt.Printf("You see some bed in a room... You have %v stamina level\n", p.Stamina)
		fmt.Println("Want to have some rest?\n1. Yes\n2. No")
		CheckStamina(p)
		if p.Stamina < 70 {
			fmt.Scanln(&choice)
			if choice == 1 {
				fmt.Println("For how many hours?\n1. 1h\n2. More than 1h\n")
				fmt.Scanln(&choice)
				if choice == 2 {
					fmt.Println("It's always a bad idea have a good sleep and abandoned asylum.\nYou were killed by somebody...")
					world.EndOfTheGame()
				}
				if choice == 1 {
					p.Stamina = 100
					fmt.Printf("Your stamina now %v \nLet's see around what's in the room\n", p.Stamina)
				}
				if choice == 2 {
					fmt.Println("Let's see what else do we have here...")
				}

			}
		}
		if p.Stamina > 70 {
			fmt.Scanln(&choice)
			if choice == 1 {
				fmt.Println("C'mon! It's not the time! I'm fresh!\nLet's see around what's in the room\n")
			}
			if choice == 2 {
				fmt.Println("Good choice\n")
			}
		}

	}
	if r.HasBarsOnWindows {
		fmt.Println("o! There is a window!\n1. Go to the window and try to exit\n2.Ignore it")
		fmt.Scanln(&choice)
		if choice == 1 {
			fmt.Println("Foolish! Of course there is a bars and you can't broken them")
			p.Stamina -= 10
		}
		if choice == 2 {
			fmt.Println("Never-mind than")
		}
		CheckStamina(p)

	}
	if r.HasWardrobe {
		fmt.Println("There is wardrobe...\nShould you hide?\n1. Yes\n2. No")
		fmt.Scanln(&choice)
		if choice == 1 {
			fmt.Println("You're in wardrobe\n Waiting... but it's quite nervous...\n1. Want to exit?\n2. Stay inside")

			for choice == 1 {
				fmt.Println("Waiting... but it's quite nervous.\n1. Stay inside \n2. Exit")
				fmt.Scanln(&choice)
				//if choice == 1 {
				//
				//}
				p.Stamina -= 5
				CheckStamina(p)
			}
		}
		if choice == 2 {
			fmt.Println("Let's see what we have here...")
		}
		CheckStamina(p)
	}

}

// OpenDoor is quite straight logic of opening any room in the asylum.
func (p *Protagonist) OpenDoor(room *world.Room, choice int) {
	if room.HasLockedDoor == true {
		flippytext.New().Write("\nYou trying to open the door but it's locked. You feel your hands starts shaking\n")
		room.Visited = true
	}
	if room.IsExit == true {
		fmt.Println("You won the price! Freedom. Run Forest run!")
		os.Exit(0)
	}
	if room.HasEnemyInside == true {
		if p.Stamina > 70 && p.Tool != nil {
			world.OpenDoorMsg()
			fmt.Printf("You see the man inside!\nYou have some stamina and you %v with you!", p.Tool)
		}
		if p.Stamina < 70 && p.Tool != nil {
			world.OpenDoorMsg()
			fmt.Printf("You see the man inside!\nYou have %v\\% stamina and you %v with you!", p.Stamina, p.Tool)
			fmt.Println("This is it! What would you choose?\n 1. Fight\n2. Run!")
			fmt.Scanln(&choice)
			if choice == 2 {
				// this should be refactored to method of rooms left
				fmt.Println("You see rooms left")
			}
			if choice == 1 {
				fmt.Println("You fight as a hero but you had no enough power for it, this is the end\n")
				world.EndOfTheGame()
			}
		}

	}
	room.Visited = true
	p.WhatsInRoom(room)
}
