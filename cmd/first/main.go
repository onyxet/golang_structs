package main

import (
	"structs/pkg/first"
)

func main() {
	nums := make(chan int)
	avg := make(chan float32)
	go first.GenerateRandomNums(nums)
	go first.FindAvg(nums, avg)
	first.PrintAvg(avg)
}
