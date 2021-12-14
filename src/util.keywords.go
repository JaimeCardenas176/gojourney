package main

import "fmt"

func main() {

	//defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")
	//se pueden usar mas, pero es buena practica solo un defer por funcion

	//continue y break
	for i := 0; 1 < 10; i++ {
		fmt.Println(i)
		//continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		//break
		if i == 8 {
			fmt.Println("break")
			break
		}
	}
}
