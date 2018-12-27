package net

import "fmt"

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER      = '\t'
)

var logSn = 0

func printLog(format string, args ...interface{}) {
	fmt.Printf("%d: %s", logSn, fmt.Sprintf(format, args...))

}
