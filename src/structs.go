package main

//st es el alias que le puse
import st "fmt"

//modificador de acceso privado
type car struct {
	brand string
	year  int
	plate string
}

//si empieza por mayusculas el mod. de acceso es public
// -> ocurre igual con las funciones (...y structs...)
type CarPublic struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2016, plate: "1234"}
	st.Println(myCar)

	//otra manera
	var yourCar car
	yourCar.brand = "Renault"

	st.Println(yourCar)

}
