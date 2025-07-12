package main

import (
	"fmt"
	"net/http"
)

func formHandller(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Form submitted successfully!\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Email: %s\n", email)
}

func helloHandller(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandller)
	http.HandleFunc("/hello", helloHandller)

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", nil)
}