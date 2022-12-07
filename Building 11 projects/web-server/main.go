/*
Building a Web server in Go with 3 endpoints:

root: provides a index.html file
/hello: hello function
/form: form function and form.html


*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler( w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return 
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
	}
	fmt.Fprint(w, "Hello!")
}

func formHandler( w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() failed, err: %v ", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	// tells server to look at the static directory for files
	fileServer := http.FileServer(http.Dir("./static"))
	// handles the url and searches the file server for the right file 
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
