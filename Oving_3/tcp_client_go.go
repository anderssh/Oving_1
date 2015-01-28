package main;

import (
	"fmt"
	"net"
	"strconv"
	//"os"
	//"bufio"
	
)

func main(){
	//tcpConnection, err := net.Dial("tcp", google.com:80);
	serverPort := 33546
	serverIp := "129.241.187.136"
	response := make([]byte,1024);

	tcpServerAddr, _ := net.ResolveTCPAddr("tcp", serverIp + ":" + strconv.Itoa(serverPort));
	tcpConnection, err := net.DialTCP("tcp", nil, tcpServerAddr);
	message := []byte("Connect to: 78.91.71.186 \x00")
	n,err := tcpConnection.Write(message)
	fmt.Println(message)
	fmt.Println("Antall bytes som er blitt sendt er:",n)
	fmt.Println("Feilmeldingen:",err)
	tcpConnection.Read(response)
	fmt.Println(string(response))
}