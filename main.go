package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Jogo - Adivinhe o número!")
	var tentativas int = 0
	var chute int
	randomNumber := rand.Intn(100) + 1
	for {
		fmt.Println("Digite um número entre 1 e 100: ")
		tentativas++

		_, error := fmt.Scan(&chute)
		if error != nil {
			fmt.Println("Erro ao ler o número")
			return
		}

		if chute == randomNumber {
			fmt.Println("Parabéns! Você acertou!")
			fmt.Println("Você acertou em ", tentativas, " tentativas")
			return
		}
		if chute < randomNumber {
			fmt.Println("Seu número é menor que o número secreto")
		}
		if chute > randomNumber {
			fmt.Println("Seu número é maior que o número secreto")
		}
	}

}
