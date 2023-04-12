package main

import (
	"fmt"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

var cache = make(map[string]string)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		defer conn.Close()
		go func(c net.Conn) {
			for {
				buf := make([]byte, 64)
				n, err := c.Read(buf)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println(n, processReq(buf))

				response := []byte(processReq(buf))
				_, err = c.Write(response)
				if err != nil {
					fmt.Println(err)
					_, err = c.Write([]byte(fmt.Sprintf("-ERROR: %s\r\n", err)))
					return
				}
			}

		}(conn)
	}
}

func processReq(buf []byte) string {
	req := string(buf)
	reqSlice := strings.Split(req, "\r\n")
	command := strings.ToUpper(reqSlice[2])
	if command == "COMMAND" {
		return "+OK\r\n"
	}
	if command == "PING" {
		return "+PONG\r\n"
	}
	if command == "ECHO" {
		if len(reqSlice) < 6 {
			return "+\r\n"
		}
		return "+" + reqSlice[4] + "\r\n"
	}
	if command == "SET" {
		cache[reqSlice[4]] = reqSlice[6]
		return "+OK\r\n"
	}
	if command == "GET" {
		return "+" + cache[reqSlice[4]] + "\r\n"
	}
	return "-Invalid Command\r\n"
}
