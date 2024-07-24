package main

import "testing"

func TestVerifyNumber(t *testing.T) {
	if verifyNumber(5) == false {
		t.Error("5 is not a number")
	}

	// Exécuter le test avec: go test tests/example_test.go
}
