package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

//HTML login form and form data
func RenderForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	//Serve html template
	if r.Method == "GET" {
		t, _ := template.ParseFiles("api_client.tpl")
		t.Execute(w, nil)
	}
}

func DeIndentify(w http.ResponseWriter, r *http.Request) {
	//Get data from html login form
	r.ParseForm()
	// a := Helper(r.Form["birthdate"])
	// fmt.Println(a)
	fmt.Println("BirthDate:", Helper(r.Form["birthdate"]))
	fmt.Println("ZipCode:", Helper(r.Form["zipcode"]))
	fmt.Println("Admission Date:", Helper(r.Form["admissiondate"]))
	fmt.Println("Discharge Date:", Helper(r.Form["dischargedate"]))
	fmt.Println("Notes:", Helper(r.Form["notes"]))
}

func Helper(data []string) string {
	//convert from []string to string
	res := strings.Join(data, " ")
	return res
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
