package zoo

import "fmt"

type Zookeeper struct {
	Human
	BachelorDegree bool
	ComputerSkills bool
}

func (z Zookeeper) OnDuty() {
	fmt.Printf("As a usual %v in his %vs on duty I'm watching UFC because animals are sleeping\n", z.Name, z.Age)
}

func (z Zookeeper) OffDuty() {
	fmt.Printf("As a humanbeing %v is dreaming about end of working day and going back home\n", z.Name)
}

func (z Zookeeper) Introduction() {
	fmt.Printf("Hi, my name is %v %v and I am a zookeeper! :)\nI'm %v old and I my computer skills are: %v\n", z.Name, z.Surname, z.Age, z.ComputerSkills)
}

func (z Zookeeper) IsCatchingAnimal(a AnimalInterface) AnimalInterface {
	fmt.Printf("\n%v is going to catching animals!\nBut %v years is significant blocker of doing this\n", z.Name, z.Age)
	fmt.Printf("So %v is going to catch %v and he got it!\n", z.Name, a.GetName())
	return a
}

func (z Zookeeper) PutAnimalToCage(a AnimalInterface, c *Cage) AnimalInterface {
	fmt.Printf("%v was caught and placed to cage!\n", a.GetName())
	c.Animals = append(c.Animals, a)
	return a
}

func (z Zookeeper) CheckWhatInCage(c *Cage) {
	fmt.Printf("\nZookeeper %v is checking what in cage:\n", z.Name)
	for _, animal := range c.Animals {
		fmt.Printf("\n%s is the cage now!\n", animal.GetName())
	}
}

func (z Zookeeper) EscapingFromAnimals() {
	fmt.Printf("%v is escaping from animals! There is chaos in city and finally animals are free.\nEND\n", z.Name)
}
