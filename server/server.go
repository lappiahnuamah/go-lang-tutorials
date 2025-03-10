package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
		fmt.Fprintf(w, "This is my site serving with Go Lang Web Server")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is my about page using go lang")
		fmt.Fprintf(w, "This is my site serving with Go Lang Web Server. Enjoy")
	})

	fmt.Println("Server started on :8081")
	http.ListenAndServe(":8081", nil)
}
