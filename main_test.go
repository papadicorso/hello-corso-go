package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello Corso!!" {
		t.Fatal("Test fail")
	}
}