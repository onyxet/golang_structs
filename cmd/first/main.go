package main

import (
	"context"
	"flag"
	"structs/pkg/first"
	"sync"
	"time"
)

func main() {
	items := make(chan first.Item)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	buyerName := flag.String("n", "fakeRandomName", "buyer name")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	go first.GenerateFakeRandomOrder(&wg, ctx, *buyerName, items)
	wg.Add(1)
	go first.CreateFakeOrder(&wg, ctx, items)
	wg.Wait()
}
