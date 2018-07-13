package channel

import (
	"testing"
	"fmt"
	"time"
)

func Test_test_1(t *testing.T)  {
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Hello, World")
	}()

	fmt.Println("Exist...")
}




// 无缓冲channel
func Test_channel_no_buffer_test(t *testing.T)  {
	var ch = make(chan string)
	go func(message string) {

		time.Sleep(time.Second * 3)
		fmt.Println("Hello, Go!")
		ch <- message

	}("Ping!")

	<- ch

	fmt.Println("Exist...")
}




// 有缓冲channel
func Test_test_8(t *testing.T) {
	queue := make(chan int,3)

	go func(ch chan<- int) {             // chan<- 只能写的channel
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

	for i:= 0; i<3 ; i++{
		<- queue
	}

	fmt.Println("All players finished")
}

