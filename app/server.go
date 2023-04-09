package main

import (
	"bufio"
	"fmt"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

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

				fmt.Println(n, string(buf))
				response := []byte("+PONG\r\n")
				_, err = c.Write(response)
				if err != nil {
					fmt.Println(err)
					_, err = c.Write([]byte(fmt.Sprintf("-ERROR: %s\r\n", err)))
					return
				}
			}

		}(conn)

		//go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// create a new scanner to read data from the connection
	scanner := bufio.NewScanner(conn)

	// read the request line
	scanner.Scan()
	requestLine := scanner.Text()
	fmt.Println("Request line:", requestLine)

	// // read the request headers
	// headers := make(map[string]string)
	// for scanner.Scan() {
	// 	headerLine := scanner.Text()
	// 	if headerLine == "" {
	// 		break
	// 	}
	// 	parts := strings.SplitN(headerLine, ":", 2)
	// 	if len(parts) == 2 {
	// 		headers[parts[0]] = parts[1]
	// 	}
	// }

	// // read the request body, if any
	// if contentLength, ok := headers["Content-Length"]; ok {
	// 	bodyLength := contentLength
	// 	fmt.Println("Content-Length:", bodyLength)
	// 	body := make([]byte, int(1024))
	// 	_, err := conn.Read(body)
	// 	if err != nil {
	// 		fmt.Println("Error reading request body:", err)
	// 	} else {
	// 		fmt.Println("Request body:", string(body))
	// 	}
	// }
}
