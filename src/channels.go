package main

import "fmt"

//se le puede indicar si es de salida  <-chan o de entrada chan<-
func sayWithChannel(text string, c chan string) {
	c <- text
}

func main() {
	//creamos un canal, le decimos que tipo de dato
	//maneja el canal
	//y cuantos datos simultaneamente
	c := make(chan string, 1)
	fmt.Println("hello")

	go sayWithChannel("Byeee", c)

	//utilizamos la salida del canal
	fmt.Println(<-c)

	otroCanal := make(chan string, 2)
	otroCanal <- "Hola"

	fmt.Println(len(otroCanal), cap(otroCanal))
	otroCanal <- "Adios"
	fmt.Println(len(otroCanal), cap(otroCanal))

	close(otroCanal)

	//si el canal esta cerrado o
	//lleno, no podemos darle mas inputs
	// esto no funcionaria daria error en runtime
	// otroCanal <- "otro"

	canal1 := make(chan string)
	canal2 := make(chan string)

	go sayWithChannel("mensaje numero 1", canal1)
	go sayWithChannel("mensaje numero 2", canal2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-canal1:
			fmt.Println("Recibido memsaje del canal 1", m1)
		case m2 := <-canal2:
			fmt.Println("Recibido memsaje del canal 2", m2)
		}
	}

}
