package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8001", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello User!!!</h1>"))
	
	name := os.Getenv("NAME")
	fmt.Fprintf(w, "Hello, I'm %s.", name)
}