package main

import "fmt"

func main() {
	figuras := []string{"circulo", "cuadrado", "triangulo", "rombo", "trapecio", "heptagono"}
	figurasCopia := make([]string, len(figuras))
	copy(figurasCopia, figuras)
	figuras = append(figuras[:1], figuras[2:]...)
	fmt.Println(figuras)
	fmt.Println(figurasCopia)
}
