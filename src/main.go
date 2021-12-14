package main

import "fmt"

func main() {
	//constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println("base", base)
	fmt.Println("altura", altura)
	fmt.Println("altura", area)

	//Zero values -> valores x default
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseC2 = 10
	areaC2 := baseC2 + baseC2
	println(areaC2)

}
