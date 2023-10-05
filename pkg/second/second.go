package second

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func GenerateRandomNums(size int, nums chan int, res chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		nums <- rand.Intn(size)
	}
	close(nums)
	min := <-res
	max := <-res
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}

func FindMaxAndMin(nums chan int, res chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var resNums []int
	for num := range nums {
		resNums = append(resNums, num)
	}
	sort.Ints(resNums)
	res <- resNums[0]
	res <- resNums[len(resNums)-1]
	close(res)
}
