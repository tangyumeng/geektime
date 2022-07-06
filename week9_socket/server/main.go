package main

import (
	"bufio"
	"fmt"
	"net"
	"socket/protocol"
)

// 服务端 fix length
func fixedLength(conn net.Conn) {
	fmt.Println("server, fixed length")

	const LENGTH = 1024

	for {
		var buf = make([]byte, LENGTH)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("client data :", string(buf))
	}
}

// 服务端 delimiter based
func delimiter(conn net.Conn) {
	fmt.Println("server, delimiter based")

	reader := bufio.NewReader(conn)
	for {
		slice, err := reader.ReadSlice('\n')
		if err != nil {
			continue
		}
		fmt.Printf("%s", slice)
	}
}

// 服务端 frameDecoder
func frameDecoder(conn net.Conn) {
	fmt.Println("server, frameDecoder")
	var buf = make([]byte, 0)
	var readerChannel = make(chan []byte, 16)
	go func() {
		select {
		case data := <-readerChannel:
			fmt.Println("channel=", string(data))
		}
	}()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		protocol.Unpack(append(buf, buffer[:n]...), readerChannel)
	}
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:9000")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		defer conn.Close()

		if err != nil {
			fmt.Println(err)
			return
		}
		go frameDecoder(conn)
	}

}
