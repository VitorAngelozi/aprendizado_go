package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		escrever("ola")
		wg.Done()
	}()

	go func() {
		escrever("mundo")
		wg.Done()
	}()

	wg.Wait()

}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
}
