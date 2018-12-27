package net

import (
	"bytes"
	"fmt"
	"math"
	"net"
	"strconv"
	"time"
)

func serverGo() {
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printLog("Listen Error: %s\n", err)
		return
	}
	defer listener.Close()
	printLog("Go listener for the server.(local address: %s)\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			printLog("Accept Error: %s\n", err)
		}
		printLog("Established a connection with a client application.(remote address: %s)\n", conn.RemoteAddr())

	}

}

func handleConn(conn net.Conn) {
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {

		}
		printLog("Received request: %s (Server)\n", strReq)

		i32Req, err := strconv.Atoi(strReq)
		f64Resp := math.Cbrt(float64(i32Req))
		respMsg := fmt.Sprintf("The cube root of %d is %f.", i32Req, f64Resp)
		n, err := write(conn, respMsg)
		printLog("Sent response (written %d bytes): %s (Server)\n", n, respMsg)
	}
}

func read(conn net.Conn) (string, error) {
	// 将readBytes长度设置为1，防止从conn中读出多余的数据，而对后续的读取操作造成影响
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}
