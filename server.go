package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8001", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello User!!!</h1>"))
	
	name := os.Getenv("NAME")
	fmt.Fprintf(w, "Hello, I'm %s.", name)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w, "My Family: %s.", string(data))
}