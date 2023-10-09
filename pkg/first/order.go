package first

import (
	"context"
	"fmt"
	"sync"
)

func GenerateFakeRandomOrder(wg *sync.WaitGroup, ctx context.Context, buyer string, items chan Item) {
	defer wg.Done()
	fakeRandomProducts := Items{
		Item{
			Title: "Milk",
			Cost:  1.99,
			Buyer: buyer,
		},
		Item{
			Title: "Bread",
			Cost:  0.99,
			Buyer: buyer,
		},
		Item{
			Title: "Eggs",
			Cost:  2.99,
			Buyer: buyer,
		},
		Item{
			Title: "Cheese",
			Cost:  4.99,
			Buyer: buyer,
		},
	}
	for _, item := range fakeRandomProducts {
		select {
		case <-ctx.Done():
			return
		case items <- item:
		}
	}
	close(items)
}

func CreateFakeOrder(wg *sync.WaitGroup, ctx context.Context, items chan Item) {
	defer wg.Done()
	var order Item
	for {
		select {
		case <-ctx.Done():
			return
		case item, ok := <-items:
			if !ok {
				fmt.Printf("Order for %s: %s, total cost: %f\n", order.Buyer, order.Title, order.Cost)
				return
			}
			order.Title += item.Title + ", "
			order.Cost += item.Cost
			order.Buyer = item.Buyer
		}
	}
}
