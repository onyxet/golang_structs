package main

import "structs/pkg/zoo"

/*
0. Declaration
1. Different types
3. Creating instances, instantiation
4. Zero values
5. Consecutive fields with same type
6. User defined type. Structs are the only way to create concrete user-defined types in Golang.
7. Struct Visibility
8. Fields Visibility
9. Embedding
10. Recursive (almost)
11. Anonymous/unnamed structs
12. Structural comparison
13. methods or pointer and value receivers
14. Tagging
15. Implicit dereferencing
16. Assignment
*/

func main() {
	// Struct declaration
	somePanda := zoo.Panda{
		Animal: zoo.Animal{
			Name:    "Joe",
			Species: "Bear",
			Age:     10,
		},
		HasTail:  true,
		Color:    "Black and White",
		Weight:   100,
		IsLovely: true,
	}
	somePuma := zoo.Puma{
		Animal: zoo.Animal{
			Name:    "Liza",
			Species: "Cat",
			Age:     5,
		},
		HasFangs: true,
		Color:    "Black",
		Weight:   50,
	}
	someElephant := zoo.Elephant{
		Animal: zoo.Animal{
			Name:    "Dumbo",
			Species: "Elephant",
			Age:     15,
		},
		TrunkLength: 100,
		Color:       "Grey",
		Weight:      500,
		LoveBanas:   true,
	}
	someChimpanzee := zoo.Chimpanzee{
		Animal: zoo.Animal{
			Name:    "Bobo",
			Species: "Chimpanzee",
			Age:     7,
		},
		HasTail: true,
		Color:   "Brown",
		Weight:  70,
		IsFunny: true,
	}
	someLeopard := zoo.Leopard{
		Animal: zoo.Animal{
			Name:    "Leo",
			Species: "Cat",
			Age:     12,
		},
		HasTail:         true,
		HasFangs:        true,
		Color:           "Yellow",
		Weight:          80,
		IsNotFunnyAtAll: true,
	}
	someZookeeper := zoo.Zookeeper{
		Human: zoo.Human{
			Name:    "John",
			Surname: "Smith",
			Age:     50,
		},
		BachelorDegree: true,
		ComputerSkills: true,
	}

	someCage := zoo.Cage{
		Enclosure: zoo.Enclosure{
			Size:     100.2,
			Material: "Steel",
		},
		Purpose: "Predators",
	}
	windOfChange := zoo.MagicWind{}

	someZookeeper.Introduction()
	someZookeeper.OnDuty()
	someZookeeper.OffDuty()
	someCage.Close()
	windOfChange.Magic()
	someCage.Open()
	somePanda.Run()
	somePuma.Run()
	someElephant.Run()
	someChimpanzee.Run()
	someLeopard.Run()

	someZookeeper.IsCatchingAnimals()
	someZookeeper.EscapingFromAnimals()
}
