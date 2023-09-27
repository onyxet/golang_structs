package vehicle

type Vehicle interface {
	Move(r *Route)
	Stop()
	ChangeSpeed()
	Board(p *Passenger)
	TakeOff(p *Passenger)
	GetName() string
}
