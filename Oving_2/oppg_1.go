// go run oppg_4.go

package main

import (
    "fmt"
    "runtime"
    "time"
)

var i int = 20

func Goroutine1() {
    for j := 0; j < 1000001; j++ {
        i++;
    }
}

func Goroutine2() {
    for j := 0; j < 1000000; j++ {
        i--;
    }
}
func main() {

    runtime.GOMAXPROCS(runtime.NumCPU());

    i_chan := make(chan int, 1);
    
    go Goroutine1();
    go Goroutine2();

    time.Sleep(100*time.Millisecond);
    fmt.Println(i);
}
