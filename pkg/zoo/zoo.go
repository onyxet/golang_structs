package zoo

import "fmt"

type Enclosure struct {
	Size     float32
	Material string
}

type Cage struct {
	Enclosure
	Purpose string
	Animals []AnimalInterface
}

func (c Cage) Open() {
	fmt.Printf("\nCage for %v opened. And it's quite dangerous!\n", c.Purpose)
}

func (c Cage) Closed(a AnimalsInSlice) {
	fmt.Println("Cage closed. Next animals are inside:")
	for _, animal := range a.Animals {
		fmt.Printf("Name: %s\n", animal.GetName())
	}
}
