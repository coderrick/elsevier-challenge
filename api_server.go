package main

import (
	"fmt"
	"log"
	"net/http"
)

func RenderForm(w http.ResponseWriter, r *http.Request) {
	//TODO
	fmt.Fprintf(w, "RenderForm")
}

func DeIndentify(w http.ResponseWriter, r *http.Request) {
	//TODO
	fmt.Fprintf(w, "DeIdentify")
}

func main() {
	log.Print("Listening on Port 8080; http://localhost:8080")
	http.HandleFunc("/", RenderForm) // setting router rule
	http.HandleFunc("/deidentify", DeIndentify)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
