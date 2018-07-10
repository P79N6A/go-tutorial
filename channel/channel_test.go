package channel

import (
	"testing"
	"fmt"
	"time"
)

func Test_test_1(t *testing.T)  {
	var ch = make(chan string)
	go func(message string) {
		ch <- message

	}("Ping!")

	fmt.Println(<- ch)
}

var ch = make(chan int)    // 无缓冲channel

func loop(){
	for i:=0; i<10 ;i++ {
		fmt.Print(i)
	}
	ch <- 0
}

func Test_test_2(t *testing.T)  {
	go loop()
	<- ch
}

func Test_test_3(t *testing.T)  {
	ch := make(chan int,3)
	ch <- 1
	println("Put 1 into channel")
	ch <- 2
	println("Put 2 into channel")
	ch <- 3
	println("Put 3 into channel")
	//ch <- 4
	//println(4)
	//ch <- 5
	//println(5)

	for v := range ch{
		fmt.Println(v)
		if len(ch) <= 0 {
			break
		}
	}
}

func Test_test_4(t *testing.T)  {
	ch := make(chan int,3)
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	for v := range ch{
		fmt.Println(v)
		if len(ch) <= 0 {
			break
		}
	}
}

func Test_test_5(t *testing.T)  {
	var ch = make(chan int,10)

	for i := 0; i<10 ; i++ {
		go func(id int) {
			println(id)
			ch <- 0
		}(i)
	}

	for i:= 0; i< 10 ; i++ {
		<- ch
	}
	println("Finished")
}




func Test_test_6(t *testing.T)  {
	var ch = make(chan int)

	go func() {
		println("hello")
		ch <- 0       // 发送一个信号，是什么值无所谓
	}()

	// do something here...
	<- ch
}
// 如果是无缓冲channel:
// 1. 发送方会一直阻塞，直到接收方把数据取出。
// 2. 接收方会一直阻塞，直到有数据到来

// 如果是有缓冲channel:
// 1. 如果缓冲区没满，发送方将数据写到缓冲区就退出了，如果缓冲区已满，会阻塞，直到接收方取出数据，缓冲区有空间了，将数据放入缓冲区，然后退出。
// 2. 接收方从缓冲区拿完数据就退出了，如果缓冲区里没数据，就一直等，直到有数据来。

func Produce(queue chan<- int) {
	for i:= 0; i<100 ;i++ {
		queue <- i
	}
}

func Consume(queue <-chan int) {
	for i:= 0;i<100 ;i++ {
		v := <-queue
		fmt.Println("received data: ",v)
	}
}

func Test_test_7(t *testing.T) {
	queue := make(chan int,1)

	go Produce(queue)
	go Consume(queue)

	time.Sleep(time.Second * 5)
}







