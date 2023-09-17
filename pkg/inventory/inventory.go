package inventory

type Torch struct {
	BatteryPower int
	CanBeUsed    bool
}

type Knife struct {
	CanBeUsed bool
}

type Camera struct {
	BatteryPower int
	CanBeUsed    bool
}

type Tool interface {
	Broken()
}

func (t *Torch) Broken() *Torch {
	return &Torch{
		CanBeUsed: false,
	}
}

func (k *Knife) Broken() *Knife {
	return &Knife{
		CanBeUsed: false,
	}
}

func (c *Camera) Broken() *Camera {
	return &Camera{
		CanBeUsed: false,
	}
}
