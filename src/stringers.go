package main

import "fmt"

//Stringers
//una manera de customizar el output al imprimir structs

type pc struct {
	ram  int
	disk int
	cpu  string
}

//esto es el stringer?
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GBs de RAM, %d GBs de Disco y my procesador es un %s", myPc.ram, myPc.disk, myPc.cpu)
}

func main() {

	myPC := pc{ram: 16, disk: 512, cpu: "intel"}
	fmt.Println(myPC)

}
