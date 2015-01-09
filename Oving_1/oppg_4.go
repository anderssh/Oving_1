// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"     // Using '.' to avoid prefixing functions with their package names
                //   This is probably not a good idea for large projects...
    "runtime"
    "time"
)

var i int = 20

func Goroutine1() {
    for j := 0; j < 1000000; j++ {
        i++;
    }
}

func Goroutine2() {
    for j := 0; j < 1000000; j++ {
        i--;
    }
}
func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    go Goroutine1()
    go Goroutine2()                      
    time.Sleep(100*time.Millisecond)
    Println(i)

}
