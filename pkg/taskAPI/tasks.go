package taskAPI

import (
	"flag"
	"github.com/dustinkirkland/golang-petname"
	"math/rand"
	"time"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

var (
	words     = flag.Int("words", 2, "The number of words in the pet name")
	separator = flag.String("separator", "-", "The separator between words in the pet name")
	FakeTasks = GenerateFakeTasks()
)

func GenerateFakeTasks() []Task {
	rand.Seed(time.Now().UnixNano())
	var tasks []Task
	startDate := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)

	duration := endDate.Sub(startDate)

	for i := 0; i < 10; i++ {
		randomDuration := time.Duration(rand.Int63n(int64(duration)))
		randomDate := startDate.Add(randomDuration)
		tasks = append(tasks, Task{ID: rand.Intn(100), Name: petname.Generate(*words, *separator), Date: randomDate.Format("2006-01-02")})
	}
	return tasks
}
