package channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_test_1(t *testing.T) {
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Hello, World")
	}()

	fmt.Println("Exist...")
}

// 无缓冲channel
func Test_channel_no_buffer_test(t *testing.T) {
	var ch = make(chan string)
	go func(message string) {

		time.Sleep(time.Second * 3)
		fmt.Println("Hello, Go!")
		ch <- message

	}("Ping!")

	<-ch

	fmt.Println("Exist...")
}

// 有缓冲channel
func Test_test_8(t *testing.T) {
	queue := make(chan int, 3)

	go func(ch chan<- int) { // chan<- 只能写的channel
		time.Sleep(time.Second * 2)
		fmt.Println("Player 1 finished")
		ch <- 0
	}(queue)

	go func(ch chan<- int) {
		time.Sleep(time.Second * 3)
		fmt.Println("Player 2 finished")
		ch <- 0
	}(queue)

	go func(ch chan<- int) {
		time.Sleep(time.Second * 5)
		fmt.Println("Player 3 finished")
		ch <- 0
	}(queue)

	for i := 0; i < 3; i++ {
		<-queue
	}

	fmt.Println("All players finished")
}

func Test_test_range_chan_1(t *testing.T) {
	var ch = make(chan string)
	go func() {
		// range ch:只要该channel没有close，就一直从channel里取数据。如果没有数据就阻塞。
		// 直到channel被close，才会退出for语句。
		for str := range ch {
			fmt.Println(str)
		}
		fmt.Println("channel is closed, all data is received")
	}()

	ch <- "1"
	time.Sleep(3 * time.Second)
	ch <- "2"
	time.Sleep(3 * time.Second)
	ch <- "3"
	time.Sleep(3 * time.Second)
	ch <- "4"
	time.Sleep(3 * time.Second)
	ch <- "5"
	time.Sleep(3 * time.Second)

	close(ch)

	time.Sleep(10 * time.Second)

}

func Test_test_select_chan_1(t *testing.T) {
	var ch = make(chan string)
	go func() {
		for {
			select {
			case s := <-ch:
				fmt.Println("Get[ ", s, " ] from chan")
			default:
			}
			time.Sleep(1 * time.Second)
		}
	}()

	ch <- "1"
	time.Sleep(2 * time.Second)

	ch <- "2"
	time.Sleep(2 * time.Second)

	ch <- "3"
	time.Sleep(2 * time.Second)

	close(ch)

	time.Sleep(10 * time.Second)

}

func Test_test_chan_capacity_length_1(t *testing.T) {
	chan1 := make(chan string)    // 无缓冲的
	chan2 := make(chan string, 1) // 有一个缓冲
	chan3 := make(chan string, 2) // 有两个缓冲

	fmt.Println(cap(chan1)) // 0
	fmt.Println(cap(chan2)) // 1
	fmt.Println(cap(chan3)) // 2
}
