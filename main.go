package main

import (
	"fmt"
	"log"
	"net/http"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Contact Form Submit Successful\n")
	firstName := r.FormValue("firstName")
	surname := r.FormValue("surname")
	address := r.FormValue("address")
	phone := r.FormValue("phone")
	email := r.FormValue("email")

	fmt.Fprintf(w, "First Name = %s\n", firstName)
	fmt.Fprintf(w, "Surame = %s\n", surname)
	fmt.Fprintf(w, "Address = %s\n", address)
	fmt.Fprintf(w, "Phone = %s\n", phone)
	fmt.Fprintf(w, "Name = %s\n", email)
}

func msgHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/About Me Msg Goes Here" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", contactHandler)
	http.HandleFunc("/aboutMe", msgHandler)

	fmt.Printf("starting Server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
