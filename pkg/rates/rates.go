package rates

import "fmt"

type Subject struct {
	SubjectName string
	StudentName string
	Rates       []float64
}

func AVGRateForSubject(subject Subject) {
	result := 0.0
	for _, v := range subject.Rates {
		result += v
	}
	result = result / float64(len(subject.Rates))
	fmt.Printf("Student: %v\nSubject: %v\nAverage Rate: %.2f\n", subject.StudentName, subject.SubjectName, result)
}
