package array

import (
	"testing"
	"fmt"
)

func Test_create_array(t *testing.T)  {

	array := [5]string{"tom","jerry","leo","shuke","beita"}
	fmt.Println(array)
}

func updateArray(a [5]string)  {
	a[0] = "hello"
}

func updateArray2(a *[5]string)  {
	a[0] = "hello"
}

func Test_int(t *testing.T)  {
	array := createArray()

	updateArray(array)
	printArray(array)

	updateArray2(&array)
	printArray(array)
}

func createArray() [5]string {
	array := [5]string{"tom","jerry","leo","shuke","beita"}
	return array
}

func printArray(array [5]string)  {
	fmt.Println(array)
}