package main

import (
	"context"
	"structs/pkg/first"
)

func main() {
	nums := make(chan int)
	ctx := context.Context()

	//avg := make(chan float32)
	go first.GenerateRandomOrder(nums)
}
