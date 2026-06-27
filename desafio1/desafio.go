package main

import "fmt"
const TAM = 100
func main(){
	fmt.Println("Iniciando o Sistema...")

	for i := 1; i < TAM; i++{
		if i % 3 == 0{
			fmt.Println(i)
		}
	}
}