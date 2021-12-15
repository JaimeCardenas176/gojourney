package main

import "fmt"

type pc struct {
	ram  int
	disk int
	cpu  string
}

//asi le decimos que es una funcion del struct pc
func (myPc pc) ping() {
	fmt.Println(myPc.cpu, "pong")
}

//ahora vamos a utilizar un puntero para acceder a la direccion
//donde hemos guardado la variable del tipo de nuestro struct
//y modificar sus atributos
func (myPc *pc) duplicateRam() {
	//en la cabecera le indicamos el * antes del tipo para indicar
	//que vamos a acceder a sus valores mediante el puntero
	myPc.ram = myPc.ram * 2
}

func main() {
	//cuando creamos una variable lo que hace el runtime
	//es crear/buscar una direccion de memoria
	//y en ella almacena el valor
	a := 50

	//un puntero no es mas que un acceso de memoria
	//esto es el valor del puntero, que apunta a la
	//direccion de memoria donde se encuentra a
	//el valor de este puntero es una direccion de memoria
	b := &a

	//valor de a, alojado en una direccion
	fmt.Println(a)
	//direccion de memoria
	fmt.Println(b)
	//el valor de la direccion a la que esta apuntando
	fmt.Println(*b)

	//estoy modificando el valor en la direccion a
	//la que esta apuntando b
	*b = 100

	fmt.Println(*b)

	//como consecuencia tambien cambia el valor de a
	fmt.Println(a)

	//structs

	myPC := pc{ram: 16, disk: 512, cpu: "intel"}
	fmt.Println(myPC)

	//aqui estamos usando una funcion del struct

	myPC.ping()

	fmt.Println(myPC)

	myPC.duplicateRam()
	fmt.Println(myPC)

	myPC.duplicateRam()
	fmt.Println(myPC)

}
