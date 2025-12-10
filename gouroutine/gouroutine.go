package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá")
	escrever("Mundo")

}

func escrever(a string) {
	fmt.Println(a)
	time.Sleep(time.Second)
}
