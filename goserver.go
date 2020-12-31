package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "CIAO FROM A SIMPLE GOLANG HTTP SERVER!")
	fmt.Fprintf(w, "\n")
}

func main() {
	fmt.Printf("START\n")
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Printf("END\n")
}
