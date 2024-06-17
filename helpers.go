package main

import (
	"fmt"
	"os"
	"strings"
	// "sync"
	"syscall"
	"unsafe"
	// "github.com/creack/pty"
)

// Winsize represents the terminal size
type Winsize struct {
	Rows uint16
	Cols uint16
	X    uint16
	Y    uint16
}

// getWinsize gets the size of the terminal
func getWinsize() (*Winsize, error) {
	ws := &Winsize{}
	_, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
	if err != 0 {
		return nil, err
	}
	return ws, nil
}

// displayTextInCenter displays the given text in the center of the terminal
func DisplayTextInCenter(text string) error {
	ws, err := getWinsize()
	if err != nil {
		return err
	}

	lines := strings.Split(text, "\n")
	maxLineLength := 0
	for _, line := range lines {
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}

	// Calculate the starting position
	startRow := int(ws.Rows)/2 - len(lines)/2
	startCol := int(ws.Cols)/2 - maxLineLength/2

	// Print blank lines to move to the starting row
	for i := 0; i < startRow; i++ {
		fmt.Println()
	}

	// Print each line of the text at the center
	for _, line := range lines {
		fmt.Printf("%*s%s\n", startCol, "", line)
	}

	return nil
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func CalculateAverageScore(students *[]Student) {
	for i := range *students {
		totalScore := 0.0

		for _, grade := range (*students)[i].grade {
			totalScore += grade.score
		}

		(*students)[i].avgScore = totalScore / float64(3)
	}
}

func DetermineOverallPerformance(student Student, studentPerformanceCh chan string) {

	if student.avgScore >= 70 {
		student.overAllPerformance = "A"
	} else if student.avgScore >= 60 {
		student.overAllPerformance = "B"
	} else if student.avgScore >= 50 {
		student.overAllPerformance = "C"
	} else if student.avgScore >= 40 {
		student.overAllPerformance = "D"
	} else {
		student.overAllPerformance = "F"
	}

	studentPerformanceCh <- student.overAllPerformance
}

// without using goroutines
func DetermineOverallPerformanceW(student Student) string {
	totalScore := 0.0

	for _, grade := range student.grade {
		totalScore += grade.score
	}

	avgScore := totalScore / float64(len(student.grade))

	if avgScore >= 70 {
		student.overAllPerformance = "A"
	} else if avgScore >= 60 {
		student.overAllPerformance = "B"
	} else if avgScore >= 50 {
		student.overAllPerformance = "C"
	} else if avgScore >= 40 {
		student.overAllPerformance = "D"
	} else {
		student.overAllPerformance = "F"
	}
	return student.overAllPerformance
}

// DetermineStatus determines the status of the student wihtout using goroutines
func DetermineStatusW(student Student) string {
	if student.overAllPerformance == "F" {
		student.status = "Fail"
	} else {
		student.status = "Pass"
	}
	return student.status
}

// with goroutines
func DetermineStatus(student Student, studentStatusCh chan string) {

	if student.overAllPerformance == "F" {
		student.status = "Fail"
	} else {
		student.status = "Pass"
	}

	studentStatusCh <- student.status
}

// display students details as a table
func DisplayStudents(students []Student) {
	fmt.Print("\n----------------------------------------------------------------------------------------------------------------------------\n")
	fmt.Printf("%5s\t%-5s %-10s\t %-10s\t%-10s\t%-10s\t%-10s\t%-10s\t%-10s\n", "Rank", "ID", "Name", "Math", "Eng", "Bio", "Avg", "Overall", "Status")
	fmt.Print("----------------------------------------------------------------------------------------------------------------------------\n")

	for i, student := range students {
		fmt.Printf("%5d\t%-5s %-10s\t", i+1, student.ID, student.name)
		for _, grade := range student.grade {
			fmt.Printf("%-10.2f\t", grade.score)
		}
		fmt.Printf("%-10.2f\t", student.avgScore)
		fmt.Printf("%-10s\t", student.overAllPerformance)
		fmt.Printf("%-10s\n", student.status)
	}
}

// filter out passing students
func FilterPassingStudents(students []Student) ([]Student, []Student) {
	var passingStudents []Student
	var failingStudents []Student
	for _, student := range students {
		if student.status == "Pass" {
			passingStudents = append(passingStudents, student)
		} else {
			failingStudents = append(failingStudents, student)
		}
	}
	return passingStudents, failingStudents
}

// sort the passing students by their avgScore decending
func SortPassingStudents(passingStudents []Student) {
	for i := 0; i < len(passingStudents); i++ {
		for j := i + 1; j < len(passingStudents); j++ {
			if passingStudents[i].avgScore < passingStudents[j].avgScore {
				passingStudents[i], passingStudents[j] = passingStudents[j], passingStudents[i]
			}
		}
	}
}
