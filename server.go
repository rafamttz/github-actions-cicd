package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/", sendHello)
	fmt.Println("Server started in port 8089")
	http.ListenAndServe(":8089", nil)
}

func sendHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Opa, bão?")
}
