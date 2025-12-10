package main

import (
	"fmt"
)

func main() {
	canal := make(chan int, 3) // canal com capacidade para 3 valores

	canal <- 1 // imediato
	canal <- 2 // imediato
	canal <- 3 // imediato — ainda há espaço

	// Se eu tentasse canal <- 4 aqui, travaria! (canal cheio)

	fmt.Println(<-canal) // lê o primeiro valor
	fmt.Println(<-canal) // lê o segundo valor
	fmt.Println(<-canal) // lê o terceiro valor

	canal <- 4
	canal <- 5
	canal <- 6
	fmt.Println(<-canal) // lê o primerio valor
	fmt.Println(<-canal) // lê o segundo valor
	fmt.Println(<-canal) // lê o terceiro valor
}
