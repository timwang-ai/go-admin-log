package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", ":8889")
	log.Println("start server....")
	if err != nil {
		fmt.Println("listen failed! msg :", err)
		log.Panic(err)
		return
	}
	for {
		client, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	for {
		buf := make([]byte, 100)
		n, err := client.Read(buf) //读取客户端数据
		if err != nil {
			fmt.Println(err)
			return

		}
		fmt.Printf("read data size %d msg:%s", n, string(buf[0:n]))
		msg := []byte("hello,world\n")
		_, _ = client.Write(msg) //发送数据
	}
}
