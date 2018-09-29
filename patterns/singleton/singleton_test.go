package singleton

import (
	"testing"
	"fmt"
)

func Test_Ticker(t *testing.T){

	s := New()

	s["hello"] = "world"

	s2 := New()

	fmt.Println(s2["hello"])
}


