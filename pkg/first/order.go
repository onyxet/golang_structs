package first

// GenerateRandomOrders generates random orders and sends them to the orders channel
func GenerateRandomOrders(orders chan Items) {
	items := Items{
		{"iPhone 14", 1000},
		{"MacBook M1", 1200},
		{"Thinkpad", 1400},
		{"Dell Vostro", 800},
		{"Samsung S10", 900},
	}
	orders <- items
}

func ProcessItems(orders chan Items) Items {
	ordersRes := Items{}
	for {
		select {
		case items := <-orders:
			for _, item := range items {
				ordersRes = append(ordersRes, item)
			}
		}

	}
}
