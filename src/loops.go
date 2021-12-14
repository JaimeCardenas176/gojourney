package main

import "fmt"

func main() {
	//for, el unico ciclo iterativo de golang
	//for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}

	//for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//for forever
	//esto no para nunca, hasta que yo explicitamente lo pare
	counterForever := 0
	for {
		fmt.Println(counterForever)
	}

}
