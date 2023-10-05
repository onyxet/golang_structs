package main

import (
	"structs/pkg/second"
	"sync"
)

func main() {
	sizeOnRandom := 10001
	nums := make(chan int)
	res := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go second.GenerateRandomNums(sizeOnRandom, nums, res, &wg)
	go second.FindMaxAndMin(nums, res, &wg)
	wg.Wait()
}
