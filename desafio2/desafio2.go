package main

import "fmt"

const TAM = 100

func main() {
	fmt.Println("Iniciando o Sistema...")

	for i := 1; i <= TAM; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pan Pan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("Fim do Sistema...")
}
