package first

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomNums(nums chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums <- rand.Intn(100)
	}
	close(nums)
}

func FindAvg(nums chan int, avg chan float32) {
	sum := 0
	count := 0
	for num := range nums {
		sum += num
		count++
	}
	avg <- float32(sum) / float32(count)
}

func PrintAvg(avg chan float32) {
	fmt.Println(<-avg)
}
