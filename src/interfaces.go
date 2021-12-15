package main

import "fmt"

type figura interface {
	area() float64
}

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calculcarArea(f figura) {
	fmt.Println("Area: ", f.area())
}

func main() {

	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calculcarArea(myCuadrado)
	calculcarArea(myRectangulo)

	//*lista de interfaces* -> bonus para formar listas
	//que admitan diferentes tipos de valor

	myInterface := []interface{}{"un string", 125, 4.87, true}
	fmt.Println(myInterface...)
}
