package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
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

//Obfuscate patient info in form data
func DeIndentify(w http.ResponseWriter, r *http.Request) {
	//Get data from html login form
	r.ParseForm()
	fmt.Println("BirthDate:", BirthDateHelper(r.Form["birthdate"]))
	fmt.Println("ZipCode:", ZipCodeHelper(r.Form["zipcode"]))
	fmt.Println("Admission Date:", DateHelper(r.Form["admissiondate"]))
	fmt.Println("Discharge Date:", DateHelper(r.Form["dischargedate"]))
	fmt.Println("Notes:", NotesHelper(r.Form["notes"]))
}

//Parse birthdate form data
func BirthDateHelper(data []string) string {
	//convert from []string to string
	bd := strings.Join(data, " ")
	then, err := time.Parse("2006-01-02", bd)
	if err != nil {
		fmt.Println(err)
	}
	dur := time.Since(then)
	n := uint64(dur.Hours() / 8760)
	age := strconv.FormatUint(uint64(n), 10)
	return age
}

func ZipCodeHelper(data []string) string {
	//TODO
	return " "
}

//Parse and truncate admission/discharge date form data
func DateHelper(data []string) string {
	d := strings.Join(data, " ")
	return d[0:4]
}

//Find and replace SSN
func NotesHelper(data []string) string {
	d := strings.Join(data, " ")
	m := regexp.MustCompile("...-..-....")
	Str := "${1}XXX-XX-XXXX$2"
	ssn := m.ReplaceAllString(d, Str)
	return ssn
}

func ReadCSV(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func main() {
	a := "2019-03-12"
	fmt.Println(a[0:4])
	fmt.Println(time.Now().Format("2006-01-02"))
	log.Print("Listening on Port 8080; http://localhost:8080")
	http.HandleFunc("/", RenderForm) // setting router rule
	http.HandleFunc("/deidentify", DeIndentify)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
