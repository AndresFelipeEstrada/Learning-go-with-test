package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Word")
}

func main() {
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
	// sleeper := &DefaultSleeper{}
	// Countdown(os.Stdout, sleeper)
	nombre := "Andres Felipe Estrada Arias"
	archivo := "archivo.txt"
	error := os.WriteFile(archivo, []byte(nombre), 0644)

	if error != nil {
		fmt.Println("Error al escribir archivo", error)
	}

	_, err := os.Stat(archivo)
	if err != nil {
		fmt.Println("El archivo no existe", err)
	} else {

		text, err := os.ReadFile(archivo)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(text))
	}
}
