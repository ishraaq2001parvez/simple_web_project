package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := 9090
	fmt.Printf("Server is running on port %d\n", port)
	http.HandleFunc("/root", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello~"))
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatalln("Error: ", err)
	}
}
