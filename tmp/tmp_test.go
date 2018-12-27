package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_tmp(t *testing.T) {
	//r := 1.0 / 3
	//fmt.Println(r)

	//var m = map[string]string{}
	//
	//m["hello"] = "hello"
	//
	//s := m["world"]
	//
	//fmt.Println(s)
	//
	//i, err := strconv.Atoi("")
	//
	//fmt.Println(i, err)

	l := strings.Split("11.239.163.19:38269", ":")
	fmt.Println(len(l))
	fmt.Println(l[0])

}
