package main

import (
	"fmt"
	"net/http"
)

func main() {
	const port int = 9090
	fmt.Println("hello")
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
