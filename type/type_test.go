package _type

import (
	"testing"
	"fmt"
)

// 定义一个struct，用于测试值传递和指针传递
type Student struct {
	Name string
	Age int
}

func (t Student) M1() {
	t.Name = "name1"
}

func (t *Student) M2() {
	t.Name = "name2"
}

func updateName(t Student) {
	t.Name = "name1"
}

func updateName2(t *Student) {
	t.Name = "name2"
}

func Test_update_name(t *testing.T)  {
	t1 := Student{"leo",10}
	fmt.Println(t1.Name)     // 输出leo

	t1.M1()
	fmt.Println(t1.Name)     // 输出leo, 而不是name1

	t1.M2()
	fmt.Println(t1.Name)    // 输出name2
}

func Test_updateName2(t *testing.T)  {
	t1 := Student{"leo",10}
	fmt.Println(t1.Name)     // 输出leo

	updateName(t1)
	fmt.Println(t1.Name)     // 输出leo, 而不是name1

	updateName2(&t1)
	fmt.Println(t1.Name)    // 输出name2
}





func Test_delete_from_map2(t *testing.T)  {

	var a int = 10;
	fmt.Println(a)

	b := a;
	fmt.Println(b)

	b ++

	fmt.Println("a:", a)
	fmt.Println("b:",b)
}


func Test_delete_from_map3(t *testing.T)  {
	var a int = 10
	update(a)
	fmt.Println("a:", a)

	update2(&a)
	fmt.Println("a:", a)

}




func update(i int)  {
	i ++
}

func update2(i *int)  {
	*i++
}