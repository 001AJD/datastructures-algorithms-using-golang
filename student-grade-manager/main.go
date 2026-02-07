package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name     string     `json:"studentName"`
	Subjects [3]float32 `json:"subjects"`
	Average  float32    `json:"average"`
}

func main() {
	var str1 string = "This is golang programming practice!"
	report := CountWord(str1)
	fmt.Println(report)
	fmt.Println("Student Grade management system")
	var classRoom []Student = []Student{}
	var student1 Student = Student{
		Name:     "Ajinkya",
		Subjects: [3]float32{75.5, 85.5, 96.3},
	}
	var student2 Student = Student{
		Name:     "Pratik",
		Subjects: [3]float32{75.5, 89.9, 63.6},
	}
	var student3 Student = Student{
		Name:     "Ravi",
		Subjects: [3]float32{95.00, 96.00, 98.00},
	}
	var student4 Student = Student{
		Name:     "Madhav",
		Subjects: [3]float32{92.23, 99.66, 89.30},
	}
	var student5 = Student{
		Name:     "Shahid",
		Subjects: [3]float32{95.66, 65.66, 98.99},
	}
	classRoom = append(classRoom, student1)
	classRoom = append(classRoom, student2)
	classRoom = append(classRoom, student3)
	classRoom = append(classRoom, student4)
	classRoom = append(classRoom, student5)

	fmt.Println("List of Students in the class room")
	for index, value := range classRoom {
		fmt.Printf("%d) %s %0.2f\n", index+1, value.Name, value.Subjects)
	}

	var toppers []Student = []Student{}
	for _, value := range classRoom {
		average := ((value.Subjects[0] + value.Subjects[1] + value.Subjects[2]) / 3)
		value.Average = average
		if average >= 90 {
			toppers = append(toppers, value)
		}
	}
	fmt.Println("Below are the students whose average score more than 90")
	for index, value := range toppers {
		fmt.Printf("%d) %s %0.2f\n", index+1, value.Name, value.Subjects)
	}

	data, err := json.MarshalIndent(toppers, "", "  ")
	if err != nil {
		fmt.Println("Error in json encoding")
	}
	fmt.Println(string(data))
	fileName := "toppers.json"
	error := os.WriteFile(fileName, data, 0644)
	if error != nil {
		fmt.Println("Error writing to file")
	}
	fmt.Println("File created")
}
