package _struct

import (
	"fmt"
	"testing"
)

// 定义一个struct，用于测试值传递和指针传递
type Student struct {
	Name string
	Age  int
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

func Test_update_student(t *testing.T) {
	s := Student{"leo", 10}
	fmt.Println(s.Name) // 输出leo

	s.M1()
	fmt.Println(s.Name) // 输出leo, 而不是name1

	s.M2()
	fmt.Println(s.Name) // 输出name2
}

func Test_update_student2(t *testing.T) {
	s := Student{"leo", 10}
	fmt.Println(s.Name) // 输出leo

	updateStudent(s)
	fmt.Println(s.Name) // 输出leo, 而不是name1

	updateStudent2(&s)
	fmt.Println(s.Name) // 输出name2
}

func Test_print_student(t *testing.T) {
	s := Student{"leo", 10}
	printStudent(s)
}

func TestStudentAssignment(t *testing.T) {
	s1 := Student{Name: "tom", Age: 20}
	// 将s1赋值给s2，会开辟一块新的内存，将s1的值拷贝给s2。内存中有两份数据，这两份数据的值一样。s1和s2指向不同的内存块。
	s2 := s1
	s2.Name = "jerry"

	fmt.Println(s1.Name) // tom
	fmt.Println(s2.Name) // jerry
}
