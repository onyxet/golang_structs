package main

import (
	"fmt"
	"structs/pkg/post"
	"structs/pkg/vehicle"
)

func main() {
	// Implementing sort logic
	somePost := post.PostSortDept{DeptName: "Random Post Dept"}
	somePost.Sort(&post.Box{SenderAddr: "123 Main St", RecipientAddr: "456 Side St"})
	somePost.Sort(&post.Envelope{SenderAddr: "123 Main St", RecipientAddr: "456 Side St"})

	// Implementing route logic
	fmt.Println()
	kyivLviv := vehicle.Route{From: "Kyiv", To: "Lviv", Hops: []string{"Zhytomyr", "Rivne", "Lutsk"}}
	kyivDnipro := vehicle.Route{From: "Kyiv", To: "Dnipro", Hops: []string{"Brovary", "Reshetylivka", "Tsarychanka"}}
	psg1 := vehicle.Passenger{Name: "John"}
	psg2 := vehicle.Passenger{Name: "Jane"}
	psg3 := vehicle.Passenger{Name: "Jack"}
	someCar := vehicle.Car{}
	someTrain := vehicle.Train{}
	someAirplane := vehicle.Airplane{}
	kyivLviv.AddVehicleToRoute(&someCar)
	kyivLviv.AddVehicleToRoute(&someTrain)
	kyivDnipro.AddVehicleToRoute(&someAirplane)
	kyivLviv.GetAllVehicles()
	kyivDnipro.GetAllVehicles()
	someCar.Board(&psg1)
	someTrain.Board(&psg2)
	someAirplane.Board(&psg3)
	someCar.Move(&kyivLviv)
	someCar.Move(&kyivLviv)
	someCar.Move(&kyivLviv)
	someCar.Move(&kyivLviv)
	someTrain.Move(&kyivLviv)
	someTrain.Move(&kyivLviv)

}
