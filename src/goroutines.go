package main

import (
	"fmt"
)

func saySomenthing(text string) {
	fmt.Println(text)
}

func main() {

	saySomenthing("Hello")
	saySomenthing("world")

	saySomenthing("Hello2")
	//con go routine
	go saySomenthing("world2")

	//no le da tiempo a llegar a la subrutina por eso no se imprime world2
	//el hilo se muere antes de ejecutar la goroutine ()

	//1.una manera basica pero no recomedable de solucionarlo seria
	//codigo -> duerme el hilo por un seegundo para que le de tiempo a la goroutine

	//time.Sllep(time.Second *1)

	//2. Otra manera un poco mas avanzada -> agregando un waitGroup
	//codigo -> para crear un waitGroup
	//var wg sync.WaitGroup

	//func saySomenthing(text string, wg *syncWaitGroup) {
	//	defer wg.Done()
	//	fmt.Println(text)

	//}
	//usamos wg.Add(1) para indicar que añadimos 1 goroutine
	// llamamos a la gouroutine -> go saySomenthing("hola", &wg)
	//wg.Wait() -> indicamos al wg que espere

	//funciones anonimas
	go func() {
		fmt.Print("adios anonimo")
	}() //-> aqui le pasamos el parametro en caso de tener

	go func(text string) {
		fmt.Print(text)
	}("adios como param")
}
