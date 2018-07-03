package main

import (
	"log"
	"os"
	"fmt"
	"go-tutorial/goinaction/search"
)

import _"go-tutorial/goinaction/matchers"

func init() {
	log.SetOutput(os.Stdout)
}

func main()  {
	print("Hello world! From goinaction/main.go")
	search.Run("president")

	fmt.Println(os.Getwd())
}

