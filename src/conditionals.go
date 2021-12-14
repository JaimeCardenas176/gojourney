package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("es 1")
	} else {
		fmt.Println("no es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("es verdad")
	}

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("tambien es verdad")
	}

	fmt.Println(isEven(17))
	fmt.Println(checkUserPassword("12345"))
	fmt.Println(checkUserPassword("12345abcd"))

	//switch

	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	switch modulo2 := 5 % 2; modulo2 {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	value := 200
	switch {
	case value > 100:
		fmt.Println("es mayor a 100")
	case value < 100:
		fmt.Println("es menor a 100")
	default:
		fmt.Println("no menor ni mayor a 100")

	}

}

func isEven(number int) bool {
	return number%2 == 0
}

func checkUserPassword(password string) bool {
	const savedPassword string = "12345abcd"
	return password == savedPassword
}
