package main

import "fmt"

func main() {

	myFunction("Ole Mundo")
	tripleArgunment(1, 2, "papa y arro")
	value := returningStuff(17)
	fmt.Println(value)

	value1, valu2 := doubleReturn(26)
	fmt.Println("value1 & value2", value1, valu2)

	value3, _ := doubleReturn(26)
	fmt.Println("la _ hace que no recoja el 2o valor", value3)

}

func tripleArgunment(a, b int, c string) {
	fmt.Println(a, b, c)
}

func myFunction(message string) {
	fmt.Println(message)
}

func returningStuff(n int) int {
	return n * 17
}

func doubleReturn(n int) (m, o int) {
	return n, n * 13
}
