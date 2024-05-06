package main

import (
	// "errors"
	"fmt"
	"log"
	"net/http"
	// "os"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Form Error")
		return
	}

	fmt.Fprintf(w, "POST successful!")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name is %s\n", name)
	fmt.Fprintf(w, "Address is %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supproted", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello World!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting Server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
