package vehicle

import "fmt"

type Car struct {
	HopIndex int
}

func (c *Car) Move(r *Route) {
	if len(r.Hops) == 0 {
		fmt.Println("No hops available.")
		return
	}

	if c.HopIndex < len(r.Hops) {
		fmt.Printf("Car has moved to: %s\n", r.Hops[c.HopIndex])
		c.HopIndex++
	} else {
		fmt.Printf("Car has reached the %s\n", r.To)
	}
}

func (c *Car) Stop() {
	fmt.Printf("Car stopped\n")
}

func (c *Car) ChangeSpeed() {
	fmt.Printf("Car's speed changed\n")
}

func (a *Car) Board(p *Passenger) {
	fmt.Printf("%s boarded on the car\n", p.Name)
}

func (a *Car) TakeOff(p *Passenger) {
	fmt.Printf("%s took off from the car\n", p.Name)
}

func (a *Car) GetName() string {
	return "Car"
}
