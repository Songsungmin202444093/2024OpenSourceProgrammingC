package main

import "fmt"

func main() {
	var student1 struct {
		id   int
		name string
		gpa  float32
	}
	student1.id = 20241234
	student1.name = "송성민"
	student1.gpa = 4.5
	fmt.Println(student1.gpa)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 20244321
	student2.name = "아이다"
	student2.gpa = 4.43
	fmt.Println(student2.id)
}