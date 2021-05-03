package main

import (
    "fmt"
    "time"
)

func main() {
    for {
	now := time.Now()
    millis := now.UnixNano() / 1000000
    fmt.Println(millis)
    time.Sleep(1000000000)
    }
}
