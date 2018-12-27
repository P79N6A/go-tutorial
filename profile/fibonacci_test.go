package profile

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"testing"
)

func Test_Fibonacci(t *testing.T) {
	f, err := os.Create("cpu-profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	fmt.Println(fibonacci(45))
	pprof.StopCPUProfile()
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
