package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go func(conn net.Conn) {
			buff := make([]byte, 1024)
			_, err = conn.Read(buff)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Received:", string(buff))
			responseStr := fmt.Sprintf("your message: %v. time: %v",
				string(buff[:]), time.Now().Format(time.ANSIC))

			conn.Write([]byte(responseStr))
		}(conn)
	}

}
