package main

import (
	"fmt"
	"log"
	"net/http",
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "CIAO FROM A SIMPLE GOLANG HTTP SERVER!")
	fmt.Fprintf(w, "\n")
}

func main() {

	port := os.Getenv("SERVER_PORT", "8080")
	
	fmt.Printf("START\n")
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Printf("END\n")
}
