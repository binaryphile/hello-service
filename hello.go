package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/binaryphile/hello-service/services/msgsvc"
)

func main() {
	http.HandleFunc("/msg", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", html.EscapeString(msgsvc.GetMessage()))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
