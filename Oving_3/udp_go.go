package main;

import (
	"fmt"
	"net"
	"strconv"
);

func read(listen_conn *net.UDPConn, d_chan chan bool) {

	buf := make([]byte, 1024);
	//fmt.Println("R")

	for {
		//fmt.Println("YO")
		n, raddr, err := listen_conn.ReadFromUDP(buf);

		if (err != nil) {
			fmt.Println("PRoblem");
			break;
		}

		fmt.Println(string(buf));
		_ = n;
		_ = raddr;
		_ = err;
	}
}

func main() {

	broadcast_addr, err := net.ResolveUDPAddr("udp4", ":" + strconv.Itoa(30000));
	if err != nil { return; }

	tempConn, err := net.DialUDP("udp4", nil, broadcast_addr)
	defer tempConn.Close()
	temssAddr := tempConn.LocalAddr()
	laddr, err := net.ResolveUDPAddr("udp4", tempAddr.String())
	laddr.Port = 30000;
	fmt.Println(laddr)
	fmt.Println(broadcast_addr)

	//local_conn, err := net.ListenUDP("udp4", laddr);
	broad_conn, err := net.ListenUDP("udp", broadcast_addr);
	fmt.Println(err)

	d_chan := make(chan bool, 1);

	//go read(local_conn, d_chan);
	go read(broad_conn, d_chan);

	<- d_chan;
}
