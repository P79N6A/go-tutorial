package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Ticker(t *testing.T) {

	//i, _ := strconv.Atoi("")
	//fmt.Println(i)

	i := strings.Index("outbound||xxx", "outbound")
	fmt.Println(i)

	//ticker := time.NewTicker(time.Second * 5)
	//go func() {
	//	for t := range ticker.C {
	//		fmt.Println("ticked at:", t, time.Now())
	//	}
	//}()
	//
	//time.Sleep(time.Second * 100)
}
