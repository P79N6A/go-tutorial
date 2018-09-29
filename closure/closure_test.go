package closure

import (
	"testing"
	"fmt"
)

func Test_Ticker(t *testing.T){
	nextNum := getSequence()

	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())

	nextNum1 := getSequence()
	fmt.Println(nextNum1())
	fmt.Println(nextNum1())


	nextNum3 := getSequence2(10)
	fmt.Println(nextNum3())
	fmt.Println(nextNum3())
	fmt.Println(nextNum3())
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func getSequence2(x int) func() int {
	return func() int {
		x += 1
		return x
	}
}