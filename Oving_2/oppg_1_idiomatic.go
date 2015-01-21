// go run oppg_4.go

package main

import (
    "fmt"
    "runtime"
)

var i int = 20

func Increment(inc_event chan bool, done_event chan bool) {
    for j := 0; j < 1000001; j++ {
        inc_event <- true;
    }
    done_event <- true;
}

func Decrement(dec_event chan bool, done_event chan bool) {
    for j := 0; j < 1000000; j++ {
        dec_event <- true;
    }
    done_event <- true;
}
func main() {

    runtime.GOMAXPROCS(runtime.NumCPU());

    done1_event := make(chan bool);
    done2_event := make(chan bool);
    done1 := false;
    done2 := false;

    inc_event := make(chan bool);
    dec_event := make(chan bool);
    
    go Increment(inc_event, done1_event);
    go Decrement(dec_event, done2_event);

    for {
        select {
            case <- inc_event:
                i++;
            case <- dec_event:
                i--;
            case <- done1_event:
                done1 = true;
            case <- done2_event:
                done2 = true;
        }

        if done1 && done2 {
            break;            
        }
    }
    fmt.Println(i);
}
