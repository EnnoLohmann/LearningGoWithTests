package main

import (
	"learning_go_with_tests/dependency_injection"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(dependency_injection.MyGreeterHandler)))
}
