package main

import "structs/pkg/rates"

func main() {
	sbj1 := rates.Subject{
		SubjectName: "Math",
		StudentName: "John",
		Rates:       []float64{90.5, 80.0, 70.2, 12, 4},
	}
	rates.AVGRateForSubject(sbj1)
}
