package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)

}

func getNames() (string, string) {
	return "John", "Doe"
}
