package main

import "fmt"

func main() {
	//array -> inmmutables en go
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	//longitud
	fmt.Println(len(array))
	//capacidad
	fmt.Println(cap(array))

	fmt.Println(array[0])
	fmt.Println(array[:3])
	fmt.Println(array[1:3])
	fmt.Println(array[3:])

	//slices

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice)
	//longitud
	fmt.Println(len(slice))
	//capacidad
	fmt.Println(cap(slice))

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:6])
	fmt.Println(slice[4:])

	//no se pueden añadir objetos a los arrays(inmmutables)
	//pero si a los slice

	slice = append(slice, 8)
	fmt.Println(slice)

	newSlice := []int{9, 10, 11, 12}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
