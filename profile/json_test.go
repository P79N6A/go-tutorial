package profile

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Json1(t *testing.T) {

	//f, err := os.Create("cpu-profile.prof")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//pprof.StartCPUProfile(f)
	json1()
	//pprof.StopCPUProfile()
}

func Test_Json2(t *testing.T) {

	//f, err := os.Create("cpu-profile2.prof")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//pprof.StartCPUProfile(f)
	json2()
	//pprof.StopCPUProfile()
}

func json1() {
	s := `{"name":"tom","age":10,"courses":[{"name":"english","score":80},{"name":"math","score":90},{"name":"math","score":70}]}`

	for i := 0; i < 10000; i++ {
		student := &Student{}
		json.Unmarshal([]byte(s), student)
		fmt.Println(*student)
	}
}

func json2() {

	s := `{"name":"tom","age":10}`

	for i := 0; i < 10000; i++ {
		student := &Student{}
		json.Unmarshal([]byte(s), student)
		fmt.Println(*student)
	}
}

type Student struct {
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Courses []*Course `json:"courses"`
}

type Course struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
