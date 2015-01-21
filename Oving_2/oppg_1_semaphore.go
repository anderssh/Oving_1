// go run oppg_4.go

package main

import (
    "fmt"
    "runtime"
)

var i int = 20

func Goroutine1(flag_chan chan bool, done_chan chan bool) {
    for j := 0; j < 1000001; j++ {
        <- flag_chan;
        i++;
        flag_chan <- true; 
    }
    done_chan <- true;
}

func Goroutine2(flag_chan chan bool, done_chan chan bool) {
    for j := 0; j < 1000000; j++ {
        <- flag_chan;
        i--;
        flag_chan <- true; 
    }
    done_chan <- true;
}
func main() {

    runtime.GOMAXPROCS(runtime.NumCPU());

    flag_chan := make(chan bool, 1);
    flag_chan <- true;

    done1_chan := make(chan bool, 1);
    done2_chan := make(chan bool, 1);
    
    go Goroutine1(flag_chan, done1_chan);
    go Goroutine2(flag_chan, done2_chan);

    <- done1_chan;
    <- done2_chan;

    fmt.Println(i);
}
