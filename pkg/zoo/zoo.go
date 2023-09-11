package zoo

import "fmt"

type Enclosure struct {
	Size     float32
	Material string
}

type Cage struct {
	Enclosure
	Purpose string
}

func (c Cage) Open() {
	fmt.Printf("\nCage for %v opened. And it's quite dangerous!\n", c.Purpose)
}

func (c Cage) Close() {
	println("Cage closed. Animals are safe now")
}
