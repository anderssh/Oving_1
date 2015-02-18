package main;

import(
	"os"
	"fmt"
	"os/exec"
	"time"
	"net"
	"encoding/binary"
);

//-----------------------------------------------//

func listen(listenConnection *net.UDPConn, listenChannel chan int) {
	
	messageBuffer := make([]byte, 1024);

	for {
		messageLength, _, _ := listenConnection.ReadFromUDP(messageBuffer);
		
		data := binary.BigEndian.Uint64(messageBuffer[0:messageLength]);

		listenChannel <- int(data);
	}
}

func timeout(timeoutChannel chan bool, resetTimer chan bool) {

	start := time.Now();

	for {
		select {
			case <-resetTimer:
				start = time.Now();
			default:
				if time.Since(start) >= time.Millisecond*500 {
					timeoutChannel <- true;
				}
		}
	}
}

func backup(){

	fmt.Println("Creating backup...");

	i := 0;

	listenAddress, _ := net.ResolveUDPAddr("udp", "localhost:20005");
	listenConnection, _ := net.ListenUDP("udp", listenAddress);

	listenChannel := make(chan int);
	timeoutChannel := make(chan bool);
	resetTimer := make(chan bool);

	go listen(listenConnection, listenChannel);
	go timeout(timeoutChannel, resetTimer);

	for {
		select {
			case message := <- listenChannel:
				
				if message == -1{
					resetTimer <- true;
				} else {
					i = message;
				}

			case <- timeoutChannel:

				listenConnection.Close()

				close(listenChannel);
				close(timeoutChannel);
				close(resetTimer);

				go master(i);
		}
	}
}

//-----------------------------------------------//

func counting(i int, sendChannel chan int) {
	
	for {
		fmt.Println(i);
		sendChannel <- i;
		i++;
		time.Sleep(100*time.Millisecond)

	}
}
	
func notify(sendChannel chan int) {
	
	for {
		time.Sleep(20*time.Millisecond)
		sendChannel <- -1; 
	}
}

func send(sendChannel chan int) {
	
	for {
		message := <- sendChannel;

		sendAddress, _ := net.ResolveUDPAddr("udp", "localhost:20005");
		sendConnection, _ := net.DialUDP("udp", nil, sendAddress);

		bs := make([]byte, 4)
    	binary.LittleEndian.PutUint32(bs, uint32(message))

		sendConnection.Write(bs);
	}
}


func master(i int) {

	fmt.Println("Creating master...");
	
	command := exec.Command("gnome-terminal", "-e", "go run main.go");

	err := command.Run();
	if err != nil {
		fmt.Println(err);
		panic("Error in creating backup");
	}

	sendChannel := make(chan int);

	go send(sendChannel);
	go counting(i, sendChannel);
	go notify(sendChannel);

	deadChannel := make(chan bool)
	<- deadChannel;
}

//-----------------------------------------------//

func main() {
	
	args := os.Args[1:];
	role := "backup";

	if len(args) > 0 && args[0] == "master" {
		role = "master"
	}

	if role == "master" {
		go master(0);
	} else {
		go backup();
	}

	deadChannel := make(chan bool)
	<- deadChannel;
}