package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercicio 1")
	Par(0)

	fmt.Println("Exercicio 2")
	fmt.Println(Soma([]int{1, 2, 3}))

	fmt.Println("Exercicio 3")
	Vogal("chinelo")

	fmt.Println("Exercicio 4")
	fmt.Println(Maior([]int{1, 2, 3, 4, 54, 6}))

}

// 1. Imprimir números pares
// Escreva uma função que recebe um número n e imprime todos os pares de 0 até n
func Par(n int) {
	var i int
	for i = 0; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

//2. Soma de slice
//Dado um slice []int, escreva uma função Soma(valores []int) int.

func Soma(valores []int) int {
	soma := 0
	for _, valor := range valores {
		soma += valor
	}
	return soma

}

//3. Contar vogais
//Função que recebe uma string e retorna quantas vogais ela tem.

func Vogal(palavra string) {

	vogal := []string{"a", "e", "i", "o", "u"}
	for _, letra := range palavra {
		for _, letra_vogal := range vogal {
			if string(letra) == string(letra_vogal) {
				fmt.Printf("A letra %s é vogal!\n", string(letra))
			}
		}
	}

}

//4. Encontrar o maior valor
//Função que recebe um slice de inteiros e retorna o maior

func Maior(numero []int) int {
	maior := 0

	for _, n := range numero {
		if n >= maior {
			maior = n
		}
	}
	return maior
}
