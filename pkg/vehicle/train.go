package vehicle

import "fmt"

type Train struct {
	HopIndex int
}

func (t *Train) Move(r *Route) {
	if len(r.Hops) == 0 {
		fmt.Println("No hops available.")
		return
	}

	if t.HopIndex < len(r.Hops) {
		fmt.Printf("Car has moved to: %s\n", r.Hops[t.HopIndex])
		t.HopIndex++
	} else {
		fmt.Printf("Car has reached the %s", r.To)
	}
}

func (c *Train) Stop() {
	fmt.Printf("Train stopped\n")
}

func (c *Train) ChangeSpeed() {
	fmt.Printf("Train's speed changed\n")
}

func (a *Train) Board(p *Passenger) {
	fmt.Printf("%s boarded on the car\n", p.Name)
}

func (a *Train) TakeOff(p *Passenger) {
	fmt.Printf("%s took off from the train\n", p.Name)
}

func (a *Train) GetName() string {
	return "Train"
}
