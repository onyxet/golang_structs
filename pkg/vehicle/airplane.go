package vehicle

import "fmt"

type Airplane struct {
	HopIndex int
}

func (a *Airplane) Move(r *Route) {
	if len(r.Hops) == 0 {
		fmt.Println("No hops available.")
		return
	}

	if a.HopIndex < len(r.Hops) {
		fmt.Printf("Car has moved to: %s\n", r.Hops[a.HopIndex])
		a.HopIndex++
	} else {
		fmt.Printf("Car has reached the %s", r.To)
	}
}

func (a *Airplane) Stop() {
	fmt.Printf("Airplane stopped\n")
}

func (a *Airplane) ChangeSpeed() {
	fmt.Printf("Speed changed on Airplane\n")
}

func (a *Airplane) Board(p *Passenger) {
	fmt.Printf("%s boarded on the plane\n", p.Name)
}

func (a *Airplane) TakeOff(p *Passenger) {
	fmt.Printf("%s took off from the plane\n", p.Name)
}

func (a *Airplane) GetName() string {
	return "Airplane"
}
