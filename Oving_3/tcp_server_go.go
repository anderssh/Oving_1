package main;

import (
	"fmt"
	"net"
	//"strconv"
	"os"
	//"bufio"
)

const (
	CONN_HOST = "localhost"
    CONN_PORT = "33546"
    CONN_TYPE = "tcp"
)

func handleConnection(conn net.Conn) {
	  // Make a buffer to hold incoming data.
	  buf := make([]byte, 1024)
	  fmt.Println("HEY")
	  // Read the incoming connection into the buffer.
	  msgLength, err := conn.Read(buf)
	  if err != nil {
	    fmt.Println("Error reading:", err.Error())
	  }
	  fmt.Println(msgLength)
	  fmt.Println(string(buf))
	  // Send a response back to person contacting us.
	  conn.Write([]byte("Message received."))
	  // Close the connection when you're done with it.
	  conn.Close()
	}

func main(){

	listner, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
	        fmt.Println("Error listening:", err.Error())
	        os.Exit(1)
	}
	fmt.Println("Del 1")
	TCPconnection, err := listner.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
	    os.Exit(1)
	}
	fmt.Println("Del 2")
	handleConnection(TCPconnection)
}
