package concurrent

import (
	"testing"
	"sync"
	"fmt"
	"runtime"
)

var counter int
var wg sync.WaitGroup

func Test_test_7(t *testing.T) {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Result:",counter)
}

func incCounter(id int){
	defer wg.Done()

	for count := 0; count<2 ;count++  {
		value := counter

		// goroutine从当前线程退出，给其他goroutine运行的机会
		runtime.Gosched()

		value ++

		counter = value
	}
}