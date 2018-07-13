// 互斥锁

package concurrent

import (
	"testing"
	"sync"
	"runtime"
	"fmt"
)

var mutex sync.Mutex

func Test_mutex_test(t *testing.T) {
	wg.Add(2)

	go incCounter2(1)
	go incCounter2(2)

	wg.Wait()
	fmt.Println("Result:",counter)
}

func incCounter2(id int)  {
	defer wg.Done()

	for count:=0; count<2;count++{
		mutex.Lock()            // mutex开始临界区，同一时刻只有一个goroutine可以进入临界区
		{                       // 使用大括号让临界区看起来更清晰，并不是必要的
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}