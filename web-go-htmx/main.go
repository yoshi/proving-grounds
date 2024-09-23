package main

import (
	"fmt"
	//	"html/template"
	"net/http"
)

func main() {

	fmt.Println("starting")
	http.ListenAndServe(":3000", nil)
}
