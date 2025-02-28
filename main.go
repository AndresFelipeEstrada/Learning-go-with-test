package main

import "fmt"

func main() {
	var usuarios int
	fmt.Println("cuantos usuarios quieres crear?")
	fmt.Scan(&usuarios)
	nombres := make([]string, usuarios)

	for i := range nombres {
		var nombre string
		fmt.Printf("Escoje el nombre del usuario %d\n", i+1)
		fmt.Scan(&nombre)
		nombres[i] = nombre
	}

	fmt.Println("\nLista completa de usuarios: ")
	fmt.Println(nombres)
}
