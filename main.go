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
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
