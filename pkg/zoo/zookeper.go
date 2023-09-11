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

func (z Zookeeper) IsCatchingAnimals() {
	fmt.Printf("\n%v is going to catching animals!\nBut %v years is significant blocker of doing this\n", z.Name, z.Age)
}

func (z Zookeeper) EscapingFromAnimals() {
	fmt.Printf("%v is escaping from animals! There is chaos in city and finally animals are free.\nEND\n", z.Name)
}
