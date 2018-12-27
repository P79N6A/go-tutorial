package net

import (
	"net"
	"time"
)

func clientGo(id int) {
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()
	printLog("Connected to server.(remote address: %s, local address: %s) (client[%d]) \n", conn.RemoteAddr(), conn.LocalAddr(), id)
	time.Sleep(200 * time.Millisecond)
}
