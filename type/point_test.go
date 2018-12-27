package _type

import (
	"fmt"
	"testing"
)

func Test_pointer(t *testing.T) {

	var a int = 20
	var ip *int

	ip = &a

	fmt.Printf("a value address is %x\n", &a)

	fmt.Printf("ip is %x\n", ip)

	fmt.Printf("*ip is %d\n", *ip)

	m := map[string]string{
		"hello": "world",
	}

	fmt.Printf("m is %v\n", m)
	fmt.Printf("&m is %x\n", &m)

}
