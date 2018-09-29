package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_Ticker(t *testing.T) {

	ticker := time.NewTicker(time.Second * 5)
	go func() {
		for t := range ticker.C {
			fmt.Println("ticked at:", t, time.Now())
		}
	}()

	time.Sleep(time.Second * 100)
}
