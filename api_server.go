package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//HTML login form and form data
func RenderForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("api_client.tpl")
		t.Execute(w, nil)
	} else {
		//Get data from html login form
		r.ParseForm()
		fmt.Println("BirthDate:", r.Form["birthdate"])
		fmt.Println("ZipCode:", r.Form["zipcode"])
		fmt.Println("Admission Date:", r.Form["admissiondate"])
		fmt.Println("Discharge Date:", r.Form["dischargedate"])
		fmt.Println("Notes:", r.Form["notes"])
	}
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
