package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_4_Reflect_1(t *testing.T) {
	u := User{"tom", 20, "beijing"}

	v := reflect.ValueOf(u)
	fmt.Println(v)

	typ := reflect.TypeOf(u)
	fmt.Println(typ)

	fmt.Println(reflect.TypeOf(u).Kind()) // struct

	fmt.Println(reflect.TypeOf(&u).Kind()) // ptr

	var tom Hello = &User{"tom", 20, "beijing"}

	fmt.Println(reflect.TypeOf(tom).Kind()) // struct

}

type User struct {
	Name string
	Age  int
	addr string
}

func (u User) Hello() {
	fmt.Printf("Name is %s, Age is %d \n", u.Name, u.Age)
}

type Hello interface {
	Hello()
}
