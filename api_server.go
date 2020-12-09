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
	d := strings.Join(data, " ")
	then, err := time.Parse("2006-01-02", d)
	if err != nil {
		fmt.Println(err)
	}
	dur := time.Since(then)
	n := uint64(dur.Hours() / 8760)
	age := strconv.FormatUint(uint64(n), 10)
	return age
}

//Maps csv data to a map and calculates new zipcode
func ZipCodeHelper(data []string) string {
	d := strings.Join(data, " ")
	zcta := make(map[string]string)
	lines, _ := ReadCSV("population_by_zcta_2010.csv")
	//populate map with csv data
	for _, line := range lines {
		zcta[line[0]] = line[1]
	}
	population, _ := strconv.Atoi(zcta[d])
	if population < 20000 {
		return "00000"
	} else {
		return d[0:3] + "00"
	}
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

//Read CSV file
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
	log.Print("Listening on Port 8080; http://localhost:8080")
	http.HandleFunc("/", RenderForm) // setting router rule
	http.HandleFunc("/deidentify", DeIndentify)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
