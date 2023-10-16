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
}

var (
	words     = flag.Int("words", 2, "The number of words in the pet name")
	separator = flag.String("separator", "-", "The separator between words in the pet name")
	FakeTasks = GenerateFakeTasks()
)

func GenerateFakeTasks() []Task {
	rand.Seed(time.Now().UnixNano())
	var tasks []Task
	for i := 0; i < 10; i++ {
		tasks = append(tasks, Task{ID: rand.Intn(100), Name: petname.Generate(*words, *separator)})
	}
	return tasks
}
