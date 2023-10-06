package first

type Customer struct {
	Name   string  `json:"name"`
	Orders []Order `json:"orders"`
}

type Customers []Customer

func (c *Customers) AddCustomer(customer Customer) {
	*c = append(*c, customer)
}
