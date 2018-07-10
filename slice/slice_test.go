package slice

import (
	"testing"
	"fmt"
)

func Test_update_map(t *testing.T)  {
	m1 := make([]string, 1)
	m1[0] = "old"
	fmt.Println("调用 func1 前 m1 值:", m1)
	func1(m1)
	fmt.Println("调用 func1 后 m1 值:", m1)
}

func Test_update_map2(t *testing.T)  {
	m1 := make([]string, 10)
	m1[0] = "test"
	fmt.Println("调用 func1 前 m1 值:", m1)
	func2(m1)
	fmt.Println("调用 func1 后 m1 值:", m1)
}

func func1 (a []string) {
	a[0] = "new"
	fmt.Println("func1中:", a)
}

func func2 (a []string) {
	 a = append(a, "val1")
	//a = append(a, "val1")
	fmt.Println("func1中:", a, cap(a))
}

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
