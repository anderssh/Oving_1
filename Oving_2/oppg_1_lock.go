// go run oppg_4.go

package main

import (
    "fmt"
    "runtime"
    //"time"
)

var i int = 20

func Goroutine1(i_chan chan int, done_chan chan bool) {
    for j := 0; j < 1000001; j++ {
        i = <- i_chan;
        i++;
        i_chan <- i; 
    }
    done_chan <- true;
}

func Goroutine2(i_chan chan int, done_chan chan bool) {
    for j := 0; j < 1000000; j++ {
        i = <- i_chan;
        i--;
        i_chan <- i; 
    }
    done_chan <- true;
}
func main() {

    runtime.GOMAXPROCS(runtime.NumCPU());

    i_chan := make(chan int, 1);
    i_chan <- i;

    done1_chan := make(chan bool, 1);
    done2_chan := make(chan bool, 1);
    
    go Goroutine1(i_chan, done1_chan);
    go Goroutine2(i_chan, done2_chan);

    <- done1_chan;
    <- done2_chan;

    fmt.Println(i);
}
