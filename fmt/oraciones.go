package main

import "fmt"

func main() {
	fmt.Print("Hola Mundo\n")

	fmt.Printf("Un string: %s\n", "Apple Vision Pro")
	fmt.Printf("Un float: %f\n", 3.143542)
	fmt.Printf("Un booleano: %t\n", true)
	fmt.Printf("Cualquier valor: %v %v %v\n", 4.13, "\n", false)
	fmt.Printf("%d\n", 50)
	fmt.Printf("Un integer %d\n", 30)
	fmt.Printf("Signo tanto por ciento: %%\n")
	fmt.Printf("Tipos: %T %T %T\n", 1.18, "Filippo", true)

	fmt.Print("\n")

	fmt.Println("Esta funcion de fmt proporciona valores en la consola, agregando una nueva lineal al final")


	my_string := fmt.Sprint("Hello", "World")
	fmt.Println(my_string)

	//fmt.Sprintf testing. fmt.Sprintf no imprime en la pantalla.
	//Debes ponerlo dentro de una variable
	const nombre, edad = "Filippo", 22
	salida := fmt.Sprintf("%s tiene %d años\n", nombre, edad)
	fmt.Println(salida)
}
