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

func transmit(conn *net.UDPConn, serverAddr *net.UDPAddr) {
	
	for {
		time.Sleep(1000*time.Millisecond);

		message := "Yo server!";
		conn.WriteToUDP([]byte(message), serverAddr);
		fmt.Println("Sent: message to server");
	}
}

func main() {

	listenPort := 20016;

	broadcastIP := "129.241.187.255";
	broadcastAddr, _ := net.ResolveUDPAddr("udp4", broadcastIP + ":" + strconv.Itoa(listenPort));

	tempConn, _ := net.DialUDP("udp4", nil, broadcastAddr);
	defer tempConn.Close();
	tempAddr := tempConn.LocalAddr();
	localAddr, _ := net.ResolveUDPAddr("udp4", tempAddr.String());
	localAddr.Port = listenPort;

	localConn, _ := net.ListenUDP("udp4", localAddr);

	go receive(localConn);
	go transmit(localConn, broadcastAddr);

	d_chan := make(chan bool, 1);
	<- d_chan;
}