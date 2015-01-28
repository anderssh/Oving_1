package main;

import (
	"fmt"
	"net"
	"strconv"
	"time"
);

func receive(conn *net.UDPConn) {
	
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

	broadcastIP := "129.241.187.255";
	broadcastPort := 30000;
	broadcastAddr, _ := net.ResolveUDPAddr("udp", broadcastIP + ":" + strconv.Itoa(broadcastPort));

	serverIP := "129.241.187.255";
	serverPort := 20016;
	serverAddr, _ := net.ResolveUDPAddr("udp", serverIP + ":" + strconv.Itoa(serverPort));

	localAddr, _ := net.ResolveUDPAddr("udp", ":" + strconv(serverPort));

	fmt.Println(localAddr);
	fmt.Println(serverAddr);

	listenConn, _ := net.ListenUDP("udp", localAddr);
	transmitConn, _ := net.DialUDP("udp", nil, serverAddr);

	go receive(listenConn);
	go transmit(transmitConn);

	d_chan := make(chan bool, 1);
	<- d_chan;
}