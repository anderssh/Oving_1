package main;

import(
	"os"
	"fmt"
	"os/exec"
)


func  backup(){
	//listen
}

func master(){
	//print
	//send
}

func main() {
	
	args := os.Args[1:];
	role := "backup";

	if len(args) > 0 && args[0] == "master" {
		role = "master"
	}

	if role == "master" {
		//make backup
		fmt.Println("asd");
		command := exec.Command("gnome-terminal", "-e", "");
		fmt.Println(command);
		err := command.Run();
		fmt.Println("Ferdig", err);
		//go master();
	} else {
		//go backup();
	}
	
	
}