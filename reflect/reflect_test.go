package reflect

import (
	"testing"
	"reflect"
	"fmt"
)

func Test_Ticker(t *testing.T){
	s := Student{Name:"tom", Age: 20}

	typ := reflect.TypeOf(s)
	//v := reflect.ValueOf(s)

	//fmt.Println("type: ", typ.Kind())
	//fmt.Println("value: ", v)

	fmt.Println(typ.NumMethod())

	//for m := 0; m < typ.NumMethod(); m++ {
	//	method := typ.Method(m)
	//	mtype := method.Type
	//	//mname := method.Name
	//	// Method must be exported.
	//	if method.PkgPath != "" {
	//		continue
	//	}
	//
	//	fmt.Println(mtype.NumIn())
	//
	//	// Method needs four ins: receiver, context.Context, *args, *reply.
	//	}
	}

type Student struct {
	Name string
	Age int
}

func hello(s1 string, s2 string){
	fmt.Println(s1,s2)
}

func (s *Student) hello(s1 string, s2 string) {
	fmt.Println(s1, s2)
}

