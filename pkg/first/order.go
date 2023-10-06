package first

import "context"

func GenerateRandomOrder(ctx context.Context, items <-chan Item, order chan<- Item) {
		select {
		case <-ctx.Done():
			close(order)
			return
		case <-items:


		}
	}

}
