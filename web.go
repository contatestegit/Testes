package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	// a := 2
	// b := 3
	// c := a + b
	// fmt.Fprintf(w, "Numero %d", c)
	fmt.Fprintf(w, "<h1>Titulo</h1>")
	fmt.Fprintf(w, "<p>Titulo</p>")

}

func value_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Number")
}

// func main() {
// 	http.HandleFunc("/", index_handler)
// 	http.HandleFunc("/continue", value_handler)
// 	http.ListenAndServe(":8000", nil)
// }
