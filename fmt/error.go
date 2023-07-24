package main

import "fmt"

func main() {

	edad := 15

	if edad < 17 {
		err := fmt.Errorf("Es menor de edad %d. No puede pasar\n", edad)
		fmt.Println(err.Error())
		fmt.Println(err)
	}

	if edad >= 18 {
		respuesta := fmt.Sprintf("Tienes %d años de edad. Puedes pasar\n", edad)
		fmt.Println(respuesta)
	}

}
