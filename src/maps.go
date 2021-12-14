package main

import "fmt"

func main() {
	//diccionarios mapas
	mapa := make(map[string]int)

	mapa["jose"] = 14
	mapa["pepe"] = 18

	fmt.Println(mapa)

	//recorrer el map
	//no siempre vienen en el orden en el que se ingresaron
	//es aleatorio o concurrente
	for i, v := range mapa {
		fmt.Println(i, v)
	}

	//encontrar un valor
	value := mapa["jose"]
	test := mapa["gustav"]
	fmt.Println(value)
	fmt.Println(test)

	//checkear valor
	value2, ok1 := mapa["jose"]
	test2, ok2 := mapa["gustav"]
	fmt.Println(value2, ok1)
	fmt.Println(test2, ok2)
}
