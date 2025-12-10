package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string)

	go escrever("olá mundo", canal)
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}
func escrever(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second)
	}
	close(canal)
}
