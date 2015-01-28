package main;

import (
	"fmt"
	"net"
	//"strconv"
	//"os"
	//"bufio"
	
)
const(
SERVER_PORT = "3333"
SERVER_IP = "78.91.71.186"
CONN_TYPE = "tcp"
)

func main(){
	//tcpConnection, err := net.Dial("tcp", google.com:80);

	response := make([]byte,1024);

	tcpServerAddr, _ := net.ResolveTCPAddr(CONN_TYPE, SERVER_IP + ":" + SERVER_PORT);
	tcpConnection, err := net.DialTCP("tcp", nil, tcpServerAddr);
	message := []byte("Connect to: 78.91.71.186 \x00")
	n,err := tcpConnection.Write(message)
	fmt.Println(message)
	fmt.Println("Antall bytes som er blitt sendt er:",n)
	fmt.Println("Feilmeldingen:",err)
	tcpConnection.Read(response)
	fmt.Println(string(response))
}