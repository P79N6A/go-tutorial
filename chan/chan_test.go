package _chan

import (
	"strconv"
	"fmt"
	"testing"
)

func makeCakeAndSend(count int) chan string{
	cs := make(chan string, count)
	for i := 0; i < count ; i++  {
		cakeName := "Cake " + strconv.Itoa(i)
		cs <- cakeName
	}

	close(cs)

	return cs
}

func receiveCakeAndPack(cs chan string) {
	for s := range cs{
		fmt.Println("Packing received cake: ", s)
	}
}

func Test_chan(t *testing.T){
	cs := makeCakeAndSend(8)
	receiveCakeAndPack(cs)
}


