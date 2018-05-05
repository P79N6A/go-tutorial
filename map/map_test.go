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


	dict3 := map [int] []string {}
	dict3[1] = []string{"hello","world hhh"}

	fmt.Println(dict3)
}

func Test_create_nil_map(t *testing.T)  {
	// 这种方式创建的map是nil，如果直接操作它，会抛异常
	var dict map[int]string

	dict[1] = "hello"
}

func Test_create_nil_map2(t *testing.T)  {
	colors := map[string]string{"red":"100","yellow":"200"}

	value, exists := colors["red"]
	if exists {
		fmt.Println(value)
	}
}

func Test_iterate_map(t *testing.T)  {
	colors := createStringStringMap()

	for key,value := range colors{
		fmt.Printf("key: %s  value: %s\n",key,value)
	}
}

func createStringStringMap() map[string]string {
	colors := map[string]string{
		"Blue" : "#f0f8ff",
		"Coral": "#ff7F50",
		"DarkGray": "#a9a9a9",
		"Green": "#228b22",
	}
	return colors
}


func Test_delete_from_map(t *testing.T)  {
	colors := createStringStringMap()
	delete(colors, "Blue")
	printMap(colors)
}

func printMap(m map[string]string)  {
	for key,value := range m{
		fmt.Printf("key: %s  value: %s\n",key,value)
	}
}

func deleteColor(colors map[string]string, key string)  {
	delete(colors,key)
}



func Test_delete_from_map2(t *testing.T)  {
	// 创建一个map
	colors := createStringStringMap()
    // 打印map
	printMap(colors)
    // 调用方法删除一个颜色
	deleteColor(colors,"Blue")
    // 打印map
	printMap(colors)

	// 可以看到Blue被删掉了。
	// 在函数中传参map时，不会创建一个该map的一个副本
	// 依然使用同一个map
}

