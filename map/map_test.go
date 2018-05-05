package _map

import (
	"testing"
	"fmt"
)

func Test_create_map(t *testing.T)  {
	dict1 := make(map[string]int)
	dict1["foo"] = 10
	fmt.Println(dict1["foo"])

	dict2 := map[string]int{"tom":10,"jerry":20}
	dict2["leo"] = 30
	fmt.Println(dict2["tom"])
	fmt.Println(dict2["leo"])
}
