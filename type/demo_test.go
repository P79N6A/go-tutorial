package _type

import (
	"testing"
	"strconv"
	"fmt"
)

// 自定义类型
// go中没有class，如果要定义一个包含数据和方法的类，需要先定义一个struct来包含数据，然后通过方法绑定给sruct添加方法
type User struct {
	Name string
	Age int
}

/////////////////////////////////
// 简单的SayHello方法
/////////////////////////////////
func (u User) SayHello() string{
	return "Hello, I'm " + u.Name + ", " + strconv.Itoa(u.Age) + " years old"
}

func Test_say_hello(t *testing.T)  {
	var tom = User{
		Name:"tom",
		Age:20,
	}

	println(tom.SayHello())      // Hello, I'm tom, 20 years old
}

/////////////////////////////////
// ChangeUserName
////////////////////////////////
func (u User) ChangeUserName(name string){
	u.Name = name
}

func Test_change_username(t *testing.T)  {
	var tom = User{
		Name:"tom",
		Age:20,
	}

	tom.ChangeUserName("jerry")

	println(tom.Name)        // jerry?
}

func (u *User) ChangeUserName2(name string){
	u.Name = name
}

func Test_change_username_2(t *testing.T)  {
	var tom = User{
		Name:"tom",
		Age:20,
	}

	tom.ChangeUserName2("jerry")

	println(tom.Name)        // jerry?
}

//////////////////////////////////////
// ChangeUserAge
/////////////////////////////////////
func (u User) ChangeUserAge(age int){
	u.Age = age
}

func (u *User) ChangeUserAge2(age int){
	u.Age = age
}

func Test_change_age_1(t *testing.T)  {
	var tom = User{
		Name:"tom",
		Age:20,
	}

	tom.ChangeUserAge(30)

	println(tom.Age)        // tom
}

func Test_change_age_2(t *testing.T)  {
	var tom = User{
		Name:"tom",
		Age:20,
	}

	tom.ChangeUserAge2(30)

	println(tom.Age)        // tom
}


///////////////////////////////////////
// ChangeUserAge 修改入参age
//////////////////////////////////////

func (u *User) ChangeUserAge3(age int){
	age ++
	u.Age = age
}

func (u *User) ChangeUserage4(age *int){
	*age ++
	u.Age = *age
}

func Test_change_age(t *testing.T)  {
	var tom = User{
		Name:"tom",
		Age:20,
	}

	var age int = 15
	tom.ChangeUserAge3(age)

	println(age)
	println(tom.Age)        // tom
}

func Test_change_age2(t *testing.T)  {
	var tom = User{
		Name:"tom",
		Age:20,
	}

	var age int = 15

	tom.ChangeUserage4(&age)

	println(age)
	println(tom.Age)        // tom
}


////////////////////////////////
// array
////////////////////////////////
func updateArray(a [5]int) {
	a[0] = 15
}

func updateArray2(a *[5]int) {
	(*a)[0] = 15
}

func Test_change_array(t *testing.T)  {

	var array = [5]int {10,20,30,40,50}

	updateArray(array)
	fmt.Println(array)
}

func Test_change_array2(t *testing.T)  {

	var array = [5]int {10,20,30,40,50}

	updateArray2(&array)
	fmt.Println(array)
}

//////////////////////////////////////////
// Slice
/////////////////////////////////////////
func updateSlice(s []int)  {
	s[0] = 15
}

func Test_update_slice(t *testing.T)  {

	var slice = []int {10,20,30,40,50}

	updateSlice(slice)
	fmt.Println(slice)
}

/////////////////////////////////////////
// Map
/////////////////////////////////////////
func updateMap(m map[string]int){
	m["tom"] = 15
}

func Test_change_map(t *testing.T)  {
	m := map[string]int {
		"tom":10,
		"jerry":20,
	}

	updateMap(m)
	fmt.Println(m)
}


/////////////////////////////////////////
// Slice Append
/////////////////////////////////////////
func Test_test(t *testing.T)  {
	s := []int{1,2,3,4,5,6,7,8,9}

	s1 := s[3:]
	fmt.Println(s1)         // 4 5 6 7 8 9
	fmt.Println(len(s1))    // 6
	fmt.Println(cap(s1))    // 6

	s = s[3:7]
	fmt.Println(s)         // 4 5 6 7
	fmt.Println(len(s))    // 4
	fmt.Println(cap(s))    // 6

	s = append(s,100)

	fmt.Println(s)        // 4 5 6 7 100
	fmt.Println(len(s))   // 5
	fmt.Println(cap(s))   // 6

	fmt.Println(s1)          // 4 5 6 7 100 9
	fmt.Println(len(s1))     // 6
	fmt.Println(cap(s1))     // 6
}

func Test_test_3(t *testing.T)  {
	s := []int{5, 7, 9}
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s,x,y)
}

func Test_test_2(t *testing.T)  {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
}
