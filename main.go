package main

import (
	"fmt"
	// "sync"
	// "time"
)

type Student struct {
	ID                 string
	name               string
	grade              []Grade
	avgScore           float64
	overAllPerformance string
	status             string
}

type Grade struct {
	subject string
	score   float64
}

// student mock data
var mockStudents = GenerateStudent(20)

func main() {
	// ClearScreen()
	// text := "Hello, Welcome to the student ranking system!\nPress any key to continue..."
	// if err := DisplayTextInCenter(text); err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Scanln()

	// ClearScreen()

	// var num int
	// fmt.Println("Enter the number of students: ")
	// fmt.Scanln(&num)

	// // make a slice of students with the given number
	// students := make([]Student, num)

	// // iterate over the students slice
	// for i := 0; i < num; i++ {
	// 	var name string

	// 	id := i + 1
	// 	fmt.Print("\nEnter Student Name: ")
	// 	fmt.Scanln(&name)

	// 	// make a slice of grades with 3 subjects: Math, English, and Science
	// 	subjects := []string{"Math", "English", "Science"}
	// 	grades := make([]Grade, 3)

	// 	// assign the subjects to the grades
	// 	for i, subject := range subjects {
	// 		grades[i].subject = subject
	// 	}

	// 	for j := 0; j < len(grades); j++ {
	// 		var score float64
	// 		fmt.Printf("Enter %s score: ", grades[j].subject)
	// 		fmt.Scanln(&score)
	// 		grades[j].score = score
	// 	}

	// 	// assign the student details to the students slice
	// 	students[i] = Student{
	// 		ID:                 id,
	// 		name:               name,
	// 		grade:              grades,
	// 		overAllPerformance: "", // default
	// 		status:             "", // default
	// 	}

	// 	ClearScreen()

	// }

	// // display the student details
	// for _, student := range students {
	// 	fmt.Printf("ID: %d\nName: %s\n", student.ID, student.name)
	// 	for _, grade := range student.grade {
	// 		fmt.Printf("%s: %.2f\n", grade.subject, grade.score)
	// 	}
	// 	fmt.Println()
	// }

	// calculate avgScore
	CalculateAverageScore(&mockStudents)

	var studentPerformanceCh = make(chan string)
	var studentStatusCh = make(chan string)

	// determine the overall performance of each student
	for i := 0; i < len(mockStudents); i++ {
		go DetermineOverallPerformance(mockStudents[i], studentPerformanceCh)
		go DetermineStatus(mockStudents[i], studentStatusCh)
		// mockStudents[i].overAllPerformance = DetermineOverallPerformanceW(mockStudents[i])
		// mockStudents[i].status = DetermineStatusW(mockStudents[i])

		// assign into student[i].overAllPerformance
		mockStudents[i].overAllPerformance = <-studentPerformanceCh
		mockStudents[i].status = <-studentStatusCh

	}

	// DisplayStudents(mockStudents)

	// // filter the passing students
	passingStudents, _ := FilterPassingStudents(mockStudents)

	ClearScreen()
	fmt.Print("\n----------------------------------------------------------------------------------------------------------------------------\n")
	fmt.Print("Passing Students")

	SortPassingStudents(passingStudents)
	DisplayStudents(passingStudents)

	// fmt.Println("\nFailing Students")
	// DisplayStudents(failingStudents)

}
