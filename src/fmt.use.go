package main

import "fmt"

func main() {

	helloMessage := "Hello"
	worldMessage := "World"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	//%v cuando no sabes que tipo de dato puede ser
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Println(message)

	//Averiguar tipo
	fmt.Printf("tipo de hello: %T\n", helloMessage)
}
