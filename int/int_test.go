package int

import (
	"testing"
	"fmt"
)

func update(i int)  {
	i++
}

func update2(i *int)  {
	*i++
}

func Test_int(t *testing.T)  {
	var i int = 10
	update(i)
	fmt.Println(i)

	update2(&i)
	fmt.Println(i)
}
