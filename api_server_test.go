package main

import "testing"

func TestBirthDateHelper(t *testing.T) {
	a := []string{"2000-01-01"}
	total := BirthDateHelper(a)
	if total != "20" {
		t.Errorf("BirthDate was incorrect, got: %s, want: %s.", total, "10")
	}
}
