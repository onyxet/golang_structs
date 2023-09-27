package vehicle

import "fmt"

type Route struct {
	From     string
	To       string
	Vehicles []Vehicle
	Hops     []string
}

func (r *Route) GetAllVehicles() {
	fmt.Printf("On the route from %s to %s there are next vehicles\n", r.From, r.To)
	for _, v := range r.Vehicles {
		fmt.Printf("%s\n", v.GetName())
	}
}

func (r *Route) AddVehicleToRoute(v Vehicle) {
	r.Vehicles = append(r.Vehicles, v)
}
