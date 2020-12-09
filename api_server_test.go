package main

import "testing"

func TestBirthDateHelper(t *testing.T) {
	a := []string{"2000-01-01"}
	result := BirthDateHelper(a)
	if result != "20" {
		t.Errorf("BirthDate was incorrect, got: %s, want: %s.", result, "20")
	}
}

func TestZipCodeHelper(t *testing.T) {
	a := []string{"10013"}
	result := ZipCodeHelper(a)
	if result != "10000" {
		t.Errorf("ZipCode was incorrect, got: %s, want: %s.", result, "10000")
	}
}

func TestDateHelper(t *testing.T) {
	a := []string{"2019-03-12"}
	result := DateHelper(a)
	if result != "2019" {
		t.Errorf("ZipCode was incorrect, got: %s, want: %s.", result, "2019")
	}
}
