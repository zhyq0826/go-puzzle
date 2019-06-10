package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func pase_student() {
	m1 := make(map[string]*Student)
	m2 := make(map[string]*Student)
	stus := []Student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for index, stu := range stus {
		fmt.Println(index)
		m1[stu.Name] = &stu
		m2[stu.Name] = &stus[index]
	}

	fmt.Println(m1)
	fmt.Println(m2)

}

func main() {
	pase_student()
}
