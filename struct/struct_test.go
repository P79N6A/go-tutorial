package _struct

import (
	"testing"
	"fmt"
)

// 定义一个struct，用于测试值传递和指针传递
type Student struct {
	Name string
	Age int
}

func (s Student) M1() {
	s.Name = "name1"
}

func (s *Student) M2() {
	s.Name = "name2"
}

func updateStudent(s Student) {
	s.Name = "name1"
}

func updateStudent2(s *Student) {
	s.Name = "name2"
}

func printStudent(s Student) {
	fmt.Println(s.Name)
	fmt.Println(s.Age)
}

func Test_update_student(t *testing.T)  {
	s := Student{"leo",10}
	fmt.Println(s.Name)     // 输出leo

	s.M1()
	fmt.Println(s.Name)     // 输出leo, 而不是name1

	s.M2()
	fmt.Println(s.Name)    // 输出name2
}

func Test_update_student2(t *testing.T)  {
	s := Student{"leo",10}
	fmt.Println(s.Name)     // 输出leo

	updateStudent(s)
	fmt.Println(s.Name)     // 输出leo, 而不是name1

	updateStudent2(&s)
	fmt.Println(s.Name)    // 输出name2
}

func Test_print_student(t *testing.T)  {
	s := Student{"leo",10}
	printStudent(s)
}