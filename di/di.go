package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
	if err != nil {
		fmt.Println(err)
	}
}

// MyGreeterHandler says Hello, word over HTTP.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "word")
}

// Greet sends a personalised greeting to writer.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
