package main


import (
	"fmt"
	"net"
	"flag"
)

func sendMessage(msg string) {
	networkType, localhost := "tcp", "localhost:8080"
	conn, err := net.Dial(networkType, localhost)
	if err != nil {
		//handle error here
		fmt.Println("Error connecting to:", localhost)
	} else {
		fmt.Println("Sending msg: ", msg)
		bMsg := []byte(msg)
		n, err := conn.Write(bMsg)
		if err != nil {
			fmt.Println("Error writing to connection")
		}
		fmt.Println("SUCCESS! Message length sent: ", n)
		conn.Close()
	}
}

func main() {
	msgPtr := flag.String("msg", "hello world", "the message you want to send")

	flag.Parse()

	fmt.Println("Trying to send message...")
	sendMessage(*msgPtr)

	//dial = net.Dial(
}
