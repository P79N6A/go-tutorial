package _func

import (
	"fmt"
	"testing"
)

func Test_func(t *testing.T) {
	users := []string{"tom", "jerry"}
	sayHello("Welcome", users...)
}

func sayHello(greet string, who ...string) {
	for _, name := range who {
		fmt.Println(greet, name)
	}
}
