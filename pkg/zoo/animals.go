package zoo

import "fmt"

type AnimalInterface interface {
	Run()
	GetName() string
}

type AnumalsInSluce struct {
	Animals []Animal
}

type Animal struct {
	Name    string
	Species string
	Age     int
}

type Puma struct {
	Animal
	HasFangs bool
	Color    string
	Weight   int
}

type Elephant struct {
	Animal
	TrunkLength int
	Color       string
	Weight      int
	LoveBanas   bool
}

type Chimpanzee struct {
	Animal
	HasTail bool
	Color   string
	Weight  int
	IsFunny bool
}

type Panda struct {
	Animal
	HasTail  bool
	Color    string
	Weight   int
	IsLovely bool
}

type Leopard struct {
	Animal
	HasTail         bool
	HasFangs        bool
	Color           string
	Weight          int
	IsNotFunnyAtAll bool
}

type AnimalsInSlice struct {
	Animals []AnimalInterface
}

func (s *AnimalsInSlice) AddAnimal(i AnimalInterface) {
	s.Animals = append(s.Animals, i)
}

func (p Puma) Run() {
	fmt.Printf("%v %v over %vkgs is escaping from the zoo holding its tail high\n", p.Color, p.Name, p.Weight)
}

func (e Elephant) Run() {
	fmt.Printf("%v %v is escaping from the zoo smashing everything on its way having over %vkgs! OH MY GOSH\n", e.Color, e.Name, e.Weight)
}

func (c Chimpanzee) Run() {
	fmt.Printf("%v %v is escaping from the zoo by trees!\n", c.Color, c.Name)
}

func (p Panda) Run() {
	fmt.Printf("%v %v was thinking about escaping from the zoo, but it is too lazy and still slipping\n", p.Color, p.Name)
}

func (p Leopard) Run() {
	fmt.Printf("%v hunting for a prey! And it's quite dangerous!\n", p.Name)
}

func (a Animal) Run() {
	fmt.Printf("%v hunting for a prey! And it's quite dangerous!\n", a.Name)
}

func (a Animal) GetName() string {
	return a.Name
}

func (p Puma) GetColor() string {
	return p.Color
}

func (c Chimpanzee) GetColor() string {
	return c.Color
}

func (e Elephant) GetColor() string {
	return e.Color
}

func (l Leopard) GetColor() string {
	return l.Color
}

func (p Panda) GetColor() string {
	return p.Color
}
