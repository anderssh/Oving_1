package main;

import (
	"fmt"
	"net"
	"strconv"
	"time"
);

func listen(conn *net.UDPConn) {
	
	buffer := make([]byte, 1024);

	for {
		messageSize, _, _ := conn.ReadFromUDP(buffer);
		fmt.Println("Received: " + string(buffer[0:messageSize]));
	}
}

func transmit(conn *net.UDPConn) {
	
	for {
		time.Sleep(2000*time.Millisecond);

		message := "Hello server";
		conn.Write([]byte(message));
		fmt.Println("Sent: " + message);
	}
}

func main() {

	serverIP := "129.241.187.255";
	serverPort := 20016;
	serverAddr, _ := net.ResolveUDPAddr("udp", serverIP + ":" + strconv.Itoa(serverPort));

	listenPort := 20016;
	listenAddr, _ := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(listenPort));

	fmt.Println(listenAddr);
	fmt.Println(serverAddr);

	listenConn, _ := net.ListenUDP("udp", listenAddr);
	transmitConn, _ := net.DialUDP("udp", nil, serverAddr);

	go listen(listenConn);
	go transmit(transmitConn);

	d_chan := make(chan bool, 1);
	<- d_chan;
}