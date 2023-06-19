package main

import (
	"flag"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	var message string
	flag.StringVar(&message, "message", "Hello World!", "Message to send to server")
	flag.Parse()
	tcpAddr, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		println("Error resolving address", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP(TYPE, nil, tcpAddr)
	if err != nil {
		println("Error Dialing", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		println("Error writing", err.Error())
		os.Exit(1)
	}

	buff := make([]byte, 1024*10)
	_, err = conn.Read(buff[0:])
	if err != nil {
		println("Error reading", err.Error())
		log.Fatal(err)
		os.Exit(1)
	}
	println("Recieved message:", string(buff))

}
