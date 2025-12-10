package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan int), make(chan int)

	go func() {
		for {
			time.Sleep(time.Second)
			canal1 <- 1
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- 2
		}
	}()
	for {
		select {
		case msg1 := <-canal1:
			fmt.Println(msg1)
		case msg2 := <-canal2:
			fmt.Println(msg2)
		}
	}
}
