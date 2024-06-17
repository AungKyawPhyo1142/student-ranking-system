package main

import (
	"fmt"
	"math/rand"
)

func randomScore() float64 {
	return 50 + rand.Float64()*50
}

func generateRandomGrades() []Grade {
	subjects := []string{"Math", "English", "Science"}
	grades := make([]Grade, len(subjects))
	for i, subject := range subjects {
		grades[i] = Grade{subject: subject, score: randomScore()}
	}
	return grades
}

func GenerateStudent(num int) []Student {

	students := make([]Student, num)

	for i := 0; i < num; i++ {
		students[i] = Student{
			ID:    fmt.Sprintf("S%02d", i+1),
			name:  fmt.Sprintf("Student %02d", i+1),
			grade: generateRandomGrades(),
		}
	}

	return students
}
